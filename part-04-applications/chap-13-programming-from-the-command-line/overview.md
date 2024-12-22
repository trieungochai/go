## Introduction
UIs don’t always have to be a web application frontend web page. End users can interact with software through engaging command-line interfaces, as well as by using a command-line interface (`CLI`).

---
### Reading arguments
Command-line arguments are a fundamental aspect of building versatile and interactive command-line applications. Reading arguments allow developers to make their applications more dynamic and adaptable to user input.

Command-line arguments serve as a means for users to customize the behavior of a program without modifying its source code. By capturing input parameters from the command line, developers can create versatile applications that cater to different use cases and scenarios.

In Go, the os package serves as a straightforward way to access these arguments. The `os.Args` slice provides a convenient way to access command-line arguments. This allows developers to retrieve information such as file paths, configuration parameters, or any other input relevant to the application’s functionality. The ability to read command-line arguments enhances the user experience by making applications more interactive and user-friendly.

Moreover, command-line arguments enable automation and scripting, allowing users to pass inputs programmatically. This flexibility is particularly valuable in scenarios where the same program needs to be executed with different parameters, making it a powerful tool for scripting and automation tasks.

---
### Using flags to control behavior
The flags package provides a higher-level and more structured approach to reading arguments compared to directly using the os package. Flags simplify the process of parsing and handling command-line input, making it easier for developers to create robust and user-friendly command-line applications.

The flags package allows you to define flags with associated types and default values, making it clear what kind of input a user is expected to provide. It also automatically generates help messages, making your program more self-documenting.

Here’s a brief overview of how the flags package can help in reading and handling command-line arguments:
- <b>Define flags</b>: You can define flags, along with their types and default values. This provides a clear and structured way to specify expected inputs.
- <b>Parse flags</b>: After defining flags, you can parse the command-line arguments. This initializes flag variables with the values provided by a user.
- <b>Access flag values</b>: Once you have parsed the flag values that have been passed in, you can access the defined flags through variables and continue to work with them throughout the application.

Flags allow you to customize the behavior of your program without the need to modify the source code. For example, you can create flags that allow you to toggle behavior based on if a flag value is set. You can also use basic conditional logic pending the values set for certain flags.

---
### Streaming large amounts of data in and out of your application
In command-line applications, it is crucial to handle large amounts of data efficiently for performance and responsiveness purposes. Often, command-line applications may be a small part of a larger pipeline processing data. Most people are not going to want to sit around typing out a large amount of data, such as a dataset, piece by piece.

Go allows you to stream data to your applications so that you can process information in chunks, rather than all at once.

This allows you to effectively process large amounts of data, reduce memory overhead, and provide better scalability in the future.

When dealing with large amounts of data, it’s often stored in files. This can range from financial CSV files, analysis Excel files, or machine learning datasets.

There are a few main benefits of streaming data with Go:

- <b>Memory efficiency</b>: The program can read and process data line by line, reducing memory consumption, as you then don’t have to read the entire data into memory.
- <b>Real-time analysis</b>: Users can observe a real-time analysis of the results of processing their data.
- <b>Interactive interface</b>: You can enhance the command-line interface so that it accepts dynamic information or displays additional details when processing large amounts of data.

---
### Exit codes and command line best practices
Exit codes provide a way for command-line applications to communicate their status to the calling application. A well-defined exit code system allows users and the other scripts to understand whether the application executed successfully or encountered an issue when running.

In Go, the os package provides a straightforward way to set exit codes using the `os.Exit` function. Conventionally, an exit code of 0 indicates success, while any non-zero code signals an error.

For example, you can check the status code of the previous exercise and verify the successful status code. To do this, run echo `$?` in the terminal. `$?` is a special shell variable that holds the exit status of the last command that was executed, and the echo command prints it out. You’ll see the 0 exit code printout denoting a successful execution status, and no error. You can manually catch errors in the program and return non-zero code signals to denote errors. You can even create custom exit codes, such as the following:
```go
const (
  ExitCodeSuccess = 0
  ExitCodeInvalidInput = 1
  ExitCodeFileNotFound = 2
)
```
These can easily be used using `os.Exit`, by placing `os.Exit(ExitCodeSuccess)` in successful cases you want to exit, and by using one of the other error codes when you want to exit in certain circumstances.

While using proper exit codes is an important command line best practice, there are a few others to keep in mind:
- <b>Consistent logging</b>: Use meaningful messages to aid troubleshooting.
- <b>Clear usage information</b>: Provide clear and concise usage information, including flags and arguments. Also, some packages allow you to provide example commands. Those should be used to let others see how to use the commands easily.
- <b>Handle help and versioning</b>: Implement flags to display help and version information. This is good for making your application more user-friendly and providing a means to ensure they are on the latest version by checking the version information.
- <b>Graceful termination</b>: Exit codes should be considered and terminated gracefully, ensuring proper cleanup tasks are performed as needed.

---
### Knowing when to stop by watching for interrupts

