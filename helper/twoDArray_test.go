package helper_test

import (
	"github.com/rxue92/leetcode-go/helper"
	"testing"
	"fmt"
)

func TestString2Matrix(t *testing.T) {
	input := " [[3,40 ], [2,150]] "
	fmt.Println(helper.StringTo2DArray(input))

}
