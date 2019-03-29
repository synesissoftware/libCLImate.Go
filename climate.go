/* /////////////////////////////////////////////////////////////////////////
 * File:        climate.go
 *
 * Purpose:     Main source file for libCLImate.Go
 *
 * Created:     22nd March 2019
 * Updated:     29th March 2019
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
	"io"
	"os"
	"path"
)

// Type of flags passed to the libclimate.Init() method
type InitFlag int

// Type of flags passed to the libclimate.Parse() method
type ParseFlag int

// Type of flags passed to the libclimate.AddAlias(), libclimate.AddFlag(),
// and libclimate.AddOption() methods
type AliasFlag int

type exiter interface {

	Exit(exitCode int)
}

// Structure representing a CLI parsing context, obtained from
// libclimate.Init()
type Climate struct {

	Aliases			[]*clasp.Alias
	ParseFlags		clasp.ParseFlag
	Version			interface{}
	VersionPrefix	string
	InfoLines		[]string
	ProgramName		string

	initFlags_		InitFlag
	stream_			io.Writer
	exiter_			exiter
}

// Structure representing CLI results, obtained from Climate.Parse()
type Result struct {

	Flags		[]*clasp.Argument
	Options		[]*clasp.Argument
	Values		[]*clasp.Argument
	ProgramName	string
	Argv		[]string

	arguments_	*clasp.Arguments
	parseFlags_	ParseFlag
	stream_		io.Writer
	exiter_		exiter
}

// Callback function for specification of Climate via DSL
type InitFunc func(cl *Climate) error

type FlagFunc func()

type OptionFunc func(option *clasp.Argument, alias *clasp.Alias)

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

const (

	_libCLImate_FlagFunc	=	"_libCLImate_FlagFunc_F73BB1C0_92D7_4cd5_9C36_DB672290CBE7"
	_libCLImate_OptionFunc	=	"_libCLImate_OptionFunc_F73BB1C0_92D7_4cd5_9C36_DB672290CBE7"
)

/* /////////////////////////////////////////////////////////////////////////
 * helper functions
 */

func parse_Exiter_from_options_(options ...interface{}) (result exiter, err error) {

	for _, option := range(options) {

		switch v := option.(type) {

		case exiter:

			return v, nil
		}
	}

	return
}

func parse_Stream_from_options_(options ...interface{}) (result io.Writer, err error) {

	for _, option := range(options) {

		switch v := option.(type) {

		case io.Writer:

			return v, nil
		}
	}

	return
}

func parse_InitFlags_from_options_(options ...interface{}) (result InitFlag, err error) {

	for _, option := range(options) {

		switch v := option.(type) {

		case InitFlag:

			result |= v
		}
	}

	return
}

func parse_ParseFlags_from_options_(options ...interface{}) (result ParseFlag, err error) {

	for _, option := range(options) {

		switch v := option.(type) {

		case ParseFlag:

			result |= v
		}
	}

	return
}

func pointer_aliases_to_value_aliases(input []*clasp.Alias) (result []clasp.Alias) {

	result	=	make([]clasp.Alias, len(input))

	for i, a := range(input) {

		result[i] = *a
	}

	return
}

/* /////////////////////////////////////////////////////////////////////////
 * API functions
 */

// Initialises a Climate instance, according to the given function (which
// may not be nil) and arguments
func Init(initFn InitFunc, options ...interface{}) (climate *Climate, err error) {

	initFlags, err := parse_InitFlags_from_options_(options...)

	if err != nil {

		return
	}

	climate	=	&Climate{

		Aliases:		[]*clasp.Alias { },
		initFlags_:		initFlags,
		ProgramName: 	path.Base(os.Args[0]),
	}

	if 0 == (initFlags & InitFlag_NoHelpFlag) {

		climate.AddFlag(clasp.HelpFlag())
	}

	if 0 == (initFlags & InitFlag_NoVersionFlag) {

		climate.AddFlag(clasp.VersionFlag())
	}

	err = initFn(climate)

	if err != nil {

	}

	return
}

// Adds a (copy of the) alias to the Climate instance
func (cl *Climate) AddAlias(alias clasp.Alias, flags ...AliasFlag) {

	cl.Aliases = append(cl.Aliases, &alias)
}

// Adds a (copy of the) flag to the Climate instance
func (cl *Climate) AddFlag(flag clasp.Alias, flags ...AliasFlag) {

	cl.Aliases = append(cl.Aliases, &flag)
}

// Adds a (copy of the) flag to the Climate instance
func (cl *Climate) AddFlagFunc(flag clasp.Alias, flagFn FlagFunc, flags ...AliasFlag) {

	newFlag := flag.SetExtra(_libCLImate_FlagFunc, flagFn)

	cl.Aliases = append(cl.Aliases, &newFlag)
}

// Adds a (copy of the) option to the Climate instance
func (cl *Climate) AddOption(option clasp.Alias, flags ...AliasFlag) {

	cl.Aliases = append(cl.Aliases, &option)
}

