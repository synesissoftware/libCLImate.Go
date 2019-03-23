/* /////////////////////////////////////////////////////////////////////////
 * File:        climate.go
 *
 * Purpose:     Main source file for libCLImate.Go
 *
 * Created:     22nd March 2019
 * Updated:     23rd March 2019
 *
 * Home:        http://synesis.com.au/software
 *
 * Copyright (c) 2019, Matthew Wilson
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 * - Redistributions of source code must retain the above copyright notice,
 *   this list of conditions and the following disclaimer.
 * - Redistributions in binary form must reproduce the above copyright
 *   notice, this list of conditions and the following disclaimer in the
 *   documentation and/or other materials provided with the distribution.
 * - Neither the names of Matthew Wilson and Synesis Software nor the names
 *   of any contributors may be used to endorse or promote products derived
 *   from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS
 * IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO,
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
 * PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR
 * CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 * PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 * LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 * NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 * ////////////////////////////////////////////////////////////////////// */


package libclimate

import (

	clasp "github.com/synesissoftware/CLASP.Go"

	"fmt"
	"os"
)

// Type of flags passed to the libclimate.Init() method
type InitFlag int

// Type of flags passed to the libclimate.Parse() method
type ParseFlag int

// Type of flags passed to the libclimate.AddAlias(), libclimate.AddFlag(),
// and libclimate.AddOption() methods
type AliasFlag int

// Structure representing a CLI parsing context, obtained from
// libclimate.Init()
type Climate struct {

	Aliases		[]clasp.Alias
	ParseFlags	clasp.ParseFlag
	Version		interface{}
	InfoLines	[]string
}

// Structure representing CLI results, obtained from Climate.Parse()
type Result struct {

	Flags		[]*clasp.Argument
	Options		[]*clasp.Argument
	Values		[]*clasp.Argument
	ProgramName	string
	Argv		[]string

	arguments	*clasp.Arguments
}

// Callback function for specification of Climate via DSL
type InitFunc func(cl *Climate) error

const (

	InitFlag_None				InitFlag	=	1 << iota
	InitFlag_PanicOnFailure		InitFlag	=	1 << iota
	InitFlag_NoHelpFlag			InitFlag	=	1 << iota
	InitFlag_NoVersionFlag		InitFlag	=	1 << iota
)

const (

	ParseFlag_None				ParseFlag	=	1 << iota
	ParseFlag_PanicOnFailure	ParseFlag	=	1 << iota
	ParseFlag_DontCheckUnused	ParseFlag	=	1 << iota
)

// Initialises a Climate instance, according to the given function (which
// may not be nil) and arguments
func Init(initFn InitFunc, args ...interface{}) (climate *Climate, err error) {

	climate	=	&Climate{

		Aliases: []clasp.Alias { },
	}

	if true {

		climate.Aliases = append(climate.Aliases, clasp.HelpFlag())
		climate.Aliases = append(climate.Aliases, clasp.VersionFlag())
	}

	err = initFn(climate)

	if err == nil {

	}

	return
}

// Adds an alias to the Climate instance
func (cl *Climate) AddAlias(alias clasp.Alias, flags ...AliasFlag) {

	cl.Aliases = append(cl.Aliases, alias)
}

// Adds a flag to the Climate instance
func (cl *Climate) AddFlag(flag clasp.Alias, flags ...AliasFlag) {

	cl.Aliases = append(cl.Aliases, flag)
}

// Adds an option to the Climate instance
func (cl *Climate) AddOption(option clasp.Alias, flags ...AliasFlag) {

	cl.Aliases = append(cl.Aliases, option)
}

// Parses a command line, obtaining a Result instance representing the
// arguments received by the process
func (cl Climate) Parse(argv []string, args ...interface{}) (result Result, err error) {

	parse_params := clasp.ParseParams {

		Aliases: cl.Aliases,
	}

	arguments := clasp.Parse(argv, parse_params)

	if arguments.FlagIsSpecified(clasp.HelpFlag()) {

		clasp.ShowUsage(cl.Aliases, clasp.UsageParams{

			Version: cl.Version,
			InfoLines: cl.InfoLines,
		})
	}

	if arguments.FlagIsSpecified(clasp.VersionFlag()) {

		clasp.ShowVersion(cl.Aliases, clasp.UsageParams{ Version: cl.Version })
	}

	// Check for any unrecognised flags or options

	result = Result{

		Flags: arguments.Flags,
		Options: arguments.Options,
		Values: arguments.Values,

		ProgramName: arguments.ProgramName,
		Argv: argv,

		arguments: arguments,
	}

	return
}

// Verifies that all given arguments received are recognised according to
// the specified flags and options
func (result Result) Verify(args ...interface{}) {

	if unused := result.arguments.GetUnusedFlagsAndOptions(); 0 != len(unused) {

		fmt.Fprintf(os.Stderr, "%s: unrecognised flag/option: %s\n", result.arguments.ProgramName, unused[0].Str())

		os.Exit(1)
	}
}

// Parses via Parse() and verifies via Verify()
//
// Panics, rather than returns, if the ParseFlag_PanicOnFailure flag is
// specified
func (cl Climate) ParseAndVerify(argv []string, args ...interface{}) (result Result, err error) {

	result, err = cl.Parse(argv, args...)
	if err != nil {

		return
	} else {

		result.Verify(args...)

		return
	}
}

// Determines if the given flag is specified
func (result Result) FlagIsSpecified(id interface{}) bool {

	return result.arguments.FlagIsSpecified(id)
}

// Looks for a flag with the given id - name, or the alias instance - and
// returns it and the value true if found; if not, returns nil and false
func (result Result) LookupFlag(id interface{}) (*clasp.Argument, bool) {

	return result.arguments.LookupFlag(id)
}

// Looks for an option with the given id - name, or the alias instance - and
// returns it and the value true if found; if not, returns nil and false
func (result Result) LookupOption(id interface{}) (*clasp.Argument, bool) {

	return result.arguments.LookupOption(id)
}

/* ///////////////////////////// end of file //////////////////////////// */


