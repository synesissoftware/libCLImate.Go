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

		cl.InfoLines = []string{"libCLImate.Go Examples", "", ":version:", ""}

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
