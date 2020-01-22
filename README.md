# f3s

Parses flags regardless of consistent from arguments.

### Installation

```
go get github.com/enthusiast17/f3s
```

### Usage

Attention! You must use `flags.Parse(true/false)` after setting flags, otherwise flags won't be parsed.

```Go
package main

import (
	"fmt"

	"github.com/enthusiast17/f3s"
)

func main() {
    // Create new flags
	flags := fs.SetNewFlags()

	// 1st param is flag and 2nd param is description
	flags.SetFlag("output", "This flag will write ascii output to file")

	// 1st param is flag and 2nd param is description
	flags.SetFlag("reverse", "This flag will convert ascii file to string")

	// Parse function returns parsed flags
	// Bool param is displayHelper if true displays --help/-h command when there is no argument, else displays nothing
	flags.Parse(false)

        // HasContent returns a bool, if the given tag has a content returns 'true', else returns 'false'.
	if fs.HasContent(flags, "output") {
		output := fs.Content(flags, "output") // fs.Content returns a content from parsed flags
		fmt.Println(output)
	}

	// HasContent returns a bool, if the given tag has a content returns 'true', else returns 'false'.
	if fs.HasContent(flags, "reverse") {
		reverse := fs.Content(flags, "reverse") // fs.Content returns a content from parsed flags
		fmt.Println(reverse)
	}
}

```
