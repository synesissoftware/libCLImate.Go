# libCLImate.Go Example - **show_usage_and_version**

## Summary

Simple example supporting ```--help``` and ```--version```.

## Source

```Go
// examples/show_usage_and_version.go

package main

import (

	libclimate "github.com/synesissoftware/libCLImate.Go"

	"fmt"
	"os"
)

func main() {

	// Specify aliases, parse, and checking standard flags

	climate, err := libclimate.Init(func (cl *libclimate.Climate) (err error) {

		cl.Version = "0.0.1"

		cl.InfoLines = []string { "libCLImate.Go Examples", "", ":version:", "" }

		return nil
	});
	if err != nil {

		fmt.Fprintf(os.Stderr, "failed to create CLI parser: %v\n", err)
	}

	_, _ = climate.ParseAndVerify(os.Args, libclimate.ParseFlag_PanicOnFailure)


	// Finish normal processing

	fmt.Printf("no flags specified\n")
}
```

## Usage

### No arguments

If executed with no arguments

```
    go run examples/show_usage_and_version.go
```

it gives the output:

```
no flags specified
```

### Show usage

If executed with the arguments

```
    go run examples/show_usage_and_version.go --help
```

it gives the output:

```
libCLImate.Go Examples

show_usage_and_version 0.0.1

USAGE: show_usage_and_version [ ... flags and options ... ]

flags/options:

    --help
        Shows this help and exits

    --version
        Shows version information and exits
```

### Show version

If executed with the arguments

```
    go run examples/show_usage_and_version.go --version
```

it gives the output:

```
show_usage_and_version 0.0.1
```

### Unknown option

If executed with the arguments

```
    go run examples/show_usage_and_version.go --unknown=value
```

it gives the output (on the standard error stream):

```
show_usage_and_version: unrecognised flag/option: --unknown=value
```

with an exit code of 1

