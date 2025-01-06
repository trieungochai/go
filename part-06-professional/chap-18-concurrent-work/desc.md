## Introduction

There is other software, however, that is meant to be used by several users at the same time. They are designed to serve websites or web applications that are generally used by thousands of users at the same time.

When multiple users are accessing a web server, it sometimes needs to perform a series of actions that are totally independent and whose result is the only thing that matters to the final output. All these situations call for a type of programming in which different tasks can be executed at the same time, independently from each other. Some languages allow parallel computation, where tasks are computed simultaneously.

In concurrent programming, when a task starts, all other tasks start as well, but instead of completing them one by one, the machine performs a bit of each task at the same time. While Go allows concurrent programming, tasks can also be performed in parallel when the machine has multiple cores. From the perspective of the programmer, however, this distinction is not that important, as tasks are created with the idea that they will be performed in parallel and in whatever way the machine will perform them.

---

### Goroutines

Imagine several people have some nails to hammer into a wall. Each person has a different number of nails and a different area of the wall, but there is only one hammer. Each person uses the hammer for one nail, then passes the hammer to the next person, and so on. The person with the fewest nails will finish earlier, but they will all share the same hammer.

This is how Goroutines work. Using Goroutines, Go allows multiple tasks to run at the same time (they are also called coroutines). These are routines (read tasks) that can co-run inside the same process but are totally concurrent. Goroutines do not share memory, which is why they are different from threads. However, we will see how easy it is to pass variables across them in your code and how this might lead to some unexpected behavior.

Writing a Goroutine is nothing special; they are just normal functions. Each function can easily become a Goroutine; all we must do is write the word go before calling the function.

Let us consider a function called `hello()`:

```go
func hello() {
  fmt.Println("hello world")
}
```

To call our function as a Goroutine, we do the following:

```go
go hello()
```

The function will run as a Goroutine. What this means can be understood better through the following code:

```go
func main() {
  fmt.Println("Start")
  go hello()
  fmt.Println("End")
```

The code starts by printing `Start`, then it calls the `hello()` function. Then, the execution goes straight to printing End without waiting for the `hello()` function to complete. No matter how long it takes to run the `hello()` function, the `main()` function will not care about the `hello()` function as these functions will run independently.

#### NOTE

The important thing to remember is that Go is not a parallel language but concurrent, which means that Goroutines do not work in an independent manner, but each Goroutine is split into smaller parts and each Goroutine runs one of its subparts at a time.

---

### WaitGroup

```go
func sum(fromNum, toNum int) int {
	result := 0
	for i := fromNum; i <= toNum; i++ {
		result += i
	}

	return result
}

func main() {
	var batch1, batch2 int
	go func() {
		batch1 = sum(1, 100)
	}()
	batch2 = sum(1, 10)

	time.Sleep(time.Second)
	log.Println(batch1, batch2)
}
```

In the previous exercise, we used a not-so-elegant method to ensure that the Goroutine ended by making the main Goroutine wait for a second.

The important thing to understand is that even if a program does not explicitly use Goroutines via the go call, it still uses one Goroutine, which is the main routine. When we run our program and create a new Goroutine, we are running 2 Goroutines: the main one and the one we just created. In order to synchronize these 2 Goroutines, Go gives us a function called `WaitGroup`.

`WaitGroup` needs the sync package to be imported. Typical code using the `WaitGroup` will be something like this:

```go
package main

import "sync"

func main() {
  wg := &sync.WaitGroup{}
  wg.Add(1)
  …………………..
  wg.Wait()
  ………….
  ………….
}
```

Here, we create a pointer to a new `WaitGroup`, then we mention that we are adding an asynchronous operation that adds `1` to the group using `wg.Add(1)`. This is essentially a counter holding the number of all concurrent Goroutines that are running. Later, we add the code that will run the concurrent call. At the end, we tell the `WaitGroup` to wait for the Goroutines to end using `wg.Wait()`.

