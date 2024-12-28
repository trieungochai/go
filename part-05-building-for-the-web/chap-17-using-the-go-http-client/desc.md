## Introduction
An HTTP client is something that is used to get data from or send data to a web server. Probably the most well-known example of an HTTP client is a web browser (such as Firefox, Chrome, and Microsoft Edge).

When you enter a web address into a web browser, it will have an HTTP client built in that sends a request to the server for data. The server will gather the data and send it back to the HTTP client, which will then display the web page in the browser. Similarly, when you fill out a form in a web browser, for example, when you log in to a website, the browser will use its HTTP client to send that form data to the server and then take appropriate action, depending on the response.

---
### The Go HTTP Client and its uses
The Go HTTP Client is part of the Go standard library, specifically the `net/http` package.

There are 2 main ways to use it.
- The 1st is to use the default HTTP client that is included in the `net/http` package. It’s simple to use and allows you to get up and running quickly.
- The 2nd way is to create your own HTTP client based on the default HTTP client. This allows you to customize the requests and various other things. It takes longer to configure, but it gives you much more freedom and control over the requests you send.

When using an HTTP client, you can send different types of requests. While there are many types of requests, we will discuss the two main ones: the `GET` request and the `POST` request. For instance, if you wanted to retrieve data from a server, you would send a GET request. When you enter a web address in your web browser, it will send a GET request to the server at that address and then display the data it returns. If you wanted to send data to the server, you would send a `POST` request. If you wanted to log into a website, you would `POST` your login details to the server.

---
