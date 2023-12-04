package day4

import "testing"
func Test_pow_zero(t *testing.T){
	if (pow(2,0) != 1){
		panic("Pow to the power of 0 != 1")
	}
}