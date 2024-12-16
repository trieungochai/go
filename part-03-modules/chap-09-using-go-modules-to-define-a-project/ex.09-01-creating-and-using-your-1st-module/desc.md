Your Go module does not have to be named the same as your Go package since you can have many packages to one Go module and project. It is good practice to name your module based on the main purpose of the project.

In this case, the primary purpose of the module is to manage and work with book chapters and authors, so the module’s name reflects the broader context. The name bookutil provides flexibility to include multiple packages related to book-related operations, including the author package.

In addition, there are best practices for module naming, such as <prefix>/<descriptive-text> and github.com/<project-name>/, that you can read more about in the Go documentation: https://go.dev/doc/modules/managing-dependencies#naming_module.

Now that you have successfully created a Go module named bookutil with an author package focused on book chapters, let’s explore the importance of using external Go modules and how they can enhance your project.