// Adds a (copy of the) option to the Climate instance
func (cl *Climate) AddOptionFunc(option clasp.Alias, optionFn OptionFunc, flags ...AliasFlag) {

	newOption := option.SetExtra(_libCLImate_OptionFunc, optionFn)

	cl.Aliases = append(cl.Aliases, &newOption)
}

// Parses a command line, obtaining a Result instance representing the
// arguments received by the process
func (cl Climate) Parse(argv []string, options ...interface{}) (result Result, err error) {

	parseFlags, err := parse_ParseFlags_from_options_(options...)

	if err != nil {

		return
	}

	stream, err := parse_Stream_from_options_(options...)
	if err != nil {

		return
	}

	exiter, err := parse_Exiter_from_options_(options...)
	if err != nil {

		return
	}

	parse_params := clasp.ParseParams {

		Aliases: pointer_aliases_to_value_aliases(cl.Aliases),
	}

	arguments := clasp.Parse(argv, parse_params)

	if arguments.FlagIsSpecified(clasp.HelpFlag()) {

		clasp.ShowUsage(parse_params.Aliases, clasp.UsageParams{

			Version: cl.Version,
			VersionPrefix: cl.VersionPrefix,
			InfoLines: cl.InfoLines,
			Stream: stream,
			Exiter: exiter,
			ProgramName: arguments.ProgramName,
		})
	}

	if arguments.FlagIsSpecified(clasp.VersionFlag()) {

		clasp.ShowVersion(parse_params.Aliases, clasp.UsageParams{

			Version: cl.Version,
			VersionPrefix: cl.VersionPrefix,
			Stream: stream,
			Exiter: exiter,
			ProgramName: arguments.ProgramName,
		})
	}

	for i := 0; i != len(arguments.Arguments); i++ {

		var argument *clasp.Argument = arguments.Arguments[i]
		var alias *clasp.Alias = argument.ArgumentAlias

		if alias != nil {

			if 0 != len(alias.Extras) {

				if ff, ff_ok := alias.Extras[_libCLImate_FlagFunc]; ff_ok {

					switch fn := ff.(type) {

					case FlagFunc:

						fn();

						argument.Use()
					default:

						// Issue warning
					}
				}

				if of, of_ok := alias.Extras[_libCLImate_OptionFunc]; of_ok {

					switch fn := of.(type) {

					case OptionFunc:

						fn(argument, alias);

						argument.Use()
					default:

						// Issue warning
					}
				}
			}
		}
	}

	result = Result{

		Flags: arguments.Flags,
		Options: arguments.Options,
		Values: arguments.Values,

		ProgramName: arguments.ProgramName,
		Argv: argv,

		arguments_: arguments,
		parseFlags_: parseFlags,
		stream_: stream,
		exiter_: exiter,
	}

	return
}

// Verifies that all given arguments received are recognised according to
// the specified flags and options
func (result Result) Verify(options ...interface{}) {

	// Check for any unrecognised flags or options

	if unused := result.arguments_.GetUnusedFlagsAndOptions(); 0 != len(unused) {

		fmt.Fprintf(os.Stderr, "%s: unrecognised flag/option: %s\n", result.arguments_.ProgramName, unused[0].Str())

		os.Exit(1)
	}
}

// Parses via Parse() and verifies via Verify()
//
// Panics, rather than returns, if the ParseFlag_PanicOnFailure flag is
// specified
func (cl Climate) ParseAndVerify(argv []string, options ...interface{}) (result Result, err error) {

	result, err = cl.Parse(argv, options...)
	if err != nil {

		return
	} else {

		result.Verify(options...)

		return
	}
}

// Emits the given message and, optionally, err to the standard error
// stream, prefixed with the program name, and then terminates the process
// with a non-0 exit code.
func (cl Climate) Abort(message string, err error) {

	if err != nil {

		fmt.Fprintf(os.Stderr, "%s: %s: %v\n", cl.ProgramName, message, err)
	} else {

		fmt.Fprintf(os.Stderr, "%s: %s\n", cl.ProgramName, message)
	}

	os.Exit(1)
}

// Determines if the given flag is specified
func (result Result) FlagIsSpecified(id interface{}) bool {

	return result.arguments_.FlagIsSpecified(id)
}

// Looks for a flag with the given id - name, or the alias instance - and
// returns it and the value true if found; if not, returns nil and false
func (result Result) LookupFlag(id interface{}) (*clasp.Argument, bool) {

	return result.arguments_.LookupFlag(id)
}

// Looks for an option with the given id - name, or the alias instance - and
// returns it and the value true if found; if not, returns nil and false
func (result Result) LookupOption(id interface{}) (*clasp.Argument, bool) {

	return result.arguments_.LookupOption(id)
}

/* ///////////////////////////// end of file //////////////////////////// */


