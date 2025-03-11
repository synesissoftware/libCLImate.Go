// Copyright 2019-2025, Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 22nd March 2019
 * Updated: 7th March 2025
 */

package libclimate

import (
	clasp "github.com/synesissoftware/CLASP.Go"
	"github.com/synesissoftware/libCLImate.Go/internal"

	"fmt"
	"io"
	"os"
	"path"
	"unicode"
)

// Type of flags passed to the [Init] method.
type InitFlag int64

// Type of flags passed to the [Climate.Parse] and [Climate.ParseAndVerify] methods.
type ParseFlag int64

// Type of flags passed to the [Climate.AddFlag] and [Climate.AddOption] methods.
type AliasFlag int64

// Structure representing a CLI parsing context, obtained from [Init].
type Climate struct {
	Specifications   []*clasp.Specification // The specifications created by [Init].
	ParseFlags       clasp.ParseFlag        // Parsing flags.
	Version          interface{}            // Version field that can be specified by application code in the function called by [Init].
	VersionPrefix    string                 // Version-prefix field that can be specified by application code in the function called by [Init].
	InfoLines        []string               // Information lines field that can be specified by application code in the function called by [Init].
	ValuesString     string                 // Values-string field that can be specified by application code in the function called by [Init].
	ProgramName      string                 // Program-name field that can be specified by application code in the function called by [Init]. Defaults to `os.Args[0]`.
	ValueNames       []string               // Specifies a list of value names that may be used in a contingent report when insufficient values are specified on the command-line (as determined by [Climate.ValuesConstraint]).
	ValuesConstraint []int                  // An array of 1 or 2 numbers that specify the number of values, or the minimum and maximum number of values, required. A value of -1 means "no constraint", so, for example, the constraint `{2, -1}` means 2+ values are required.
	UsageHelpSuffix  string                 // An optional string to be applied to the end of the contingent report produced by [Climate.Abort]. Defaults to nothing. Specify ":" for default suffix string of "; use --help for usage". Insert leading "; " unless first character is punctuation.

	initFlags InitFlag
	stream    io.Writer
	exiter    internal.Exiter
}

// Structure representing CLI results, obtained from [Climate.Parse].
type Result struct {
	Flags       []*clasp.Argument // Array of all flags.
	Options     []*clasp.Argument // Array of all options.
	Values      []*clasp.Argument // Array of all values.
	ProgramName string            // The program name inferred by [Init], which may be overridden in the function called by [Init].
	Argv        []string          // The original argument string array passed to [Parse].

	arguments        *clasp.Arguments
	parseFlags       ParseFlag
	stream           io.Writer
	exiter           internal.Exiter
	valueNames       []string
	valuesConstraint []int
}

// Callback function for specification of Climate via DSL.
type InitFunc func(cl *Climate) error

// Type of callback function that may be specified to [Climate.AddFlagFunc].
type FlagFunc func()

// Type of callback function that may be specified to
// [Climate.AddOptionFunc], which receives the argument and its
// specification.
type OptionFunc func(option *clasp.Argument, specification *clasp.Specification)

const (
	InitFlag_None InitFlag = 0 // No initialisation flags specified.
)

const (
	InitFlag_PanicOnFailure InitFlag = 1 << iota // Causes [Init] to panic if an error encountered during processing.
	InitFlag_NoHelpFlag                          // Suppresses the provision and processing of a help flag (aka "--help").
	InitFlag_NoVersionFlag                       // Suppresses the provision and processing of a version flag (aka "--version").
)

const (
	ParseFlag_None ParseFlag = 0 // No parse flags specified.
)

const (
	ParseFlag_PanicOnFailure  ParseFlag = 1 << iota // Causes [Climate.Parse] to panic if an error encountered during processing.
	ParseFlag_DontCheckUnused                       // Causes [Climate.Verify] to ignore unrecognised arguments.
)

const (
	_libCLImate_FlagFunc   = "_libCLImate_FlagFunc_F73BB1C0_92D7_4cd5_9C36_DB672290CBE7"
	_libCLImate_OptionFunc = "_libCLImate_OptionFunc_F73BB1C0_92D7_4cd5_9C36_DB672290CBE7"
)

