package libclimate_test

import (
	libclimate "github.com/synesissoftware/libCLImate.Go"
	"github.com/synesissoftware/libCLImate.Go/internal"

	"bytes"
	"errors"
	"fmt"
	"testing"
)

func Test_Abort_1(t *testing.T) {

	climate, _ := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		return nil
	})

	stm := new(bytes.Buffer)
	exiter := new(internal.CaptureExiter)

	msg := "Some-failure-condition"
	err := errors.New("SOME-FAILURE-REASON")

	climate.ProgramName = "myapp"
	climate.Abort(msg, err, stm, exiter)

	actual := stm.String()
	expected := fmt.Sprintf("myapp: %s: %s\n", msg, err)

	if expected != actual {

		t.Errorf("expected '%v' != actual '%v'", expected, actual)
	}
}

func Test_Abort_2(t *testing.T) {

	climate, _ := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		return nil
	})

	stm := new(bytes.Buffer)
	exiter := new(internal.CaptureExiter)

	msg := "Some-failure-condition"

	climate.ProgramName = "myapp"
	climate.Abort(msg, nil, stm, exiter)

	actual := stm.String()
	expected := fmt.Sprintf("myapp: %s\n", msg)

	if expected != actual {

		t.Errorf("expected '%v' != actual '%v'", expected, actual)
	}
}
