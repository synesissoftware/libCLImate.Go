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
