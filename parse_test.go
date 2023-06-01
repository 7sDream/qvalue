package qvalue_test

import (
	"errors"
	"testing"

	"github.com/7sDream/qvalue"
)

type test struct {
	s   string
	err error
	vs  []string
	qs  []uint
}

func newOkTest(s string, vs []string, qs []uint) test {
	return test{
		s, nil, vs, qs,
	}
}

func newErrorTest(s string, err error) test {
	return test{
		s, err, nil, nil,
	}
}

func TestParseSuccess(t *testing.T) {
	o := qvalue.NewOptions()

	tests := []test{
		// No Value
		newOkTest("", []string{""}, []uint{1000}),
		// One value without quality
		newOkTest(
			"gzip", []string{"gzip"}, []uint{1000},
		),
		// One value with quality
		newOkTest(
			"gzip;q=0.3", []string{"gzip"}, []uint{300},
		),
		// Multi value
		newOkTest(
			"gzip, br;q=0.9, deflate;q=0.8, *;q=0.7",
			[]string{"gzip", "br", "deflate", "*"},
			[]uint{1000, 900, 800, 700},
		),
		// Multi value unordered
		newOkTest(
			"br;q=0.9, gzip, *;q=0.7, deflate;q=0.8",
			[]string{"gzip", "br", "deflate", "*"},
			[]uint{1000, 900, 800, 700},
		),
		// Multi value same quality
		newOkTest(
			"br;q=0.9, gzip, *;q=0.7, deflate;q=0.9",
			[]string{"gzip", "br", "deflate", "*"},
			[]uint{1000, 900, 900, 700},
		),
		// Example from MDN: https://developer.mozilla.org/en-US/docs/Glossary/Quality_values
		newOkTest(
			"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
			[]string{"text/html", "application/xhtml+xml", "application/xml", "*/*"},
			[]uint{1000, 1000, 900, 800},
		),
	}

	for i, test := range tests {
		qvs, err := o.ParseAndSort(test.s)
		if err != nil {
			t.Fatalf("Unexpected error at test %d(%s): %s", i, test.s, err.Error())
		}
		if len(qvs) != len(test.vs) {
			t.Fatalf("Result length not match at test %d(%s): %d != %d", i, test.s, len(qvs), len(test.vs))
		}
		for j, qv := range qvs {
			targetValue := test.vs[j]
			targetQuality := test.qs[j]
			if qv.Value != targetValue || qv.Quality != targetQuality {
				t.Fatalf("Result not correct at test %d(%s)[%d]: %#v", i, test.s, j, qv)
			}
		}
	}
}

func TestParseError(t *testing.T) {
	o := qvalue.NewOptions()

	tests := []test{
		newErrorTest("br, gzip;q=what", qvalue.QualityParseError),
		newErrorTest("gzip;q=nan", qvalue.QualityRangeError),
		newErrorTest("br, gzip;q=2", qvalue.QualityRangeError),
		newErrorTest("gzip;q=2.0", qvalue.QualityRangeError),
	}

	for i, test := range tests {
		_, err := o.Parse(test.s)
		if err == nil {
			t.Fatalf("No error for test %d(%s)", i, test.s)
		}
		if !errors.Is(err, test.err) {
			t.Fatalf("Error not match for test %d(%s), wanted = %s, got = %s", i, test.s, test.err.Error(), err.Error())
		}
	}
}

func TestIgnoreError(t *testing.T) {
	o := qvalue.NewOptions().IgnoreQualityParseError(0).IgnoreQualityRangeError(1000)

	tests := []test{
		newOkTest("gzip;q=what", []string{"gzip"}, []uint{0}),
		newOkTest("br, gzip;q=2", []string{"br", "gzip"}, []uint{1000, 1000}),
	}

	for i, test := range tests {
		qvs, err := o.Parse(test.s)
		if err != nil {
			t.Fatalf("Unexpected error at test %d(%s): %s", i, test.s, err.Error())
		}
		if len(qvs) != len(test.vs) {
			t.Fatalf("Result length not match at test %d(%s): %d != %d", i, test.s, len(qvs), len(test.vs))
		}
		for j, qv := range qvs {
			targetValue := test.vs[j]
			targetQuality := test.qs[j]
			if qv.Value != targetValue || qv.Quality != targetQuality {
				t.Fatalf("Result not correct at test %d(%s)[%d]: %#v", i, test.s, j, qv)
			}
		}
	}
}
