package day8

import (
	"testing"
)

func Test1(t *testing.T) {
	in := []samplesAndDisplays{
		{Samples: []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"},
			Displays: []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"}},
		{Samples: []string{"edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec"},
			Displays: []string{"fcgedb", "cgb", "dgebacf", "gc"}},
		{Samples: []string{"fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef"},
			Displays: []string{"cg", "cg", "fdcagb", "cbg"}},
		{Samples: []string{"fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega"},
			Displays: []string{"efabcd", "cedba", "gadfec", "cb"}},
		{Samples: []string{"aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fcge", "gcbea", "fcaegb", "dgceab", "fcbdga"},
			Displays: []string{"gecf", "egdcabf", "bgf", "bfgea"}},
		{Samples: []string{"fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gcfdb", "baec", "bfadeg", "bafgc", "acf"},
			Displays: []string{"gebdcfa", "ecba", "ca", "fadegcb"}},
		{Samples: []string{"dbcfg", "fgd", "bdegcaf", "fgec", "aegbdf", "ecdfab", "fbedc", "dacgb", "gdcebf", "gf"},
			Displays: []string{"cefg", "dcbef", "fcge", "gbcadfe"}},
		{Samples: []string{"bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd"},
			Displays: []string{"ed", "bcgafe", "cdgba", "cbgef"}},
		{Samples: []string{"egadfb", "cdbfeg", "cegd", "fecab", "cgb", "gbdefca", "cg", "fgcdab", "egfdb", "bfceg"},
			Displays: []string{"gbdfcae", "bgc", "cg", "cgb"}},
		{Samples: []string{"gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc"},
			Displays: []string{"fgae", "cfgab", "fg", "bagce"}},
	}
	out := countKnownDisplays(in)
	expected := 26
	if out != expected {
		t.Errorf("wrong result: got %d, expected %d", out, expected)
		return
	}
}

func Test2(t *testing.T) {
	in := []samplesAndDisplays{
		{Samples: []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"},
			Displays: []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"}},
		{Samples: []string{"edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec"},
			Displays: []string{"fcgedb", "cgb", "dgebacf", "gc"}},
		{Samples: []string{"fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef"},
			Displays: []string{"cg", "cg", "fdcagb", "cbg"}},
		{Samples: []string{"fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega"},
			Displays: []string{"efabcd", "cedba", "gadfec", "cb"}},
		{Samples: []string{"aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fcge", "gcbea", "fcaegb", "dgceab", "fcbdga"},
			Displays: []string{"gecf", "egdcabf", "bgf", "bfgea"}},
		{Samples: []string{"fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gcfdb", "baec", "bfadeg", "bafgc", "acf"},
			Displays: []string{"gebdcfa", "ecba", "ca", "fadegcb"}},
		{Samples: []string{"dbcfg", "fgd", "bdegcaf", "fgec", "aegbdf", "ecdfab", "fbedc", "dacgb", "gdcebf", "gf"},
			Displays: []string{"cefg", "dcbef", "fcge", "gbcadfe"}},
		{Samples: []string{"bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd"},
			Displays: []string{"ed", "bcgafe", "cdgba", "cbgef"}},
		{Samples: []string{"egadfb", "cdbfeg", "cegd", "fecab", "cgb", "gbdefca", "cg", "fgcdab", "egfdb", "bfceg"},
			Displays: []string{"gbdfcae", "bgc", "cg", "cgb"}},
		{Samples: []string{"gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc"},
			Displays: []string{"fgae", "cfgab", "fg", "bagce"}},
	}
	out, err := sumAllNumbers(in)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}
	expected := 61229
	if out != expected {
		t.Errorf("wrong result: got %d, expected %d", out, expected)
		return
	}
}
