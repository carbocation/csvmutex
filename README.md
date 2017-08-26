# csvmutex
Wraps the CSV library with a mutex for the CSV writer to make it safe for use from concurrent goroutines.

# To install:
`go get github.com/carbocation/csvmutex`

# Usage example:
```go
package main

import "fmt"
import "os"

import "github.com/carbocation/csvmutex"

func main() {
  fmt.Println(func() error {
    file, err := os.Create("/tmp/output.csv")
    if err != nil {
      return err
    }

    csv := csvmutex.NewCSVMutex(file)
    csv.Writer.Comma = '\t'

    if err := csv.Write([]string{"Hello", "world"}); err != nil {
      return err
    }
    if err := csv.Flush(); err != nil {
      return err
    }

    return nil
  }())
}
```
