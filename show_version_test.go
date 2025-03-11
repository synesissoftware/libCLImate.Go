package libclimate_test

import (
	libclimate "github.com/synesissoftware/libCLImate.Go"
	"github.com/synesissoftware/libCLImate.Go/internal"

	"bytes"
	"fmt"
	"os"
	"testing"
)

func Test_ShowVersion_1(t *testing.T) {

	stm := new(bytes.Buffer)
	argv := []string{"bin/myapp", "--version"}

	climate, err := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.Version = "0.0.1"

		return nil
	})
	if err != nil {

		fmt.Fprintf(os.Stderr, "failed to create CLI parser: %v\n", err)
	}

	_, _ = climate.ParseAndVerify(argv, stm, internal.StubExiter{})

	actual := stm.String()
	expected := "myapp 0.0.1\n"

	if expected != actual {

		t.Errorf("expected '%v' != actual '%v'", expected, actual)
	}
}

func Test_ShowVersion_2(t *testing.T) {

	stm := new(bytes.Buffer)
	argv := []string{"bin/myapp", "--version"}

	climate, err := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.Version = []int{0, 1, 2}
		cl.VersionPrefix = "v"

		return nil
	})
	if err != nil {

		fmt.Fprintf(os.Stderr, "failed to create CLI parser: %v\n", err)
	}

	_, _ = climate.ParseAndVerify(argv, stm, internal.StubExiter{})

	actual := stm.String()
	expected := "myapp v0.1.2\n"

	if expected != actual {

		t.Errorf("expected '%v' != actual '%v'", expected, actual)
	}
}

func Test_ShowVersion_2_NoVersionFlag(t *testing.T) {

	stm := new(bytes.Buffer)
	argv := []string{"bin/myapp", "--version"}

	climate, err := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.Version = []int{0, 1, 2}
		cl.VersionPrefix = "v"

		cl.UsageHelpSuffix = ""

		return nil
	}, libclimate.InitFlag_NoVersionFlag)
	if err != nil {

		fmt.Fprintf(os.Stderr, "failed to create CLI parser: %v\n", err)
	}

	_, _ = climate.ParseAndVerify(argv, stm, internal.StubExiter{})

	actual := stm.String()
	expected := "myapp: unrecognised flag/option: --version\n"

	if expected != actual {

		t.Errorf("expected '%v' != actual '%v'", expected, actual)
	}
}
