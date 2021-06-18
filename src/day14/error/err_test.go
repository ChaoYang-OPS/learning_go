package err_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThenHunderError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	// if n < 0 || n > 100 {
	// 	return nil, errors.New("n should be in [0, 11]")
	// }
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThenHunderError
	}

	fibList := []int{1, 1}
	for i := 2; /* 短变量声明 := */ i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(10); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less.")
		}

		if err == LargerThenHunderError {
			fmt.Println("It is larger")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if _, err := strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	if _, err := GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}
