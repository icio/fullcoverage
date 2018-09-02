package fullcoverage_test

import (
	. "github.com/icio/fullcoverage"

	"fmt"
	"testing"
)

// TestFully ensures that the BucketName function correctly directs the shapes
// into their rightful homes. This test provides 100% test coverage against the
// BucketName function, and aims to ensure that the following provenance is
// attained:
//
//     1. Blue shapes go into the bucket for blue shapes
//     2. Yellow shapes go into the bucket for yellow shapes
//     3. Small shapes go into the bucket for small shapes
//     4. Medium-sized 4-sided shapes go into the bucket for medium quadrangles
//
// To confim everything is working:
//
//     $ go test -cover
//     PASS
//     coverage: 100.0% of statements
//     ok  	github.com/icio/fullcoverage	0.009s
//
// But something is wrong...
func TestFully(t *testing.T) {
	tests := []struct {
		shape          Shape
		expectedBucket string
	}{
		{Shape{Color: ColorBlue}, "Blue bucket"},
		{Shape{Color: ColorYellow}, "The yellow yins"},
		{Shape{Size: SizeSmall}, "Bucket o' small"},
		{Shape{Sides: 4, Size: SizeMedium}, "Medium quadrangles"},
		{Shape{}, ""}, // Default.
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			if actualBucket := BucketName(test.shape); actualBucket != test.expectedBucket {
				t.Errorf("expected bucket %q for shape %#v but got bucket %q", test.expectedBucket, test.shape, actualBucket)
			}
		})
	}
}
