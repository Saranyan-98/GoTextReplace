# GoTextReplace

GoTextReplace is a powerful Go package designed to facilitate bulk text replacement, making it effortless to find and replace content in multiple files. Whether you're managing configuration files, templating, or any other text-based document manipulation, GoTextReplace has you covered.

**Minimum Go Version Requirement: 1.15**

## Quick Start

To quickly get started with GoTextReplace, follow these simple steps:

1. Import the package:

   ```go
   import "github.com/your-username/GoTextReplace"
   ```

2. Initialize the TextReplace struct with your configuration:

```go
replacer := GoTextReplace.TextReplace{
    Filename: "sample.txt",
    YAMLfile: "tags.yaml",
}
```

3. Use the TextReplace method to execute bulk text replacement:

```go
err := replacer.Run()
if err != nil {
    fmt.Println(err)
}
```

### Sample YAML file

Here's an example of a YAML configuration file (tags.yaml):

```yaml
name: saranyan
company: XYZ
```

### Sample Text file

The source file (sample.txt) contains handlebar tags that need replacement:

```text
My name is {{name}} and I work at {{company}}
```

### Custom Output File

You can also specify a custom output file and path for the modified content:

```go
replacer := GoTextReplace.TextReplace{
    Filename:       "sample.txt",
    YAMLfile:       "tags.yaml",
    OutputFileName: "final.txt",
    OutputPath:     "./output",
}

err := replacer.Run()
if err != nil {
    fmt.Println(err)
}

```