When building robust applications, it’s crucial to handle interrupts gracefully, ensuring the software can respond appropriately to signals indicating it should stop or perform a specific action. In Go, the standard way to achieve this is by monitoring interrupt signals, allowing the application to shut down or clean up resources in an orderly manner.

Graceful shutdown, or termination, is an important concept in computer science in general. Unforeseen events, server maintenance, or external factors might require your application to stop gracefully. This could involve releasing resources, saving state, or notifying connected clients. A graceful shutdown ensures your application remains reliable and predictable, minimizing the risk of data corruption or loss.

Abruptly terminating an application without proper cleanup can lead to various issues, such as incomplete transactions, resource leaks, corrupted data, and more. Graceful shutdowns mitigate these risks by providing an opportunity to finish ongoing tasks and release acquired resources.

The operating systems communicate with running processes through signals. A process is simply a program running on a computer. The `os/signal` package provides a convenient way to handle these signals within Go programs. There are common interrupt signals, such as `SIGINT (Ctrl + C)` and `SIGTERM`, that provide a comprehensive strategy for graceful terminations.

The `signal.Notify` function allows you to register channels to receive specified signals. This sets the foundation for creating a mechanism to gracefully shut down your application upon receiving an interrupt signal. There are also effective shutdown patterns and best practices to keep in mind, such as ensuring closed network connections, saving state, and signaling goroutines, to finish up their tasks before exiting. Additionally, using timeouts and the context package enhances your application’s responsiveness during shutdown, preventing it from getting stuck indefinitely.

---
### Starting other commands from your application
Launching external commands from your Go application opens opportunities for interaction with other programs, processes, and system utilities. The `os/exec` package in Go provides functionalities for starting and interacting with external processes. You can run basic commands, capture their output, and handle errors seamlessly with this package. It serves as a foundation for more advanced command execution scenarios.

For example, the `os/exec` package allows you to customize the execution of commands by configuring attributes such as the working directory, environment variables, and more. You can also provide inputs to the subcommand through standard input streams from the originating command-line application.

By running other commands from within your command-line application, you can run some processes in the background, allowing the application to continue its execution while monitoring or interacting with parallel processes. You can even establish bidirectional communication with commands, enabling real-time interaction and data exchange between the command-line application and the external processes.

When starting other applications, cross-platform considerations must be kept in mind. There are differences in shell behavior and command paths across different operating systems. So, when executing subcommands from a command-line application, it is important to keep a consistent and reliable command execution in mind, regardless of the compute an end user may be using. Thankfully, the `os/exec` package provides a cross-platform solution for executing external commands in Go, making it easier to write code across different operating systems.

---
### Terminal UIs
Some of the latest updates to command-line programming with Go have included terminal UIs, or TUIs for short. Creating a TUI in Go opens a world of possibilities for building interactive command-line applications.

There are a few fundamental concepts that are involved with building UIs for the terminal:

- <b>Components</b>: TUIs are composed of various components, such as buttons, input fields, and/or lists
- <b>Layouts</b>: Arranging components in a structured layout is crucial for a clean and intuitive design
- <b>User input handling</b>: Processing user input from keyboard events is fundamental to interactive interfaces

Some TUI packages provide support for event handling, such as mouse events and key presses, dynamic updates based on input data, or customizing the appearance of UI components. There are several popular TUI packages available.

---
### go install
You can install Go command-line applications using the go install command. This command is a powerful tool that’s provided (among the many) by the Go toolchain to compile and install Go applications in your workspace’s bin directory. This allows you to run your applications globally from any terminal window. To install a Go application, you can simply navigate to the project’s root directory and run go install.

This command considers cross-platform compilation by providing the `GOOS` flag, where you can specify which operating system to target, as well as the `GOARCH` flag, where you can specify the underlying architecture to target.

An example of a common Go package that you can use to generate command-line interfaces in Go is the cobra package. This is also a tool you can use to bootstrap application scaffolding to rapidly develop Cobra-based applications if you would like to dive further into developing your programming in terms of command-line skills. This package provides a simple example of using the go install command:
```
go install github.com/spf13/cobra-cli@latest
```

The preceding command installs all of the dependencies to use the Cobra CLI. As a result, my machine knows about the tool, and you can easily work with the command-line program that you just installed, as shown here:
```
cobra-cli –help
```
Cobra is a CLI library for Go that empowers applications.

It can generate the necessary files to quickly create a Cobra application:
```
Usage:
 cobra-cli [command]
Available Commands:
 add Add a command to a Cobra Application
completion Generate the autocompletion script for the specified shell
help Help about any command
init Initialize a Cobra Application
Flags:
“-a, --author string author name for copyright attribution (default "YOUR NAME")
--config string config file (default is $HOME/.cobra.yaml)
-h, --help help for cobra-cli
-l, --license string name of license for the project
--viper use Viper for configuration
Use "cobra-cli [command] --help" for more information about a command.
```