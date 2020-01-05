package main

import (
	"fmt"
	"time"
)

func testIf() {
	ok := true
	if ok {
		fmt.Println("ok is true")
	}

	day := "Friday"
	if day == "Friday" {
		fmt.Println("明天不上班")
	} else if day == "Sunday" {
		fmt.Println("周末好快")
	} else {
		fmt.Println("干活啦")
	}

	m := make(map[string]string)
	m["王八"] = "绿豆"
	if v, ok := m["王八"]; ok {
		fmt.Println(v)
	}
}

func testSwitch() {
	switch day := 20; day {
	case 0, 6:
		fmt.Println("周末")
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	default:
		fmt.Println("不合法")
	}

	a, b := 1, 2
	switch {
	case a < b:
		fmt.Println("a < b")
	case a > b:
		fmt.Println("a > b")
	}
}

func testFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for {
		time.Sleep(time.Second)
		fmt.Println("sleep")
	}
}

func main() {
	// testIf()
	// testSwitch()
	testFor()
}
