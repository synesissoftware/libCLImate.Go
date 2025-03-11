package libclimate_test

import (
	angols_slices "github.com/synesissoftware/ANGoLS/slices"
	clasp "github.com/synesissoftware/CLASP.Go"
	libclimate "github.com/synesissoftware/libCLImate.Go"
	"github.com/synesissoftware/libCLImate.Go/internal"

	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func lines_to_display_string_(lines []string) string {

	stm := new(bytes.Buffer)

	for i, line := range lines {

		fmt.Fprintf(stm, "\t[% 2d]\t%s\n", i, line)
	}

	return stm.String()
}

func Test_ShowUsage_1(t *testing.T) {

	stm := new(bytes.Buffer)
	argv := []string{"bin/myapp", "--help"}

	climate, err := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		return nil
	})
	if err != nil {

		fmt.Fprintf(os.Stderr, "failed to create CLI parser: %v\n", err)
	}

	_, _ = climate.ParseAndVerify(argv, stm, internal.StubExiter{})

	actual := strings.Split(stm.String(), "\n")
	expected := []string{

		"USAGE: myapp [ ... flags and options ... ]",
		"flags/options:",
		"\t--help",
		"\t\tShows this help and exits",
		"\t--version",
		"\t\tShows version information and exits",
	}

	actual, _ = angols_slices.SelectSliceOfString(actual, func(_ int, line string) (bool, error) {

		return 0 != len(line), nil
	})

	if !angols_slices.EqualSliceOfString(expected, actual) {

		t.Errorf("expected \n'%v'\n != actual \n'%v'",
			lines_to_display_string_(expected),
			lines_to_display_string_(actual),
		)
	}
}

func Test_ShowUsage_2(t *testing.T) {

	stm := new(bytes.Buffer)
	argv := []string{"bin/myapp", "--help"}

	climate, err := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.Version = []int{0, 1, 2}
		cl.VersionPrefix = "v"

		cl.InfoLines = []string{

			"ShowUsage tests",
			":version:",
		}

		return nil
	})
	if err != nil {

		fmt.Fprintf(os.Stderr, "failed to create CLI parser: %v\n", err)
	}

	_, _ = climate.ParseAndVerify(argv, stm, internal.StubExiter{})

	actual := strings.Split(stm.String(), "\n")
	expected := []string{

		"ShowUsage tests",
		"myapp v0.1.2",
		"USAGE: myapp [ ... flags and options ... ]",
		"flags/options:",
		"\t--help",
		"\t\tShows this help and exits",
		"\t--version",
		"\t\tShows version information and exits",
	}

	actual, _ = angols_slices.SelectSliceOfString(actual, func(_ int, line string) (bool, error) {

		return 0 != len(line), nil
	})

	if !angols_slices.EqualSliceOfString(expected, actual) {

		t.Errorf("expected \n'%v'\n != actual \n'%v'",
			lines_to_display_string_(expected),
			lines_to_display_string_(actual),
		)
	}
}

func Test_ShowUsage_3(t *testing.T) {

	stm := new(bytes.Buffer)
	argv := []string{"bin/myapp", "--help"}

	climate, err := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.Version = []int{0, 1, 2}
		cl.VersionPrefix = "v"

		cl.InfoLines = []string{

			"ShowUsage tests",
			":version:",
		}
		cl.ValuesString = "<path-1> <path-2>"

		cl.AddFlag(clasp.Option("--verbosity").
			SetHelp("Specifies the verbosity").
			SetValues("low", "medium", "high"))
		cl.AddAlias("--verbosity=high", "-v")

		return nil
	})
	if err != nil {

		fmt.Fprintf(os.Stderr, "failed to create CLI parser: %v\n", err)
	}

	_, _ = climate.ParseAndVerify(argv, stm, internal.StubExiter{})

	actual := strings.Split(stm.String(), "\n")
	expected := []string{

		"ShowUsage tests",
		"myapp v0.1.2",
		"USAGE: myapp [ ... flags and options ... ] <path-1> <path-2>",
		"flags/options:",
		"\t--help",
		"\t\tShows this help and exits",
		"\t--version",
		"\t\tShows version information and exits",
		"\t-v --verbosity=high",
		"\t--verbosity=<value>",
		"\t\tSpecifies the verbosity",
		"\t\twhere <value> one of:",
		"\t\t\tlow",
		"\t\t\tmedium",
		"\t\t\thigh",
	}

	actual, _ = angols_slices.SelectSliceOfString(actual, func(_ int, line string) (bool, error) {

		return 0 != len(line), nil
	})

	if !angols_slices.EqualSliceOfString(expected, actual) {

		t.Errorf("expected \n'%v'\n != actual \n'%v'",
			lines_to_display_string_(expected),
			lines_to_display_string_(actual),
		)
	}
}

