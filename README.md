# goini

Minimal Golang INI parsing function

### Usage

Documentation is available via
[![GoDoc](https://godoc.org/github.com/karrick/goini?status.svg)](https://godoc.org/github.com/karrick/goini).

### Description

Sometimes a program needs to parse an INI file.  This library reads from either a file or an
`io.Reader` and parses the text into a `map[string]map[string]string` instance.  The default section
of an INI file is the `General` section.  Key-Value pairs in the INI text are assigned to the
`General` section until a new section header is read.

### Supported Use Cases

#### Parse

```Go
    config, err := goini.Parse(someReader)
    if err != nil {
        panic(err)
    }
    
    for sectionName, sectionMap := range config {
        fmt.Printf("[%s]\n", sectionName)
        for key, value := range sectionMap {
            fmt.Printf("%s = %s\n", key, value)
        }
    }
```

#### ParseFile

```Go
    config, err := goini.ParseFile(configPathname)
    if err != nil {
        panic(err)
    }
    
    for sectionName, sectionMap := range config {
        fmt.Printf("[%s]\n", sectionName)
        for key, value := range sectionMap {
            fmt.Printf("%s = %s\n", key, value)
        }
    }
```
