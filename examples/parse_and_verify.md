# libCLImate.Go Example - **parse_and_verify**

## Summary

Exactly equivalent example to **flag_and_option_specifications**, but using ``AddFlagFunc()`` and ``AddOptionFunc()`` methods that specify callbacks executed when the flag/option is encountered in the parsed command-line.

## Source

```Go
// examples/parse_and_verify.go

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

	is_debug := false
	verbosity := ""

	climate, err := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.AddFlagFunc(flag_Debug, func() {

			is_debug = true
		})
		cl.AddOptionFunc(option_Verbosity, func(o *clasp.Argument, a *clasp.Specification) {

			verbosity = o.Value
		})
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

	_, _ = climate.ParseAndVerify(os.Args, libclimate.ParseFlag_PanicOnFailure)

	// Program-specific processing of flags/options

	if 0 != len(verbosity) {

		fmt.Printf("verbosity is specified as: %s\n", verbosity)
	}

	if is_debug {

		fmt.Printf("Debug mode is specified\n")
	}

	// Finish normal processing

	return
}
```

## Usage

### No arguments

If executed with no arguments

```bash
go run examples/parse_and_verify.go
```

it gives the output:

```
```

### Show usage

If executed with the arguments

```bash
go run examples/parse_and_verify.go --help
```

it gives the output:

```
libCLImate.Go Examples

parse_and_verify 0.0.1

USAGE: parse_and_verify [ ... flags and options ... ]

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
go run examples/parse_and_verify.go --debug --verbosity=silent
```

it gives the output:

```
verbosity is specified as: silent
Debug mode is specified
```

### Specify flags and options in short-form

If executed with the arguments

```bash
go run examples/parse_and_verify.go -v silent -d
```

it gives the (same) output:

```
verbosity is specified as: silent
Debug mode is specified
```

### Specify flags and options in short-form, including an alias for an option-with-value

If executed with the arguments

```bash
go run examples/parse_and_verify.go -c -d
```

it gives the output:

```
verbosity is specified as: chatty
Debug mode is specified
```

### Specify flags and options with combined short-form

If executed with the arguments

```bash
go run examples/parse_and_verify.go -dc
```

it gives the (same) output:

```
verbosity is specified as: chatty
Debug mode is specified
```


<!-- ########################### end of file ########################### -->

