// libclimate_example.go

package main

import (

	clasp "github.com/synesissoftware/CLASP.Go"
	libclimate "github.com/synesissoftware/libCLImate.Go"

	"fmt"
	"os"
)

func main() {

	climate, err := libclimate.Init(func (cl *libclimate.Climate) (error) {

        // specify features HERE
		cl.Version = "0.0.1"
        cl.InfoLines = []string{ "Example program", "", ":version:", "", }

		cl.AddFlag(clasp.Flag("--debug").SetHelp("runs in Debug mode").SetAlias("-d"))

        o_Verbosity := clasp.Option("--verbosity").SetHelp("specifies verbosity").SetAlias("-v").SetValues("terse", "quiet", "silent", "chatty")

        cl.AddOptionFunc(o_Verbosity, func (o *clasp.Argument, a *clasp.Specification) {

			fmt.Printf("verbosity specified as: %v\n", o.Value)
		})


		return nil
	});
	if err != nil {

		fmt.Fprintf(os.Stderr, "failed to create CLI parser: %v\n", err)
	}

	_, _ = climate.ParseAndVerify(os.Args, libclimate.ParseFlag_PanicOnFailure)

    // rest of program
}