const (
	UsageHelpSuffix_Default = "; use --help for usage"
)

/* /////////////////////////////////////////////////////////////////////////
 * helper functions
 */

func parse_Exiter_from_options_(options ...interface{}) (result internal.Exiter, err error) {

	for _, option := range options {

		switch v := option.(type) {

		case internal.Exiter:

			return v, nil
		}
	}

	return
}

func parse_Stream_from_options_(options ...interface{}) (result io.Writer, err error) {

	for _, option := range options {

		switch v := option.(type) {

		case io.Writer:

			return v, nil
		}
	}

	return
}

func parse_InitFlags_from_options_(options ...interface{}) (result InitFlag, err error) {

	for _, option := range options {

		switch v := option.(type) {

		case InitFlag:

			result |= v
		}
	}

	return
}

func parse_ParseFlags_from_options_(options ...interface{}) (result ParseFlag, err error) {

	for _, option := range options {

		switch v := option.(type) {

		case ParseFlag:

			result |= v
		}
	}

	return
}

func pointer_specifications_to_value_specifications(input []*clasp.Specification) (result []clasp.Specification) {

	result = make([]clasp.Specification, len(input))

	for i, a := range input {

		result[i] = *a
	}

	return
}

/* /////////////////////////////////////////////////////////////////////////
 * API functions
 */

// Initialises a Climate instance, according to the given function (which
// may not be nil) and arguments.
func Init(initFn InitFunc, options ...interface{}) (climate *Climate, err error) {

	var initFlags InitFlag
	var stream io.Writer
	var exiter internal.Exiter

	if err == nil {

		initFlags, err = parse_InitFlags_from_options_(options...)
	}

	if err == nil {

		stream, err = parse_Stream_from_options_(options...)
	}

	if err == nil {

		exiter, err = parse_Exiter_from_options_(options...)
	}
	if err == nil && exiter == nil {

		exiter = new(internal.DefaultExiter)
	}

	if err == nil {

		var program_name string
		if 0 != len(os.Args[0]) {
			program_name = path.Base(os.Args[0])
		}

		climate = &Climate{

			Specifications: []*clasp.Specification{},
			//ParseFlags:
			//Version:
			//VersionPrefix:
			//InfoLines:
			// ValuesString:
			ProgramName: program_name,
			// ValueNames:
			// ValuesConstraint:

			initFlags: initFlags,
			stream:    stream,
			exiter:    exiter,
		}

		if 0 == (initFlags & InitFlag_NoHelpFlag) {

			climate.AddFlag(clasp.HelpFlag())
		}

		if 0 == (initFlags & InitFlag_NoVersionFlag) {

			climate.AddFlag(clasp.VersionFlag())
		}

		err = initFn(climate)
	}

	if err != nil {

		if 0 != (InitFlag_PanicOnFailure & initFlags) {

			panic(err)
		}
	}

	return
}

// Adds an alias to the Climate instance
//
// The resolved_name param can be the name of a flag or option, or an
// option-with-value. The alias param is the alias (which must not
// contain an equals sign.
func (cl *Climate) AddAlias(resolved_name, alias string) {

	f := clasp.Flag(resolved_name).SetAlias(alias)

	cl.Specifications = append(cl.Specifications, &f)
}

// Adds a (copy of the) flag to the Climate instance.
func (cl *Climate) AddFlag(flag clasp.Specification, flags ...AliasFlag) {

	cl.Specifications = append(cl.Specifications, &flag)
}

// Adds a (copy of the) flag to the Climate instance.
func (cl *Climate) AddFlagFunc(flag clasp.Specification, flagFn FlagFunc, flags ...AliasFlag) {

	newFlag := flag.SetExtra(_libCLImate_FlagFunc, flagFn)

	cl.Specifications = append(cl.Specifications, &newFlag)
}

// Adds a (copy of the) option to the Climate instance.
func (cl *Climate) AddOption(option clasp.Specification, flags ...AliasFlag) {

	cl.Specifications = append(cl.Specifications, &option)
}

