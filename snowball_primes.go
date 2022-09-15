package main

import (
	"fmt"
	"math"
)

func is_prime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var snowball int
	fmt.Print("Enter the first digit of the snowball: ")
	fmt.Scanln(&snowball)

	snowballs := []int{snowball}

	for true {
		new_snowballs := []int{}

		for _, snowball := range snowballs {
			for i := 1; i < 10; i += 2 {
				if is_prime(snowball*10 + i) {
					new_snowballs = append(new_snowballs, snowball*10+i)
				}
			}
		}

		if len(new_snowballs) == 0 {
			break
		}

		snowballs = new_snowballs
	}

	fmt.Printf("%d is the biggest snowball prime number from that starting point!", snowballs[len(snowballs)-1])

}
