# libCLImate.Go - Changes <!-- omit in toc -->


## 0.8.2 - 20th August 2026

* added **Version()** (replacing the **Version** constant), formed by **ver2go.CombineVersion()**;
* documented **VersionString()**;
* **VersionAB** now uses **ver2go.Release**;
* updated **ver2go** to 0.2.0-beta1;
* updated **ANGoLS** to 0.11.0;
* updated **CLASP.Go** to 0.18.0;
* updated **test/scratch/libver.go** to use **VersionString()**;
* added **examples/libver** program;
* CI modernisation (matrix + lint);
* fixed **golangci-lint** issues in **climate.go** (**ineffassign**) and **test/scratch/libver.go** (**gci**);
* CI reliability fixes (macOS test linking; golangci-lint config verification disabled in CI);
* boilerplate additions (scripts, markdown docs, project identity);
* removed retired Go Report Card badge from README;
* version string updated for the 0.8.2 release;


## 0.8.2-beta1 - 18th August 2025

* GitHub Actions;
* `interface{}` => `any`;
* boilerplate;
* documentation;


## 0.8.1 - 12th March 2025

* + further application of `UsageHelpSuffix` to various contingent reports, and now operative by default;


## 0.8.0-beta1 - 4th March 2025

* + added `UsageHelpSuffix`;


## 0.8.0-alpha1 - 4th March 2025

* + added values-constraints, specified by the `Climate` fields `ValuesConstraint` and `ValueNames`;


## 0.7.0-alpha3 - 4th March 2025

* ~ moved a bunch of elements to `internal`;


## 0.7.0-alpha2 - 28th February 2025

* ~ fixed use of **ANGoLS**' `slices` package;


## 0.7.0-alpha1 - 28th February 2025

* + added **examples/bit_flags.go**;
* ~ improved documentation markup;


## 0.6.1-alpha5 - 28th February 2025

* ~ updated dependencies;


## 0.6.1-alpha4 - 28th February 2025

* - removed unncessary use of `iota`;
* + added more test cases;
* ~ tidying / doc-markup;
* ~ formatting;


## 0.6.1-alpha3 - 26th February 2025

* fixed all references to **LibCLImate.Go** (=> **libCLImate.Go**);
* now `InitFlag_NoHelpFlag` and `InitFlag_NoVersionFlag` also suppress recognition of `"--help"` and `"--version"`, as well as their inclusion in usage;
* `..._None` flags now have value 0 (and their types are changed from `int` to `int64`);
* added **climate_test.go**;
* added **/test/scratch/libver.go**;
* added more test cases;
* standard go formatting (for examples);
* improved documentation markup;


## 0.6.1-alpha2 - 25th February 2025

* Fixed module name (from **libclimate** to **github.com/synesissoftware/libCLImate.Go**);


## 0.6.1-alpha1 - 24th February 2025

* updated for use of Go modules;
* standard go formatting;


## 0.6.0 - 10th April 2019

* ~ updated in light of breaking changes in **CLASP.Go** 0.15;


## 0.5.0 - 8th April 2019

* ~ ``AddAlias()`` method now takes two string parameters: resolved_name, alias_name;


## 0.4.2 - 30th March 2019

* + now supports ``InitFlag_PanicOnFailure`` flag;
* + now supports ``ParseFlag_PanicOnFailure`` flag;
* + now support ``ParseFlag_DontCheckUnused`` flag;
* ~ ensuring caller-provided exiter and stream are supported (for testing);
* + unit-tests for ``Climate.Abort()``;
* + unit-tests for ``Climate.ShowUsage()``;
* + unit-tests for ``Climate.ShowVersion()``;
* ~ abstracted out common test constructs into **common_test.go**;
* + added development/testing-only dependency on **ANGoLS** (http://github.com/synesissoftware/ANGoLS/);

NOTE: requires latest version (0.14.2) of **CLASP.Go** (https://github.com/synesissoftware/CLASP.Go)


## 0.4.1 - 29th March 2019

* + now obeys InitFlag_PanicOnFailure and ParseFlag_PanicOnFailure;
* + now obeys ParseFlag_DontCheckUnused;


## 0.4.0 - 29th March 2019

* + added ``Climate.VersionPrefix`` field;
* ~ substantial internal changes to ensure that mocking works for unit-tests;


## 0.3.0 - 28th March 2019

* + added ``Climate.ProgramName`` field;
* + added ``Climate.Abort()`` method;


## 0.2.0 - 25th March 2019

* + added ``Climate.AddFlagFunc()`` and ``Climate.AddOptionFunc()`` methods, which take callbacks to be executed when the flag/option is encountered in the parsed command-line arguments;
* + added **examples/parse_and_verify.go** example which illustrates use of the callbacks;


## 0.1.0 - 23rd March 2019

FIRST PUBLIC RELEASE


<!-- ########################### end of file ########################### -->

