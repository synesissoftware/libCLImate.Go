// Copyright 2019-2025, Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 30th March 2019
 * Updated: 4th March 2025
 */

package internal

import "os"

// Defines if/how a process should be exited.
type Exiter interface {
	Exit(exitCode int)
}

// An [Exiter] that performs the default action of exiting the process (via
// a call to [os.Exit]).
type DefaultExiter struct {
}

func (DefaultExiter) Exit(exitCode int) {

	os.Exit(exitCode)
}

// An [Exiter] that does not cause process exit, rather it captures the
// exit-code for future reference. This type is intended to serve as a test
// double.
type CaptureExiter struct {
	ExitCode int
}

func (se *CaptureExiter) Exit(exitCode int) {

	se.ExitCode = exitCode
}

// An [Exiter] that does not cause process exit. This type is intended to
// serve as a test double.
type StubExiter struct {
}

func (StubExiter) Exit(exitCode int) {

	// Do nothing
}
