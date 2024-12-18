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

### Package structure

It does not matter to Go how many different files are in a package. You should separate code into as many files as makes sense for readability and logic grouping.

However, all the files that are in a package must live in the same directory. The source files should contain code that is related, meaning that if the package is for configuration parsing, you should not have code in there for connecting to a database.

The basic structure of a package consists of a directory and contains one or more Go files and related code. The following diagram summarizes the core components of a package structure:
![package-structure](package-structure.png)

---

### Package naming

The name of your package is significant. It represents what your package contains and identifies its purpose. You can think of a package name as self-documentation. Careful consideration needs to go into naming a package.

The name of the package should be short and concise. It should not be verbose. Simple nouns are often chosen for a package name. The following would be poor names for a package:

```
stringconversion
synchronizationprimitives
measuringtime
```

Better alternatives would be the following:

```
strconv
sync
time
```

In Go, package names should be all lowercase with no underscores. Don’t use camel-case or snake-case styling. There are multiple packages with pluralized names.

Avoid package names such as misc, util, common, or data. These package names make it harder for the user of your package to understand its purpose.

![package-naming-conventions](package-naming-conventions.png)

---

### Package declarations

Every Go file starts with a package declaration. The package declaration is the name of the package. The first line of each file in a package must be the package declaration:

```
package <packageName>
```

All functions, types, and variables that are defined in the Go source file are accessible within that package. Though your package could spread across multiple files, it is all part of the same package. Internally, all code is accessible across the files. Simply stated, the code is visible within the package. Notice that not all of the code is visible outside of the package. The preceding snippet is from the official Go libraries. For a further explanation of the code, visit the links in the preceding Go snippet.

---

### Exported & Unexported Code

- Exported means that variables, types, functions, and so on are visible from outside of the package. If a function, type, variable, and so on starts with an uppercase letter, it is exportable.
- Unexported means it is only visible from inside the package. If it starts with a lowercase letter, it is unexportable.

```
NOTE
It is good practice to only expose code that we want other packages to see. We should hide everything else that is not needed by external packages.
```

---

### Main package

There are 2 basic types of packages in Go: `executable` and `non-executable`.

The main package is a special package. The main package is an executable package in Go. Logic that resides in this package may not be consumed by other packages.

The main package requires there to be a `main()` function in its package. The `main()` function is the entry point for a Go executable. When you perform go build on the main package, it will compile the package and create a binary. The binary is created inside the directory where the main package is located. The name of the binary will be the name of the folder it resides in:
![why-main-package-is-special](why-main-package-is-special.png)
