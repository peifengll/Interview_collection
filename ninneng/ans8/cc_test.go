package main

import (
	"fmt"
	"testing"
)

func TestCal(t *testing.T) {
	for i := 3990; i <= 4000; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println(365*8 + 366*2)
}
