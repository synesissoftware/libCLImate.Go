# libCLImate.Go Example - **flag_and_option_aliases**

## Summary

Example illustrating various kinds of *flag* and *option* aliases, including the combination of short-names.

## Source

```Go
// examples/flag_and_option_aliases.go

package main

import (

	clasp "github.com/synesissoftware/CLASP.Go"
	libclimate "github.com/synesissoftware/LibCLImate.Go"

	"fmt"
	"os"
)

func main() {

	// Specify aliases, parse, and checking standard flags

	flag_Debug			:=	clasp.Alias{ clasp.Flag, "--debug", []string{ "-d" }, "runs in Debug mode", nil, 0, nil }
	option_Verbosity	:=	clasp.Alias{ clasp.Option, "--verbosity", []string{ "-v" }, "specifies the verbosity", []string{ "terse", "quiet", "silent", "chatty" }, 0, nil }
	flag_Chatty			:=	clasp.Alias{ clasp.Flag, "--verbosity=chatty", []string{ "-c" }, "", nil, 0, nil }


	climate, err := libclimate.Init(func (cl *libclimate.Climate) (err error) {

		cl.AddFlag(flag_Debug)
		cl.AddOption(option_Verbosity)
		cl.AddAlias(flag_Chatty)

		cl.ParseFlags = 0

		cl.Version = "0.0.1"

		cl.InfoLines = []string { "libCLImate.Go Examples", "", ":version:", "" }

		return nil
	});
	if err != nil {

		fmt.Fprintf(os.Stderr, "failed to create CLI parser: %v\n", err)
	}

	result, _ := climate.Parse(os.Args, libclimate.ParseFlag_PanicOnFailure)


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

```
    go run examples/flag_and_option_aliases.go
```

it gives the output:

```
```

### Show usage

If executed with the arguments

```
    go run examples/flag_and_option_aliases.go --help
```

it gives the output:

```
libCLImate.Go Examples

flag_and_option_aliases 0.0.1

USAGE: flag_and_option_aliases [ ... flags and options ... ]

flags/options:

	--help
		Shows this helps and exits

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

```
    go run examples/flag_and_option_aliases.go --debug --verbosity=silent
```

it gives the output:

```
verbosity is specified as: silent
Debug mode is specified
```

### Specify flags and options in short-form

If executed with the arguments

```
    go run examples/flag_and_option_aliases.go -v silent -d
```

it gives the (same) output:

```
verbosity is specified as: silent
Debug mode is specified
```

### Specify flags and options in short-form, including an alias for an option-with-value

If executed with the arguments

```
    go run examples/flag_and_option_aliases.go -c -d
```

it gives the output:

```
verbosity is specified as: chatty
Debug mode is specified
```

### Specify flags and options with combined short-form

If executed with the arguments

```
    go run examples/flag_and_option_aliases.go -dc
```

it gives the (same) output:

```
verbosity is specified as: chatty
Debug mode is specified
```