How does the `WaitGroup` know that the routines are complete? We need to explicitly tell the `WaitGroup` about it inside the Goroutine with the following:

```go
wg.Done()
```

This must reside at the end of the called Goroutine.

---

---

### Race conditions

One important thing to consider is that whenever we run multiple functions concurrently, we have no guarantee in what order each instruction in each function will be performed.

In many architectures, this is not a problem. Some functions are not connected in any way with other functions, and whatever a function does in its Goroutine does not affect the actions performed in other Goroutines. This is, however, not always true. The first situation we can think of is when some functions need to share the same parameter. Some functions will read from this parameter, while others will write to this parameter. As we do not know which operation will run first, there is a high likelihood that one function will override the value updated by another function.

Let’s see an example that explains this situation:

```go
func next(v *int) {
  c := *v
  *v = c + 1
}
```

This function takes a pointer to an integer as a parameter. It is a pointer because we want to run several Goroutines with the `next()` function and update `v`. If we run the following code, we would expect that a will hold the value `3`:

```go
a := 0
next(&a)
next(&a)
next(&a)
```

This is perfectly fine. However, what if we run the following code:

```go
a := 0
go next(&a)
go next(&a)
go next(&a)
```

In this case, we might see that a holds `3`, or `2`, or `1`. Why would this happen? Because when a function executes the following statement, the value of `v` might be `0` for all functions running in independent Goroutines:

```go
c := *v
```

If this happens, then each function will set `v` to `c + 1`, which means none of the Goroutines are aware of what the other Goroutines are doing and override any changes made by another Goroutine. This problem is called a race condition and happens every time we work with shared resources without taking precautions. Fortunately, we have several ways to prevent this situation and to make sure that the same change is made only once. We will look at these solutions in the next sections, and we will explore the situation we just described in more detail, with a proper solution and race detection.

---

### Atomic operations

Let’s imagine we want to run independent functions again. However, in this case, we want to modify the value held by a variable. We still want to sum the numbers from 1 to 100, but we want to split the work into 2 concurrent Goroutines. We can sum the numbers from 1 to 50 in one routine and the numbers from 51 to 100 in another routine.

At the end, we will still need to receive the value of 5050, but two different routines can add a number at the same time to the same variable. Let’s see an example with only four numbers where we want to sum 1, 2, 3, and 4, and the result is 10.

Think of it like having a variable called `s := 0` and then making a loop where the value of `s` becomes the following:

```go
s = 0
s = 1
s = 3 // (1 + 2)
s = 6
s = 10
```

However, we could also have the following loop. In this case, the order in which the numbers are summed is different:

```go
s = 0
s = 1
s = 4 // 3 + 1, the previous value of 1
s = 6 // 2 + 4 the previous value of 4
s = 10
```

Essentially, this is just the commutative property of the sum, but this gives us a hint that we can split the sum into 2 or more concurrent calls. The problem that arises here is that all the functions need to manipulate the same variable, `s`, which can lead to race conditions and incorrect final values. A race condition happens when two processes change the same variable, and one process overrides the changes made by another process without considering the previous change. Thankfully, we have a package called atomic that allows us to safely modify variables across Goroutines.

This package has some functions for executing simple concurrent safe operations on variables. Let’s look at an example:

```go
func AddInt32(addr *int32, delta int32) (new int32)
```

This code takes a pointer to `int32` and modifies it by adding the value it points at to the value of delta. If `addr` holds a value of 2 and delta is 4, after calling this function, `addr` will hold 6.

---

### Invisible concurrency

It is easy to understand that concurrency problems are difficult to visualize as they do not manifest in the same way every time we run a program. That’s why we are focusing on finding ways to synchronize concurrent work.

One easy way to visualize it, however, but that is difficult to use in tests, is to print out each concurrent routine and see the order in which these routines are called.

