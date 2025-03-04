# libCLImate.Go Example - **values**

## Summary

Illustrates the handling of `Values`.


## Source

```Go
// examples/values.go

package main

import (
	libclimate "github.com/synesissoftware/libCLImate.Go"

	"fmt"
	"os"
)

func main() {

	// Specify specifications, parse, and checking standard flags

	climate, _ := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.Version = []int{0, 0, 1}

		cl.InfoLines = []string{
			"libCLImate.Go Examples",
			"",
			":version:",
			"",
		}

		cl.ValueNames = []string{"country name", "state id", "city name", "district id"}
		cl.ValuesConstraint = []int{2, 4}
		cl.ValuesString = "<country-name> <state-id> [ <city-name> [ <district-id> ]]"

		return nil
	}, libclimate.InitFlag_PanicOnFailure)

	r, _ := climate.ParseAndVerify(os.Args, libclimate.ParseFlag_PanicOnFailure)

	// Program logic

	fmt.Printf("%v value(s):\n", len(r.Values))
	for ix, value := range r.Values {
		fmt.Printf("\t[%d]\t'%s'\n", ix, value.Value)
	}
}
```


## Usage


### No arguments

If executed with no arguments

```bash
go run examples/values.go
```

it gives the output:

```
values: country name not specified
```


### Show usage

If executed with the arguments

```bash
go run examples/values.go --help
```

it gives the output:

```
libCLImate.Go Examples

values 0.0.1

USAGE: values [ ... flags and options ... ] <country-name> <state-name> [ <city-name> [ <district-id> ]]

flags/options:

	--help
		Shows this help and exits

	--version
		Shows version information and exits
```


### Specify one value

If executed with the arguments

```bash
go run examples/values.go England
```

it gives the output:

```
values: state id not specified
```


### Specify two values

If executed with the arguments

```bash
go run examples/values.go England Worcestershire
```

it gives the output:

```
2 value(s):
	[0]	'England'
	[1]	'Worcestershire'
```


### Specify three values

If executed with the arguments

```bash
go run examples/values.go England Worcestershire Evesham
```

it gives the output:

```
3 value(s):
	[0]	'England'
	[1]	'Worcestershire'
	[2]	'Evesham'
```


### Specify four values

If executed with the arguments

```bash
go run examples/values.go England Worcestershire Evesham Bengeworth
```

it gives the output:

```
4 value(s):
	[0]	'England'
	[1]	'Worcestershire'
	[2]	'Evesham'
	[3]	'Bengeworth'
```


### Specify five values

If executed with the arguments

```bash
go run examples/values.go England Worcestershire Evesham Bengeworth Badsley
```

it gives the output:

```
values: too many values
```


<!-- ########################### end of file ########################### -->

