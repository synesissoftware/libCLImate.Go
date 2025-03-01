# libCLImate.Go <!-- omit in toc -->

**C**ommand-**L**ine **I**nterface boilerplate mini-framework, for Go


## Introduction

**libCLImate** is a **C**ommand-**L**ine **I**nterface boilerplate
mini-framework, which encapsulates the common aspects of Command-Line
Interface boilerplate. The first
[libCLImate was a C/C++ library](https://github.com/synesissoftware/libCLImate/).
There have been several implementations in other languages. **libCLImate.Go** is the
**Go** version.


## Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
- [Background](#background)
- [Installation](#installation)
- [Components](#components)
- [Examples](#examples)
- [Project Information](#project-information)
	- [Where to get help](#where-to-get-help)
	- [Contribution guidelines](#contribution-guidelines)
	- [Dependencies](#dependencies)
		- [Development/Testing Dependencies](#developmenttesting-dependencies)
	- [Related projects](#related-projects)
	- [License](#license)


## Background

As described in the [**CLASP**](https://github.com/synesissoftware/CLASP) project, command-line arguments may be discriminated as **flags**, **options**, and **values**:

* **flags** are arguments that begin with 1+ hyphens and have a name, and whose presence/absence alter program behaviour, e.g. ``-f``, ``--verbose``;
* **options** are arguments that begin with 1+ hyphens and have a name *and* a value, e.g. ``--verbosity=terse``, ``-v t``;
* **values** are arguments that do not begin with hyphens (and are not the value of an **option**), e.g. ``~/my-files/``.

**libCLImate** builds upon this discrimination of arguments to provide higher-level facilities and more succinct specification. All variants (except for the [**C/C++** version](https://github.com/synesissoftware/libCLImate/)) provide this in the form of a ``Climate`` class type, which is specified according to a lightweight DSL. The constructed ``Climate`` instance is then used to parse the command-line arguments, and provides a number of facilities, including:

* parsing into **flags**, **options**, **values**, including recognition of special flag ``--`` as marking end of **flags**/**options** (which is done via the underlying **CLASP** library);
* automatic handling of ``--help`` flag, printing out a full description of the program features and all **flags**/**options** in details, and then terminating with exit code 0;
* automatic handling of ``--version`` flag, printing out the program name and version, and then terminating with exit code 0;
* detection and reporting of unrecognised **flags**/**options**, and then terminating with exit code 1;
* sundry other helper functions and parsed information.

**libCLImate.Go**, which is built from [**CLASP.Go**](https://github.com/synesissoftware/CLASP.Go), provides access to an instance of ``Climate`` (a **Go** ``struct``) via the ``Init()`` function's callback function argument, as illustrated in brief in the [Components](#components) section below and extensively in the [EXAMPLES.md](./EXAMPLES.md).


## Installation

Install via `go get`, as in:

```bash
go get "github.com/synesissoftware/libCLImate.Go"
```

and then import as:

```Go
import libclimate "github.com/synesissoftware/libCLImate.Go"
```

or, simply, as:

```Go
import "github.com/synesissoftware/libCLImate.Go"
```


## Components

With **libCLImate.Go**, specification of the ``Climate`` ``struct`` is done via the ``Init()`` function by specifying a callback within which the features are specified, as in:

```Go
// README_Components.go

func main() {

	climate, err := libclimate.Init(func (cl *libclimate.Climate) (error) {

        . . . // specify features HERE

		return nil
	});
	if err != nil {

		fmt.Fprintf(os.Stderr, "failed to create CLI parser: %v\n", err)
	}

	_, _ = climate.ParseAndVerify(os.Args, libclimate.ParseFlag_PanicOnFailure)

    . . . // rest of program
}
```

Inside "constructor" - at ``specify features HERE`` - the various features of the `Climate` instance under construction can be specified, including:

* the **version** as a string or an array (of numbers or strings), as in:

```Go
	cl.Version = "0.0.1"
```

* the **info lines** as an array of strings (including special value ``":version:"``, which prints program name and version), as in:

```Go
	cl.InfoLines = []string{
		"Example program",
		"",
		":version:",
		"",
	}
```

Even with just those two attributes set, the program will now respond to ``--help`` and ``--version`` with useful output:

* ``--help``
```
Example program

README_Components 0.0.1

USAGE: README_Components [ ... flags and options ... ]

flags/options:

	--help
		Shows this help and exits

	--version
		Shows version information and exits
```

* ``--version``
```
README_Components 0.0.1
```

Specification of program-specific **flags**/**options** is straightforward, e.g.

* a **flag** ``--debug``:

```Go
	cl.AddFlag(clasp.Flag("--debug").SetHelp("runs in Debug mode").SetAlias("-d"))
```

* an **option** ``--verbosity``, with a callback function:

```Go
	o_Verbosity := clasp.Option("--verbosity").SetHelp("specifies verbosity").SetAlias("-v").SetValues("terse", "quiet", "silent", "chatty")

	cl.AddOptionFunc(o_Verbosity, func (o *clasp.Argument, a *clasp.Alias) {
		fmt.Printf("verbosity specified as: %v\n", o.Value)
	})
```

Now the program will respond to the flag ``--help`` with:

```
Example program

README_Components 0.0.1

USAGE: README_Components [ ... flags and options ... ]

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
		specifies verbosity
		where <value> one of:
			terse
			quiet
			silent
			chatty
```

And will respond to the option ``-v silent`` (which is equivalent to ``--verbosity=silent``) with:

```
verbosity specified as: silent
```


## Examples

Examples are provided in the ```examples``` directory, along with a markdown description for each. A detailed list TOC of them is provided in [EXAMPLES.md](./EXAMPLES.md).


## Project Information


### Where to get help

[GitHub Page](https://github.com/synesissoftware/libCLImate.Go "GitHub Page")


### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/libCLImate.Go.


### Dependencies

**libCLImate.Go** depends on:

* [**CLASP.Go**](https://github.com/synesissoftware/CLASP.Go/)
* [**ver2go**](https://github.com/synesissoftware/ver2go/)


#### Development/Testing Dependencies

* [**testify**](github.com/stretchr/testify);
* [**ANGoLS**](https://github.com/synesissoftware/ANGoLS/);


### Related projects

* [**libCLImate**](https://github.com/synesissoftware/libCLImate/)
* [**libCLImate.Python**](https://github.com/synesissoftware/libCLImate.Python/)
* [**libCLImate.Ruby**](https://github.com/synesissoftware/libCLImate.Ruby/)


### License

**libCLImate.Go** is released under the 3-clause BSD license. See [LICENSE](./LICENSE) for details.


<!-- ########################### end of file ########################### -->