If we want to see the effects of concurrency and still be able to test it, we could use the atomic package again, this time with strings so that we can build a string containing a message from each Goroutine.

For this scenario, we will use the `sync` package again, but we will not make use of `atomic` operations. Instead, we will use a new struct called `Mutex`. A mutex, short for mutual exclusion, serves as a synchronization primitive in Go, allowing multiple Goroutines to coordinate access to shared resources. When a Goroutine acquires a mutex, it locks it, ensuring exclusive access to the critical section of code. This prevents other Goroutines from accessing the same resource until the mutex is unlocked. Once the critical section execution is complete, the mutex is unlocked, allowing other Goroutines to acquire it and proceed with their execution concurrently. Let’s see how we can use it, we create a mutex like this:

```go
mtx := sync.Mutex{}
```

But most of the time, we want to pass a mutex across several functions, so we’d better create a pointer to a mutex:

```go
mtx := &sync.Mutex{}
```

This ensures we use the same mutex everywhere. It is important to use the same mutex, but the reason why the mutex must be only one will be clear after analyzing the methods in the Mutex struct. If all Goroutines have `mtx.Lock()` before modifying a value in a critical section of code such as in the following case, then only one Goroutine at a time can modify the variable due to the lock:

```go
mtx.Lock()
s = s + 5
```

The preceding code snippet will lock the execution of all the routines, except the one that will change the variable. At this point, we will add `5` to the current value of `s`. After this, we release the lock using the following command so that any other Goroutine can modify the value of `s`:

```go
mtx.Unlock()
```

From now on, any following code will run concurrently. We will see later some better ways to ensure safety when we modify a variable, but, for now, do not worry about adding much code between the `lock/unlock` part. The more code there is between these constructs, the less concurrent your code will be. So, you should lock the execution of the program, add only the logic required to ensure safety, unlock, and then carry on with the execution of the rest of the code, which does not touch the shared variables.

One important thing to notice is that the order of asynchronously performed code can change. This is because Goroutines run independently and you cannot know which one runs first. Furthermore, mutex-protected code can only be run by one Goroutine at a time, and you should then not rely on Goroutines to order things correctly; you might need to order your results afterward if you need a specific order.

---

### Channels

A channel is what the name essentially suggests – it’s something where messages can be piped, and any Goroutine can send or receive messages through a channel. Similar to that of a slice, a channel is created the following way:

```go
var ch chan int
ch = make(chan int)
```

Of course, it is possible to instantiate the channel directly with the following:

```go
ch := make(chan int)
```

Just like with slices, we can also do the following:

```go
ch := make(chan int, 10)
```

Here, a channel is created with a buffer of 10 items.

A channel can be of any type, such as integer, Boolean, float, and any struct that can be defined, and even slices and pointers, though the last two are generally used less frequently.
Channels can be passed as parameters to functions, and that’s how different Goroutines can share content. Let’s see how to send a message to a channel:

```go
ch <- 2
```

In this case, we send the value of 2 to the preceding ch channel, which is a channel of integers. Of course, trying to send something other than an integer to an integer channel will cause an error.

After sending a message, we need to be able to receive a message from a channel. To do that, we can just do the following:

```go
<- ch
```

Doing this ensures that the message is received; however, the message is not stored. It might seem useless to lose the message, but we will see that it might make sense. Nevertheless, we might want to keep the value received from the channel, and we can do so by storing the value in a new variable:

```go
i := <- ch
```

Let’s see a simple program that shows us how to use what we’ve learned so far:

```go
package main

import "log"

func main() {
  ch := make(chan int, 1)
  ch <- 1
  i := <- ch
  log.Println(i)
}
```

This program creates a new channel, pipes the integer 1 in, then reads it, and finally prints out the value of i, which should be 1. This code is not that useful in practice, but with a small change, we can see something interesting. Let’s make the channel unbuffered by changing the channel definition to the following:

```go
ch := make(chan int)
```

If you run the code, you will get the following output:

