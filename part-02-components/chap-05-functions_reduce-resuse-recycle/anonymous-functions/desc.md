Anonymous functions are great for small snippets of code that you want to execute within a function.
We saw that when we need a small function that might not be reusable in our program, we can create an anonymous function and assign it to a variable.

Anonymous functions are used for (and in conjunction with) the following:
- Closure implementations
- defer statements
- Defining a code block to be used with a goroutine
- Defining a function for one-time use
- Passing a function to another function

The following is a basic declaration for an anonymous function:
```go
func main() {
  func() {
    fmt.Println("Greeting")
  }()
}
```
- Notice that we are declaring a function inside another function.
- With anonymous functions, there is no function name.
- The empty parentheses following the func keyword are where the function’s parameters would be defined for the function.
- Next is the open curly brace, {, which starts the function body.
- The function body is only a one-liner; it will print “Greeting”.
- The closing curly brace, }, denotes the end of the function.
- The last set of parentheses is called the <ins>execution parentheses</ins>. These parentheses invoke the anonymous function. The function will execute immediately.

------

We can also pass arguments to an anonymous function. To be able to pass arguments to an anonymous function, they must be supplied in the <ins>execution parentheses</ins>:
```go
func main() {
  message := "Greeting"
  func(str string) {
    fmt.Println(str)
  }(message)
} 
```
- `func (str string)`: The anonymous function being declared has an input parameter of the string type.
- `} (message)`: The argument message that’s being passed to the <ins>execution parentheses</ins>.

------
There are other ways to execute anonymous functions. We can also save the anonymous function to a variable.
```go
func main() {
  f := func() {
    fmt.Println("Executing an anonymous function using a variable")
  }
  fmt.Println("Line after anonymous function declaration")
  f()
}
```
- We are assigning the f variable to our anonymous function.
f is now of the func() type.
- f can now be used to invoke the anonymous function, in a fashion similar to that for a named function. You must include () after the f variable to execute the function.
