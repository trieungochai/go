## Introduction

A web server is a program that uses the HTTP protocol – hence, the HTTP server – to accept requests from any HTTP client (web browser, another program, and so on) and respond to them with an appropriate message. When we browse the internet with our browser, it will be an HTTP server that will send an HTML page to our browser and we will be able to see it. In some other cases, a server will not return an HTML page but a different message that’s appropriate to the client.

Excerpt From: Samantha Coyle. “Go Programming - From Beginner to Professional.” Apple Books.

---

### HTTP handler

To react to an HTTP request, we need to write something that, we usually say, handles the request; hence, we call this something a handler.

In Go, we have several ways to do that, and one way is to implement the handler interface of the `http` package.

```go
ServeHTTP(w http.ResponseWriter, r *http.Request)
```

So, whenever we need to create a handler for HTTP requests, we can create a struct that includes this method and we can use it to handle an HTTP request. Here’s an example:

```go
type MyHandler struct {}
func(h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
```

This is a valid HTTP handler and you can use it like so:

```go
http.ListenAndServe(":8080", MyHandler{})
```

Here, `ListenAndServe()` is a function that will use our handler to serve the requests; any struct that implements the handler interface will be fine.

The `ServeHTTP` method accepts `ResponseWriter` and a `Request` object. You can use them to capture parameters from the request and write messages to the response. The simplest thing, for example, is to let our server return a message:

```go
func(h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  _, err := w.Write([]byte("HI"))
  if err != nil {
    log.Printf("an error occurred: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
  }
}
```

The `ListenAndServe` method might return an error. If this happens, we will want the execution of our program to halt. One common practice is to wrap this function call with a fatal log:

```go
log.Fatal(http.ListenAndServe(":8080", MyHandler{}))
```

This will halt the execution and print the error message that’s returned by the `ListenAndServe` function.

---

### Simple routing

Now, we want to associate different messages with these different paths on our server. We will do this by introducing some simple routing to our server.

A path is what you see after 8080 in the URL, where 8080 is the port number we chose to run the server on. This path can be one number, a word, a set of numbers, or character groups separated by a `/`. To do this, we will use another function of the net/http package:

```go
HandleFunc(pattern string, handler func(ResponseWriter, \*Request))
```

---

### Handler versus handler function

`http.Handle` and `http.HandleFunc`, both of which have a path as their first parameter, but which differ in terms of the second parameter. These 2 functions both ensure that a specific path is handled by a function.

`http.Handler` is any struct that has a method with this signature:

```go
ServeHTTP(w http.ResponseWriter, r *http.Request)
```

In both cases, there will always be a function with `http.ResponseWriter` and `*http.Request` as parameters that will handle the path. When one or the other might be chosen may just be a matter of personal preference in many cases, but i might be important – when creating a complex project, for example – to choose the right method. Doing so will ensure that the structure of the project is optimal. Different routes may appear better organized if they’re handled by handlers that belong to different packages, or might have to perform very few actions, as in our previous case; and a simple function might prove to be the ideal choice.

In general, for simple projects where you have a handful of simple pages, you may opt for `HandleFunc`. For example, let’s say you want to have static pages and there is no complex behavior on each page. In this case, it would be overkill to use an empty struct just for returning static text. The handler is more appropriate whenever you need to set some parameters, or if you want to keep track of something. As a general rule, let’s say that if you have a counter, `Handler` is the best choice because you can initialize a struct with a count of 0 and then increment it.
