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

	flag_Debug := clasp.Specification{clasp.FlagType, "--debug", []string{"-d"}, "runs in Debug mode", nil, 0, nil}
	option_Verbosity := clasp.Specification{clasp.OptionType, "--verbosity", []string{"-v"}, "specifies the verbosity", []string{"terse", "quiet", "silent", "chatty"}, 0, nil}

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

		cl.Version = "0.0.1"

		cl.InfoLines = []string{"libCLImate.Go Examples", "", ":version:", ""}

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