```
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [chan send]:
main.main()
    /Users/ samcoyle/go/src/github.com/packt-book/Go-Programming---From-Beginner-to-Professional-Second-Edition-/Chapter19/Exercise19.04/main.go:8 +0x59Process finished with exit code 2
```

The message may be different depending on the version of Go you are using. Also, some errors such as these have been introduced in newer versions. In older versions, though, the compiler was more permissive. In this specific case, the problem is simple: if we do not know how big the channel is, the Goroutines wait indefinitely, and this is called a deadlock. You can think of an unbuffered channel as having a capacity of zero. If we try to put anything into it, it won’t hold the item – instead, it will block until we can pass the item through the channel to a variable, for example. We will see later how to handle them, as they require more than one routine running. With only one Goroutine, after we send the message, we block the execution, and there is no other Goroutine able to receive the message; hence, we have a deadlock.

Before we go further, let’s see one more characteristic of channels, which is that they can be closed. Channels need to be closed when the task they have been created for is finished. To close a channel, type in the following:

```go
close(ch)
```

Alternatively, you can defer the closing, as shown in the following code snippet:

```go
defer close(ch)
for i := 0; i < 100; i++ {
  ch <- i
}
return
```

In this case, after the return statement, the channel is closed as the closing is deferred to run after the return statement.

---

### Concurrency patterns

The way we organize our concurrent work is pretty much the same in every application.

We will look at one common pattern that is called a `pipeline`, where we have a source, and then messages are sent from one Goroutine to another until the end of the line, until all Goroutines in the pipeline have been utilized. Another pattern is the `fan out/ fan in` pattern where, work is sent to several Goroutines reading from the same channel.

All these patterns, however, are generally made of a source stage, which is the first stage of the pipeline and the one that gathers, or sources, the data, then some internal steps, and at the end, a `sink`, which is the final stage where the results of the process from all the other routines get merged. It is known as a sink because all the data sinks into it.

---

### Buffers

There are channels with a defined length and channels with an undetermined length:

```go
ch1 := make(chan int)
ch2 := make(chan int, 10)
```

A buffer is like a container that needs to be filled with some content, so you prepare it when you expect to receive that content. We said that operations on channels are blocking operations, which means the execution of the Goroutine will stop and wait whenever you try to read a message from the channel.

```go
i := <- ch
```

We know that before we can carry on with the execution of the code, we need to receive a message. However, there is something more about this blocking behavior. If the channel does not have a buffer, the Goroutine is blocked as well. It is not possible to write to a channel or to receive a channel. We’ll get a better idea of this with an example, and we will show how to use unbuffered channels to achieve the same result so that you will get a better understanding of what you’ve seen in the previous exercises.

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
fmt.Println(<-ch)
fmt.Println(<-ch)
```

If you put this code inside a function, you will see that it works perfectly and will display something as follows:

```
1
2
```

But what if you add an extra read? Let’s take a look:

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
ch <- 3
fmt.Println(<-ch)
fmt.Println(<-ch)
```

In this case, you will see an error:

```
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [chan send]:
main.main()
    /tmp/sandbox223984687/prog.go:9 +0xa0
```

This happens because the routine running this code is blocked after the buffer of size 2 is filled with a data size of 2 coming from the read operations (commonly referred to as reads), which results in the buffer being filled with data, which, in this case, has 2 data, and the buffer has a size of 2. We can increase the buffer:

```go
ch := make(chan int, 3)
```

And it will work again; we are just not displaying the third number.

Now, let’s see what happens if we remove the buffer. Try, and again you will see the previous error. This happens because the buffer is always full and the routine is blocked. An unbuffered channel is equivalent to the following:

```go
ch := make(chan int, 0)
```

We’ve used unbuffered channels without any issues. Let’s see an example of how to use them:

