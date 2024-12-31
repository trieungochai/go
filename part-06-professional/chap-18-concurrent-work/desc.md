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
