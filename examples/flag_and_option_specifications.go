// examples/flag_and_option_specifications.go

package main

import (

	clasp "github.com/synesissoftware/CLASP.Go"
	libclimate "github.com/synesissoftware/LibCLImate.Go"

	"fmt"
	"os"
)

func main() {

	// Specify specifications, parse, and checking standard flags

	flag_Debug			:=	clasp.Specification{ clasp.FlagType, "--debug", []string{ "-d" }, "runs in Debug mode", nil, 0, nil }
	option_Verbosity	:=	clasp.Specification{ clasp.OptionType, "--verbosity", []string{ "-v" }, "specifies the verbosity", []string{ "terse", "quiet", "silent", "chatty" }, 0, nil }


	climate, err := libclimate.Init(func (cl *libclimate.Climate) (err error) {

		cl.AddFlag(flag_Debug)
		cl.AddOption(option_Verbosity)
		cl.AddAlias("--verbosity=chatty", "-c")

		cl.Version = "0.0.1"

		cl.InfoLines = []string { "libCLImate.Go Examples", "", ":version:", "" }

		return nil
	});
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

