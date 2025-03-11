package libclimate_test

import (
	"github.com/stretchr/testify/require"
	libclimate "github.com/synesissoftware/libCLImate.Go"
	"github.com/synesissoftware/libCLImate.Go/internal"

	"bytes"
	"errors"
	"fmt"
	"testing"
)

func Test_Abort_1(t *testing.T) {

	climate, _ := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.ProgramName = "myapp"

		return nil
	}, libclimate.InitFlag_PanicOnFailure)

	stm := new(bytes.Buffer)
	exiter := new(internal.CaptureExiter)

	msg := "Some-failure-condition"
	err := errors.New("SOME-FAILURE-REASON")

	climate.Abort(msg, err, stm, exiter)

	actual := stm.String()
	expected := fmt.Sprintf("myapp: %s: %s\n", msg, err)

	require.Equal(t, expected, actual)
}

func Test_Abort_2(t *testing.T) {

	climate, _ := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.ProgramName = "myapp"

		return nil
	}, libclimate.InitFlag_PanicOnFailure)

	stm := new(bytes.Buffer)
	exiter := new(internal.CaptureExiter)

	msg := "Some-failure-condition"

	climate.Abort(msg, nil, stm, exiter)

	actual := stm.String()
	expected := fmt.Sprintf("myapp: %s\n", msg)

	require.Equal(t, expected, actual)
}

func Test_Abort_WITH_UsageHelpSuffix_1(t *testing.T) {

	climate, _ := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.ProgramName = "myapp"
		cl.UsageHelpSuffix = "specify --help for usage"

		return nil
	}, libclimate.InitFlag_PanicOnFailure)

	stm := new(bytes.Buffer)
	exiter := new(internal.CaptureExiter)

	msg := "Some-failure-condition"

	climate.Abort(msg, nil, stm, exiter)

	actual := stm.String()
	expected := fmt.Sprintf("myapp: %s; specify --help for usage\n", msg)

	require.Equal(t, expected, actual)
}

func Test_Abort_WITH_UsageHelpSuffix_2(t *testing.T) {

	climate, _ := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.ProgramName = "myapp"
		cl.UsageHelpSuffix = ":"

		return nil
	}, libclimate.InitFlag_PanicOnFailure)

	stm := new(bytes.Buffer)
	exiter := new(internal.CaptureExiter)

	msg := "Some-failure-condition"

	climate.Abort(msg, nil, stm, exiter)

	actual := stm.String()
	expected := fmt.Sprintf("myapp: %s; use --help for usage\n", msg)

	require.Equal(t, expected, actual)
}

func Test_Abort_WITH_UsageHelpSuffix_3(t *testing.T) {

	climate, _ := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.ProgramName = "myapp"
		cl.UsageHelpSuffix = ", specify --help for usage"

		return nil
	}, libclimate.InitFlag_PanicOnFailure)

	stm := new(bytes.Buffer)
	exiter := new(internal.CaptureExiter)

	msg := "Some-failure-condition"

	climate.Abort(msg, nil, stm, exiter)

	actual := stm.String()
	expected := fmt.Sprintf("myapp: %s, specify --help for usage\n", msg)

	require.Equal(t, expected, actual)
}
