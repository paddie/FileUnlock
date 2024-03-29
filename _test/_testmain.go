package main

import unlock "unlock"

import "testing"
import __os__ "os"
import __regexp__ "regexp"

var tests = []testing.InternalTest{
	{"unlock.Test", unlock.Test},
}

var benchmarks = []testing.InternalBenchmark{

}

var matchPat string
var matchRe *__regexp__.Regexp

func matchString(pat, str string) (result bool, err __os__.Error) {
	if matchRe == nil || matchPat != pat {
		matchPat = pat
		matchRe, err = __regexp__.Compile(matchPat)
		if err != nil {
			return
		}
	}
	return matchRe.MatchString(str), nil
}

func main() {
	testing.Main(matchString, tests, benchmarks)
}

