package array

import (
	"fmt"
	"testing"
)

func TestSpiralMatrix2(t *testing.T) {
	res := generateMatrix(4)

	fmt.Printf("%+v\n", res)

	res = generateMatrix(3)

	fmt.Printf("%+v\n", res)
}
