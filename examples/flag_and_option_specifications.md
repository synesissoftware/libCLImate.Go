# libCLImate.Go Example - **flag_and_option_specifications**

## Summary

Example illustrating various kinds of *flag* and *option* specifications, including the combination of short-names.

## Source

```Go
// examples/flag_and_option_specifications.go

package main

import (
	clasp "github.com/synesissoftware/CLASP.Go"
	libclimate "github.com/synesissoftware/libCLImate.Go"

	"fmt"
	"os"
)

func main() {

	// Specify specifications, parse, and checking standard flags

	flag_Debug := clasp.Flag("--debug").SetAlias("-d").SetHelp("runs in Debug mode")
	option_Verbosity := clasp.Option("--verbosity").SetAlias("-v").SetHelp("specifies the verbosity").SetValues("terse", "quiet", "silent", "chatty")

	climate, err := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.AddFlag(flag_Debug)
		cl.AddOption(option_Verbosity)
		cl.AddAlias("--verbosity=chatty", "-c")

		cl.Version = "0.0.2"

		cl.InfoLines = []string{
			"libCLImate.Go Examples",
			"",
			":version:",
			"",
		}

		return nil
	})
	if err != nil {

		fmt.Fprintf(os.Stderr, "failed to create CLI parser: %v\n", err)
	}

	result, _ := climate.Parse(os.Args, libclimate.ParseFlag_PanicOnFailure)
	if err != nil {

		panic(err)
	}

	// Program-specific processing of flags/options

	if opt, found := result.LookupOption("--verbosity"); found {

		fmt.Printf("verbosity is specified as: %s\n", opt.Value)
	}

	if result.FlagIsSpecified("--debug") {

		fmt.Printf("Debug mode is specified\n")
	}

	result.Verify(libclimate.ParseFlag_PanicOnFailure)

	// Finish normal processing

	return
}
```

## Usage

### No arguments

If executed with no arguments

```bash
go run examples/flag_and_option_specifications.go
```

it gives the output:

```
```

### Show usage

If executed with the arguments

```bash
go run examples/flag_and_option_specifications.go --help
```

it gives the output:

```
libCLImate.Go Examples

flag_and_option_specifications 0.0.1

USAGE: flag_and_option_specifications [ ... flags and options ... ]

flags/options:

	--help
		Shows this help and exits

	--version
		Shows version information and exits

	-d
	--debug
		runs in Debug mode

	-v <value>
	--verbosity=<value>
		specifies the verbosity
		where <value> one of:
			terse
			quiet
			silent
			chatty

	-c
	--verbosity=chatty
```

### Specify flags and options in long-form

If executed with the arguments

```bash
go run examples/flag_and_option_specifications.go --debug --verbosity=silent
```

it gives the output:

```
verbosity is specified as: silent
Debug mode is specified
```

### Specify flags and options in short-form

If executed with the arguments

```bash
	go run examples/flag_and_option_specifications.go -v silent -d
```

it gives the (same) output:

```
verbosity is specified as: silent
Debug mode is specified
```

### Specify flags and options in short-form, including a specification for an option-with-value

If executed with the arguments

```bash
go run examples/flag_and_option_specifications.go -c -d
```

it gives the output:

```
verbosity is specified as: chatty
Debug mode is specified
```

### Specify flags and options with combined short-form

If executed with the arguments

```bash
go run examples/flag_and_option_specifications.go -dc
```

it gives the (same) output:

```
verbosity is specified as: chatty
Debug mode is specified
```


<!-- ########################### end of file ########################### -->

