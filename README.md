### Introduction

This is a simple package which can be used to find and replace a file in Bulk

##### Go version > 1.15

```go
    getMap := mapper.Mapper{}

    getMap.Filename = "sample.txt"
    getMap.YAMLfile = "tags.yaml"
    err := getMap.Mapper()
    if err != nil {
        fmt.Println(err)
    }
```

Sample YAML file

```yaml
name: saranyan
company: CleverInsight
```

Sample Text file
Handlebar Tagging is expected in the source file

```text
My name is {{name}} and I work at {{company}}
```
