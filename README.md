# f3s

Parse flags regardless of consistent from arguments.

### Installation

```
go get github.com/enthusiast17/f3s
```

### Usage

```Go
package main

import (
	"fmt"

	"github.com/enthusiast17/f3s"
)

func main() {
    // Create new flags
	flags := f3s.SetNewFlags(os.Args[1:])

	// SetFlag returns a flag's content, 1st param is flag and 2nd param is description
	output := flags.SetFlag("output", "This flag will write ascii output to file")

	// SetFlag returns a flag's content, 1st param is flag and 2nd param is description
	reverse := flags.SetFlag("reverse", "This flag will convert ascii file to string")

	fmt.Println(output, reverse)
}

```

### Terminal

```Bash
user@ubuntu:~/example$ go run main.go "hello" there --reverse="example1" --output=example2
example1 example2
```
