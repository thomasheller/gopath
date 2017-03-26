# gopath

gopath returns the current user's `GOPATH`, compatible with Go 1.8+
That is, it returns the value of the environment variable `GOPATH`,
and if that is not set, the directory named `go` in the user's home
directory.

What "home directory" means is platform-specific. gopath uses
[mitchellh/go-homedir](https://github.com/mitchellh/go-homedir) to
find out what the user's home directory is.

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/thomasheller/gopath"
)

func main() {
	p, err := gopath.Get()
	if err != nil {
		log.Fatalf("Error getting GOPATH: %v", err)
	}

	fmt.Println(p)

	// Output:
	// /home/user/go (or whatever GOPATH is set)
}
```

For convenience, gopath also provides a function to join the `GOPATH`
with one or more path elements, like `filepath.Join`:

```go
package main

import (
	"fmt"
	"log"

	"github.com/thomasheller/gopath"
)

func main() {
	p, err := gopath.Join("src/github.com/thomasheller/gopath")
	if err != nil {
		log.Fatalf("Error getting GOPATH: %v", err)
	}

	fmt.Println(p)

	// Output:
	// /home/user/go/src/github.com/thomasheller/gopath (or whatever GOPATH is set)
}
```
