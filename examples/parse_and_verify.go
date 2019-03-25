// examples/parse_and_verify.go

package main

import (

	clasp "github.com/synesissoftware/CLASP.Go"
	libclimate "github.com/synesissoftware/LibCLImate.Go"

	"fmt"
	"os"
)

func main() {

	// Specify aliases, parse, and checking standard flags

	flag_Debug			:=	clasp.Alias{ clasp.FlagType, "--debug", []string{ "-d" }, "runs in Debug mode", nil, 0, nil }
	option_Verbosity	:=	clasp.Alias{ clasp.OptionType, "--verbosity", []string{ "-v" }, "specifies the verbosity", []string{ "terse", "quiet", "silent", "chatty" }, 0, nil }
	flag_Chatty			:=	clasp.Alias{ clasp.FlagType, "--verbosity=chatty", []string{ "-c" }, "", nil, 0, nil }

	is_debug			:=	false
	verbosity			:=	""

	climate, err := libclimate.Init(func (cl *libclimate.Climate) (err error) {

		cl.AddFlagFunc(flag_Debug, func() {

			is_debug = true
		})
		cl.AddOptionFunc(option_Verbosity, func (o *clasp.Argument, a *clasp.Alias) {

			verbosity = o.Value
		})
		cl.AddAlias(flag_Chatty)

		cl.Version = "0.0.1"

		cl.InfoLines = []string { "libCLImate.Go Examples", "", ":version:", "" }

		return nil
	});
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


