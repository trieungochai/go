Duck typing is a test in computer programming: <em>If it looks like a duck, swims like a duck, and quacks like a duck, then it must be a duck</em>.

If a type matches an interface, then you can use that type wherever that interface is used. Duck typing is matching a type based upon methods, rather than the expected type:

```go
type Speaker interface {
  Speak() string
}
```

Anything that matches the Speak() method can be a Speaker{} interface. When implementing an interface, we are essentially conforming to that interface by having the required method sets:

```go
package main
import (
  "fmt"
)
type Speaker interface {
  Speak() string
}
type cat struct {
}
func main() {
  c := cat{}
  fmt.Println(c.Speak())
}
func (c cat) Speak() string {
  return "Purr Meow"
}
```