// Adds a (copy of the) option to the Climate instance.
func (cl *Climate) AddOptionFunc(option clasp.Specification, optionFn OptionFunc, flags ...AliasFlag) {

	newOption := option.SetExtra(_libCLImate_OptionFunc, optionFn)

	cl.Specifications = append(cl.Specifications, &newOption)
}

// Parses a command line, obtaining a Result instance representing the
// arguments received by the process.
func (cl Climate) Parse(argv []string, options ...interface{}) (result Result, err error) {

	var parseFlags ParseFlag
	var stream io.Writer
	var exiter internal.Exiter
	var arguments *clasp.Arguments

	if err == nil {

		parseFlags, err = parse_ParseFlags_from_options_(options...)
	}

	if err == nil {

		stream, err = parse_Stream_from_options_(options...)
	}
	if err == nil && stream == nil {

		if cl.stream != nil {

			stream = cl.stream
		} else {

			stream = os.Stderr
		}
	}

	if err == nil {

		exiter, err = parse_Exiter_from_options_(options...)
	}
	if err == nil && exiter == nil {

		exiter = cl.exiter
	}

	if err == nil {

		parse_params := clasp.ParseParams{

			Specifications: pointer_specifications_to_value_specifications(cl.Specifications),
		}

		arguments = clasp.Parse(argv, parse_params)

		if 0 == (cl.initFlags & InitFlag_NoHelpFlag) {

			if arguments.FlagIsSpecified(clasp.HelpFlag()) {

				clasp.ShowUsage(parse_params.Specifications, clasp.UsageParams{

					Version:       cl.Version,
					VersionPrefix: cl.VersionPrefix,
					InfoLines:     cl.InfoLines,
					ValuesString:  cl.ValuesString,
					Stream:        stream,
					Exiter:        exiter,
					ProgramName:   arguments.ProgramName,
				})
			}
		}

		if 0 == (cl.initFlags & InitFlag_NoVersionFlag) {

			if arguments.FlagIsSpecified(clasp.VersionFlag()) {

				clasp.ShowVersion(parse_params.Specifications, clasp.UsageParams{

					Version:       cl.Version,
					VersionPrefix: cl.VersionPrefix,
					Stream:        stream,
					Exiter:        exiter,
					ProgramName:   arguments.ProgramName,
				})
			}
		}

		for i := 0; i != len(arguments.Arguments); i++ {

			var argument *clasp.Argument = arguments.Arguments[i]
			var alias *clasp.Specification = argument.ArgumentSpecification

			if alias != nil {

				if 0 != len(alias.Extras) {

					if ff, ff_ok := alias.Extras[_libCLImate_FlagFunc]; ff_ok {

						switch fn := ff.(type) {

						case FlagFunc:

							fn()

							argument.Use()
						default:

							// Issue warning
						}
					}

					if of, of_ok := alias.Extras[_libCLImate_OptionFunc]; of_ok {

						switch fn := of.(type) {

						case OptionFunc:

							fn(argument, alias)

							argument.Use()
						default:

							// Issue warning
						}
					}
				}
			}
		}
	}

	if err != nil {

		if 0 != (ParseFlag_PanicOnFailure & parseFlags) {

			panic(err)
		}
	} else {

		result = Result{

			Flags:   arguments.Flags,
			Options: arguments.Options,
			Values:  arguments.Values,

			ProgramName: arguments.ProgramName,
			Argv:        argv,

			arguments:        arguments,
			parseFlags:       parseFlags,
			stream:           stream,
			exiter:           exiter,
			valueNames:       cl.ValueNames,
			valuesConstraint: cl.ValuesConstraint,
		}
	}

	return
}

func (result Result) validateValues1(stream io.Writer, constraint int) {

	n := len(result.Values)

	if constraint < 0 {
		// do not validate
	} else {
		if constraint < n {
			fmt.Fprintf(result.stream, "%s: too many values\n", result.ProgramName)

			result.exiter.Exit(1)
		}
		if constraint > n {
			var value_name string
			if n < len(result.valueNames) {
				value_name = result.valueNames[n]
			} else {
				value_name = fmt.Sprintf("value-%d", n)
			}

			fmt.Fprintf(result.stream, "%s: %s not specified\n", result.ProgramName, value_name)

			result.exiter.Exit(1)
		}
	}
}

