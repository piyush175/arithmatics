package arithmatics

import (
	"fmt"
)

func IsPrime(num int) bool {
	for i := 2; i < int(num/2); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func WabaFeet() {
	fmt.Println("Wa waba feet!!!!!!!!!!!!!!!")
}
