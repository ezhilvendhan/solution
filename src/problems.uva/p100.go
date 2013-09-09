package p100

import (
    "fmt"
    "os"
    "strconv"
    "time"
)

func Run() {
    startTime := time.Now() 
    args := os.Args[1:]
    for i, j, n:= 0, 1, len(args); i < n && j < n; {
	num1, err := strconv.Atoi(args[i])
	if(err == nil) {
	    num2, err := strconv.Atoi(args[j])
	    if(err == nil) {
	        _100(num1, num2)
                i = i+2
                j = j+2
       	    } else {
		fmt.Println(err)
		os.Exit(2) 
	    }
	} else {
            fmt.Println(err)
	    os.Exit(2)
	}
    }
    elapsedTime := (time.Now().Sub(startTime))
    fmt.Println("Elapsed time: ", elapsedTime)
}

type _100Max struct {
    allNum []int
    cycle int
    num int
}

func _100(i, j int) {
    cycle, cnt:= 0, 0
    obj := _100Max{[]int{} , 0, 0}
    for n := i ; n <= j; n++ {
	cnt = 0
	cycle = _100Logic(n, &cnt)
	if(obj.cycle == cycle) {
	    obj.allNum = append(obj.allNum, n)
	}
	if(obj.cycle < cycle) {
	    obj.cycle = cycle
	    obj.allNum = []int {}
	    obj.allNum = append(obj.allNum, n)
	    obj.num = n
	}
    }
    fmt.Println(i, " ", j, " ",obj.num)
}

func _100Logic(n int, cnt *int) int {
    *cnt++
    if (n != 1) {
        if (n%2 == 1) {
	    _100Logic(3*n+1, cnt)
	} else {
	    _100Logic(n/2, cnt)
	}
    }
    return *cnt
}
