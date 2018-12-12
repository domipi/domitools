domitools
=========

**Several golang tools for daily usage

## Latest news

- 2018/12/12: First commit exposing first lazy tools [`PoVToV`] 

## Synopsis

```go
package main

import (
    "github.com/domipi/domitools/lazy"
    "fmt"
)

func main() {
    strValue := "helloworld"
    strPtr := &strValue
    var nilStrPtr *string

    // Display "helloworld" (inferencing if necessary)
    fmt.Println(lazy.PoVToV(strValue).(string))
    fmt.Println(lazy.PoVToV(strPtr).(string))

    // Display "" (the zero value for a string)
    fmt.Println(lazy.PoVToV(nilStrPtr).(string))
}
```