```go
package main

import "fmt"

func readThem(ch chan int) {
  for {
    fmt.Println(<- ch)
  }
}

func main() {
  ch := make(chan int)
  go readThem(ch)
  ch <- 1
  ch <- 2
  ch <- 3
}
```

If you run this program, you should see something as follows:

```
1
2
3
```

But there is a chance you could see fewer numbers. If you run this on the Go Playground, you should see this result, but if you run it on your machine, you might see fewer numbers. Try sending more numbers:

```go
ch <- 4
ch <- 5
```

At each addition, run your program; you might not see all the numbers. Basically, there are two Goroutines: one is reading messages from an unbuffered channel, and the main Goroutine is sending these messages through the same channel. Due to this, there is no deadlock. This shows that we can make use of unbuffered channels for read and write operations flawlessly by using two Goroutines. We still have, however, an issue with not all numbers showing up, which we can fix in the following way:

```go
package main

import "fmt"
import "sync"

func readThem(ch chan int, wg *sync.WaitGroup) {
  for i := range ch {
    fmt.Println(i)
  }
  wg.Done()
}

func main() {
  wg := &sync.WaitGroup{}
  wg.Add(1)
  ch := make(chan int)
  go readThem(ch, wg)
  ch <- 1
  ch <- 2
  ch <- 3
  ch <- 4
  ch <- 5
  close(ch)
  wg.Wait()
}
```

Here, we iterate over the channel inside the Goroutine, and we stop as soon as the channel gets closed. This is because when the channel gets closed, the range stops iterating. The channel gets closed in the main Goroutine after everything is sent. We make use of a WaitGroup here to know that everything is completed. If we were not closing the channel in the main() function, we would be in the main Goroutine, which would terminate before the second Goroutine would print all the numbers. There is another way, however, to wait for the execution of the second Goroutine to be completed, and this is with explicit notification, which we will see in the next exercise. One thing to notice is that even though we close the channel, the messages all still arrive at the receiving routine. This is because you can receive messages from a closed channel; you just can’t send more.

---

### Some more common practices

In all these examples, we’ve created channels and passed them through, but functions can also return channels and spin up new Goroutines. Here is an example:

```go
func doSomething() chan int {
  ch := make(chan int)
  go func() {
    for i := range ch {
      log.Println(i)
    }
  }()
  return ch
}
```

In this case, we can actually have the following in our `main()` function:

```go
ch := doSomething()
ch <- 1
ch <- 4
```

We do not need to call the `doSomething` function as a Goroutine because it will spin up a new one by itself.

Some functions can also return or accept, such as this one:

```go
<- chan int
```

Here’s another example:

```go
chan <- int
```

---

### HTTP servers

Essentially, an HTTP server runs as a single program and listens to requests in the main Goroutine. However, when a new HTTP request is made by one of the clients, a new Goroutine is created that handles that specific request. You have not done it manually, nor have you managed the server’s channels, but this is how it works internally. You do not actually need to send anything across the different Goroutines because each Goroutine and each request is independent since they have been made by different people.

However, what you must think of is how to not create race conditions when you want to keep a state. Most HTTP servers are stateless, especially if you’re building a microservice environment. However, you might want to keep track of things with a counter, or you might actually work with TCP servers, a gaming server, or a chat app where you need to keep the state and gather information from all the peers. The techniques you’ve learned in this chapter allow you to do so. You can use a mutex to make sure a counter is thread-safe or, better, routine-safe across all requests.

---

### Methods as Goroutines

So far, you’ve only seen functions used as Goroutines, but methods are simple functions with a receiver; hence, they can be used asynchronously too. This can be useful if you want to share some properties of your struct, such as for your counter in an HTTP server.
With this technique, you can encapsulate the channels you use across several Goroutines belonging to the same instance of a struct without having to pass these channels everywhere.

Here is a simple example of how to do that:

```go
type MyStruct struct {}
func (m MyStruct) doIt()
. . . . . .
ms := MyStruct{}
go ms.doIt()
```

---
