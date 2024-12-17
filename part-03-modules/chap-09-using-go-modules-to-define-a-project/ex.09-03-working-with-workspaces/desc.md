### Defining multiple modules within a project

The Go module system is designed to manage dependencies and versions for the entire module, not for subsets or subprojects within a module. However, there might be situations where you have multiple distinct components or subprojects within your main project, and each of these components or subprojects has dependencies and version requirements. In such cases, you can structure your project in a way that each component is its own module, separate from the main project module. These submodules can be maintained as separate Go modules, each with its own `go.mod` file.

For example, if you have a project with a main component and two other components, and each component has unique dependencies, you can structure your project like this:

```
myproject/
├── mainmodule/
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── ...
├── secondmodule/
│   ├── othermain.go
│   ├── go.mod
│   ├── go.sum
│   ├── ...
├── thirdmodule/
│   ├── othermain.go
│   ├── go.mod
│   ├── go.sum
│   ├── ...
```

Each subcomponent/module (that is, secondmodule and thirdmodule) is treated as a separate Go module with its own go.mod file and dependencies.
It makes sense to create submodules in the following situations:

- `Components have different dependencies`: When different components within your project have distinct sets of dependencies, creating submodules allows you to manage these dependencies separately.
- `There are separate versioning requirements`: If different components need different versions of the same dependency, using submodules can help manage these version conflicts more effectively.
- `There’s component reusability`: When you intend to reuse a component across multiple projects, structuring it as a separate module can facilitate its reuse in various contexts
  There’s maintainability: Submodules can enhance code organization and maintainability as each component can be developed, tested, and maintained separately.

While you can technically create submodules within a project, it is not a customary practice, and it should be done when there is a clear need for separate dependency management, versioning, or code organization for distinct components within your project. Each submodule should have its own `go.mod` file that defines its specific dependencies and version requirements.

---

### Go workspaces

In Go 1.18, the Go workspaces feature was released, which improved the experience of working with multiple Go modules within the same project locally. Originally, when working with multiple Go modules in the same project, you would need to manually edit the Go module files for each module with the replace directive to use your local changes. Now, with Go workspaces, we can define a go.work file, specifying to use our local changes, and not have to worry about managing several go.mod files manually ourselves. This is particularly useful when working with larger projects, or projects that span multiple repositories.
