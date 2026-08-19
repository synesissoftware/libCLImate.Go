# libCLImate.Go - TODO <!-- omit in toc -->


## Table of Contents <!-- omit in toc -->

- [Functional improvements](#functional-improvements)
- [Performance improvements](#performance-improvements)
- [Packaging improvements](#packaging-improvements)


## Functional improvements

* [ ] a "require" property for flags/options;
* [ ] verification of option values (constraining to type and/or range-of-values);
* [ ] allow succinct values, a la **libCLImate.Ruby**, e.g. `"--verbosity"` (alias `"-v"`) values range `[]string{ "[s]ilent", "[t]erse", "[n]ormal, "[c]hatty", "[v]erbose" }`, allowing for option `"-v c"`;


## Performance improvements

* \<none>


## Packaging improvements

* [ ] Before the next official release: confirm **`go.mod`** (`go 1.21`) and the CI Go-version matrix, bump Synesis `require`s to newly published tags, then run **`go mod tidy`** (not against currently published tags). Prior Synesis Go releases, in order:
  * **ver2go**;
  * **STEGoL**;
  * **ANGoLS**;
  * **CLASP.Go**;


<!-- ########################### end of file ########################### -->