func (result Result) validateValues2(stream io.Writer, min, max int) {

	if min == max {
		result.validateValues1(stream, min)
	} else {

		n := len(result.Values)

		if max > 0 && max < n {
			fmt.Fprintf(result.stream, "%s: too many values\n", result.ProgramName)

			result.exiter.Exit(1)
		}

		if min > 0 && min > n {
			var value_name string
			if n < len(result.valueNames) {
				value_name = result.valueNames[n]
			} else {
				value_name = fmt.Sprintf("value-%d", n)
			}

			fmt.Fprintf(result.stream, "%s: %s not specified\n", result.ProgramName, value_name)

			result.exiter.Exit(1)
		}
	}
}

// Verifies that all given arguments received are recognised according to
// the specified flags and options
func (result Result) Verify(options ...interface{}) {

	var err error
	var parseFlags ParseFlag

	stream, _ := parse_Stream_from_options_(options...)
	if stream == nil {

		stream = result.stream
	}
	if stream == nil {

		stream = os.Stderr
	}

	if err == nil {

		parseFlags, err = parse_ParseFlags_from_options_(options...)
	}
	parseFlags |= result.parseFlags

	if 0 == (ParseFlag_DontCheckUnused & parseFlags) {

		// Check for any unrecognised flags or options

		if unused := result.arguments.GetUnusedFlagsAndOptions(); 0 != len(unused) {

			fmt.Fprintf(stream, "%s: unrecognised flag/option: %s\n", result.arguments.ProgramName, unused[0].Str())

			result.exiter.Exit(1)
		}
	}

	switch len(result.valuesConstraint) {
	case 0:
		// do not validate
	case 1:
		result.validateValues1(stream, result.valuesConstraint[0])
	default:
		result.validateValues2(stream, result.valuesConstraint[0], result.valuesConstraint[1])
	}
}

// Parses via [Climate.Parse] and verifies via [Result.Verify].
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
func (cl Climate) Abort(message string, err error, options ...interface{}) {

	var exiter internal.Exiter

	stream, _ := parse_Stream_from_options_(options...)
	if stream == nil {

		stream = cl.stream
	}
	if stream == nil {

		stream = os.Stderr
	}

	exiter, _ = parse_Exiter_from_options_(options...)
	if exiter == nil {

		exiter = cl.exiter
	}

	var suffix string
	switch cl.UsageHelpSuffix {
	default:
		var first rune
		for _, c := range cl.UsageHelpSuffix {
			first = c
			break
		}

		if unicode.IsPunct(first) {
			suffix = cl.UsageHelpSuffix
		} else {
			suffix = "; " + cl.UsageHelpSuffix
		}

	case ":":

		suffix = UsageHelpSuffix_Default
	case "":
		// do nothing
	}

	if err != nil {

		fmt.Fprintf(stream, "%s: %s: %v%s\n", cl.ProgramName, message, err, suffix)
	} else {

		fmt.Fprintf(stream, "%s: %s%s\n", cl.ProgramName, message, suffix)
	}

	exiter.Exit(1)
}

// Determines if the given flag is specified
func (result Result) FlagIsSpecified(id interface{}) bool {

	return result.arguments.FlagIsSpecified(id)
}

// Looks for a flag with the given id - name, or the specification instance - and
// returns it and the value true if found; if not, returns nil and false.
func (result Result) LookupFlag(id interface{}) (*clasp.Argument, bool) {

	return result.arguments.LookupFlag(id)
}

// Looks for an option with the given id - name, or the specification instance - and
// returns it and the value true if found; if not, returns nil and false.
func (result Result) LookupOption(id interface{}) (*clasp.Argument, bool) {

	return result.arguments.LookupOption(id)
}

/* ///////////////////////////// end of file //////////////////////////// */
