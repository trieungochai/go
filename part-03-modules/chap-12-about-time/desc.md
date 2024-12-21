## Introduction
This chapter is dedicated to teaching you all you need to know about handling variables that represent time data. You will learn how to do it the “Go way”. First, we will start out with basic time creation, timestamps, and more; then, we will learn how to compare and manipulate time, calculate the duration between two dates, and create timestamps. Finally, we will learn how to format time according to our needs.

---
### Making time
Making time means declaring a variable that holds the time formatted in a specific way.
```go
package main

import (
  "fmt"
  "time"
)

func main(){
  // this is where the code goes.
}
```

Whenever we issue go `run <script>.go`, the `main()` function gets called and executes whatever is declared in it.

One of the most common jobs for the `time` package is to measure the duration of the execution of the script. We can do this by capturing the current time in a variable, at the beginning, and at the end so that we can calculate the difference and know how long the specific action took to complete.

The very first example is as follows:
```go
  start := time.Now()
  fmt.Println("The script has started at: ", start)
  fmt.Println("Saving the world...")
  time.Sleep(2 * time.Second)
  end := time.Now()
  fmt.Println("The script has completed at: ", end)”
```
The output from our script should look like this:
```
The script has started at: 2023-09-27 08:19:33.8358274 +0200 CEST”
m=+0.001998701
Saving the world...
The script has completed at: 2023-09-27 08:19:35.8400169 +0200 CEST m=+2.006161301
```

##### Note
Any type of operating system that you work with will provide two types of clocks to measure the time; one is called the `monotonic clock`, and the other is called the `wall clock`. The `wall clock` is what you see on a Windows machine in the taskbar; it’s subject to change and is usually synchronized with a public or corporate Network Time Protocol (`NTP`) server based on your current location. `NTP` is used to tell clients the time based on an atomic clock or from a satellite reference.
