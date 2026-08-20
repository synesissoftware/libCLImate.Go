package main

import (
	libclimate "github.com/synesissoftware/libCLImate.Go"

	ver2go "github.com/synesissoftware/ver2go"

	"fmt"
)

func main() {
	fmt.Printf("libCLImate.Go v%s\n", libclimate.VersionString())
	fmt.Printf("ver2go v%s\n", ver2go.VersionString())
}
