// examples/show_usage_and_version.go

package main

import (
	libclimate "github.com/synesissoftware/libCLImate.Go"

	"fmt"
	"os"
)

func main() {

	// Specify specifications, parse, and checking standard flags

	climate, err := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.Version = "0.0.1"

		cl.InfoLines = []string{"libCLImate.Go Examples", "", ":version:", ""}

		return nil
	})
	if err != nil {

		fmt.Fprintf(os.Stderr, "failed to create CLI parser: %v\n", err)
	}

	_, _ = climate.ParseAndVerify(os.Args, libclimate.ParseFlag_PanicOnFailure)

	// Finish normal processing

	fmt.Printf("no flags specified\n")
}
