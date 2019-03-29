
package libclimate_test


// capture_exiter (struct)
//
// Does not call Exit(), but does capture given exit-code

type capture_exiter struct {

	ExitCode	int
}

func (se capture_exiter) Exit(exitCode int) {

	se.ExitCode = exitCode
}


// stub_exiter (struct)
//
// Does not call Exit()

type stub_exiter struct {
}

func (se stub_exiter) Exit(exitCode int) {

	// Do nothing
}

