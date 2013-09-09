package p100
//UVa Problems in Golang [http://uva.onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=3]
import (
  "fmt"
  "os"
  "strconv"
  "util"
  "time"
)

func Run() {
	start:=time.Now()
	args := os.Args[1:]
	for i, j, n:= 0, 1, len(args); i < n && j < n; {
		num1, err := strconv.Atoi(args[i])
		if(err == nil) {
	    num2, err := strconv.Atoi(args[j])
	    if(err == nil) {
				printMaxCycle(num1, num2)
				i += 2
				j += 2
			} else {
				common.PrintErrorAndExit(err)
	    }
		} else {
			common.PrintErrorAndExit(err)
		}
	}
	elapsed := time.Now()
  fmt.Println("Elapsed: ", elapsed.Sub(start))
}

func printMaxCycle(num1, num2 int) {
	maxCycle, cycle:= 0, 0
	for n := num1 ; n <= num2; n++ {
		cycle = 0
		loop(n, &cycle)
		if(maxCycle < cycle) {
	    maxCycle = cycle
		}
	}
	fmt.Println(num1, " ", num2, " ",maxCycle)
}

func loop(n int, cycle *int) int {
	*cycle++
	if (n != 1) {
		if (n%2 == 1) {
	    loop(3*n+1, cycle)
		} else {
	    loop(n/2, cycle)
		}
	}
	return *cycle
}
