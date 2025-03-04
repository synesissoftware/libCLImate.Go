# libCLImate.Go Example - **values**

## Summary

Illustrates the handling of `Values`.


## Source

```Go
```


## Usage


### No arguments

If executed with no arguments

```bash
go run examples/values.go
```

it gives the output:

```
```


### Show usage

If executed with the arguments

```bash
go run examples/values.go --help
```

it gives the output:

```
```


### Specify flags and options in long-form

If executed with the arguments

```bash
```

it gives the output:

```
verbosity is specified as: silent
Debug mode is specified
```


### Specify flags and options in short-form

If executed with the arguments

```bash
go run examples/values.go -v silent -d
```

it gives the (same) output:

```
verbosity is specified as: silent
Debug mode is specified
```


### Specify flags and options in short-form, including an alias for an option-with-value

If executed with the arguments

```bash
go run examples/values.go -c -d
```

it gives the output:

```
verbosity is specified as: chatty
Debug mode is specified
```


### Specify flags and options with combined short-form

If executed with the arguments

```bash
go run examples/values.go -dc
```

it gives the (same) output:

```
verbosity is specified as: chatty
Debug mode is specified
```


<!-- ########################### end of file ########################### -->

