package utils

import (
	"fmt"
)

func Wait() {
	fmt.Printf("[Press enter to proceed]")
	fmt.Scanf(" ")
}
