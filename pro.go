package main

import (
	"fmt"
)

type counter struct {

}

func main() {
	_100(22, 22)
}

func _100(i, j int64) {
	max, cycle := 0, 0
	for n := i ; n <= j; n++ {
		counter = 1
		cycle = _100Logic(n)
		fmt.Println("Cycle: ", cycle);
		if(max < cycle) {
			max = cycle
		}
	}
	fmt.Println("max is: ", max)
}

func _100Logic(n int64) int {
	fmt.Println(n, "Count: ", counter)
	counter++
	if (n != 1) {
		if (n%2 == 1) {
			_100Logic(3*n+1)
		} else {
			_100Logic(n/2)
		}
	}
	return counter
}
