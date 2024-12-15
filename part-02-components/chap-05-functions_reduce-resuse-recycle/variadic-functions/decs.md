The nums() function shows that the variadic type of i is a slice of integers. Once in the function, i will be a slice of integers. The variadic type has a length and capacity, which is to be expected for a slice. 
In the following code snippet, we will try to pass a slice of integers to a variadic function, nums():

```go
package main
import "fmt"
func main() {
  i := []int{ 5, 10, 15}
  nums(i)
}
func nums(i ...int) {
  fmt.Println(i)
}
```

The expected output is as follows:
![alt text](err.png)

Why didn’t this code snippet work? We just proved that the variadic variable inside the function is of the slice type. The reason is that the function expects a list of arguments of the int type to be converted into a slice. Variadic functions work by converting the arguments that are passed into a slice of the type being specified. However, Go has a mechanism for passing a slice to a variadic function.

For this, we need to use the unpack operator; it is three dots (…). When you call a variadic function and you want to pass a slice as an argument to a variadic parameter, you need to place the three dots before the variable:

```go
func main() {
  i := []int{ 5, 10, 15}
  nums(i…)
}
func nums(i ...int) {
  fmt.Println(i)
}
```