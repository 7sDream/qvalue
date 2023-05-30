package qvalue_test

import (
	"fmt"

	"github.com/7sDream/qvalue"
)

func ExampleOptions_Parse() {
	o := qvalue.NewOptions()
	qvs, _ := o.Parse("gzip, deflate;q=0.8, br;q=0.7, identity;q=0.5, *;q=0.1")
	for _, qv := range qvs {
		fmt.Printf("%s: %d\n", qv.Value, qv.Quality)
	}
	// Output: gzip: 1000
	// deflate: 800
	// br: 700
	// identity: 500
	// *: 100
}
