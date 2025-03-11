# libCLImate.Go Example - **bit_flags**

## Summary

Example illustrating assocation of flag specifications with bit flags and (optionally) a receiver variable, thus simplifying the command-line handling code.

## Source

```Go
// examples/bit_flags.go

package main

import (
	"fmt"

	clasp "github.com/synesissoftware/CLASP.Go"
	libclimate "github.com/synesissoftware/libCLImate.Go"

	"os"
)

const (
	BF_Sound int = 1 << iota
	BF_Vision
)

func main() {

	// Specify specifications, parse, and checking standard flags

	var flags = 0

	flag_Sound := clasp.Flag("--enable-sound").SetAlias("-s").SetHelp("Enables sound").SetBitFlags(BF_Sound, &flags)
	flag_Vision := clasp.Flag("--enable-vision").SetAlias("-v").SetHelp("Enables vision").SetBitFlags(BF_Vision, &flags)

	climate, _ := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.AddFlag(flag_Sound)
		cl.AddFlag(flag_Vision)

		cl.Version = []int{0, 0, 1}

		cl.InfoLines = []string{
			"libCLImate.Go Examples",
			"",
			":version:",
			"",
		}

		cl.UsageHelpSuffix = "" // suppresses the default

		return nil
	}, libclimate.InitFlag_PanicOnFailure)

	_, _ = climate.ParseAndVerify(os.Args, libclimate.ParseFlag_PanicOnFailure)

	// Program logic

	if 0 != (BF_Sound & flags) {
		fmt.Println("running with sound")
	}
	if 0 != (BF_Vision & flags) {
		fmt.Println("running with vision")
	}
	if 0 == flags {
		fmt.Println("running in default mode")
	}
}
```

## Usage

### No arguments

If executed with no arguments

```bash
go run examples/bit_flags.go
```

it gives the output:

```
running in default mode
```

### Show usage

If executed with the arguments

```bash
go run examples/bit_flags.go --help
```

it gives the output:

```
libCLImate.Go Examples

bit_flags 0.0.1

USAGE: bit_flags [ ... flags and options ... ]

flags/options:

	--help
		Shows this help and exits

	--version
		Shows version information and exits

	-s
	--enable-sound
		Enables sound

	-v
	--enable-vision
		Enables vision
```

### Specify flags and options in long-form

If executed with the arguments

```bash
go run examples/bit_flags.go --enable-sound
```

it gives the output:

```
running with sound
```

### Specify flags and options in short-form

If executed with the arguments

```bash
go run examples/bit_flags.go -v
```

it gives the (same) output:

```
running with vision
```

### Specify flags and options with combined short-form

If executed with the arguments

```bash
go run examples/bit_flags.go -sv
```

it gives the (same) output:

```
running with sound
running with vision
```


<!-- ########################### end of file ########################### -->

