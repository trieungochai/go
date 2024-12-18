### What is a package?

Go follows the `Don’t Repeat Yourself (DRY)` principle.

This means that you should not write the same code twice. Refactoring your code into functions is the first step of the DRY principle.

- What if you had hundreds or even thousands of functions that you used regularly?
- How would you keep track of all those functions? Some of those functions might even have common characteristics.

You could have a group of functions that perform math operations, string manipulations, printing, or file-based operations. You may be thinking of breaking them up into individual files:
![group-functions-by-files](group-functions-by-files.png)

However, what if your string’s functionality started to grow further? You would then have a ton of string functions in one file or even multiple files. Every program you build would also have to include all of the code for `string`, `math`, and `io`. You would be copying code to every application that you built. Bugs in one code base would have to be fixed in multiple programs. That kind of code structure is not maintainable, nor does it encourage code reusability.

The packages in Go are the next step to organizing your code in a way that makes it easy to reuse the components of your code. The following diagram shows the progression of organizing code from functions to source files to packages:
![code-progression-org](code-progression-org.png)

Go organizes its code for reusability into directories called packages. A package is essentially a directory inside your workspace that contains one or more Go source files, which is used for grouping code that performs a task. It exposes only the necessary parts in order for those using your package to get a job done. The package concept is akin to using directories to organize files on a computer.

---