func Test_ShowUsage_3_NoHelpFlag(t *testing.T) {

	stm := new(bytes.Buffer)
	argv := []string{"bin/myapp", "--help"}

	climate, err := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.Version = []int{0, 1, 2}
		cl.VersionPrefix = "v"

		cl.InfoLines = []string{

			"ShowUsage tests",
			":version:",
		}
		cl.ValuesString = "<path-1> <path-2>"

		cl.UsageHelpSuffix = ""

		cl.AddFlag(clasp.Option("--verbosity").
			SetHelp("Specifies the verbosity").
			SetValues("low", "medium", "high"))
		cl.AddAlias("--verbosity=high", "-v")

		return nil
	}, libclimate.InitFlag_NoHelpFlag)
	if err != nil {

		fmt.Fprintf(os.Stderr, "failed to create CLI parser: %v\n", err)
	}

	_, _ = climate.ParseAndVerify(argv, stm, internal.StubExiter{})

	actual := stm.String()
	expected := "myapp: unrecognised flag/option: --help\n"

	if expected != actual {

		t.Errorf("expected '%v' != actual '%v'", expected, actual)
	}
}

func Test_ShowUsage_3_NoHelpFlag_AND_DontCheckUnused(t *testing.T) {

	stm := new(bytes.Buffer)
	argv := []string{"bin/myapp", "--help"}

	climate, err := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.Version = []int{0, 1, 2}
		cl.VersionPrefix = "v"

		cl.InfoLines = []string{

			"ShowUsage tests",
			":version:",
		}
		cl.ValuesString = "<path-1> <path-2>"

		cl.AddFlag(clasp.Option("--verbosity").
			SetHelp("Specifies the verbosity").
			SetValues("low", "medium", "high"))
		cl.AddAlias("--verbosity=high", "-v")

		return nil
	}, libclimate.InitFlag_NoHelpFlag)
	if err != nil {

		fmt.Fprintf(os.Stderr, "failed to create CLI parser: %v\n", err)
	}

	_, _ = climate.ParseAndVerify(argv, stm, internal.StubExiter{}, libclimate.ParseFlag_DontCheckUnused)

	actual := stm.String()
	expected := ""

	if expected != actual {

		t.Errorf("expected '%v' != actual '%v'", expected, actual)
	}
}

func Test_ShowUsage_3_NoVersionFlag(t *testing.T) {

	stm := new(bytes.Buffer)
	argv := []string{"bin/myapp", "--help"}

	climate, err := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.Version = []int{0, 1, 2}
		cl.VersionPrefix = "v"

		cl.InfoLines = []string{

			"ShowUsage tests",
			":version:",
		}
		cl.ValuesString = "<path-1> <path-2>"

		cl.AddFlag(clasp.Option("--verbosity").
			SetHelp("Specifies the verbosity").
			SetValues("low", "medium", "high"))
		cl.AddAlias("--verbosity=high", "-v")

		return nil
	}, libclimate.InitFlag_NoVersionFlag)
	if err != nil {

		fmt.Fprintf(os.Stderr, "failed to create CLI parser: %v\n", err)
	}

	_, _ = climate.ParseAndVerify(argv, stm, internal.StubExiter{})

	actual := strings.Split(stm.String(), "\n")
	expected := []string{

		"ShowUsage tests",
		"myapp v0.1.2",
		"USAGE: myapp [ ... flags and options ... ] <path-1> <path-2>",
		"flags/options:",
		"\t--help",
		"\t\tShows this help and exits",
		"\t-v --verbosity=high",
		"\t--verbosity=<value>",
		"\t\tSpecifies the verbosity",
		"\t\twhere <value> one of:",
		"\t\t\tlow",
		"\t\t\tmedium",
		"\t\t\thigh",
	}

	actual, _ = angols_slices.SelectSliceOfString(actual, func(_ int, line string) (bool, error) {

		return 0 != len(line), nil
	})

	if !angols_slices.EqualSliceOfString(expected, actual) {

		t.Errorf("expected \n'%v'\n != actual \n'%v'",
			lines_to_display_string_(expected),
			lines_to_display_string_(actual),
		)
	}
}
