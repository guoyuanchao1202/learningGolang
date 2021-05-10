package main

import "fmt"

//func main() {
//	еrr := errors.New("foo")
//	var err error
//	if еrr != nil {
//		fmt.Printf("%T %v", err, err)
//	}
//}

//func main() {
//	m := map[string]string{
//		"acbc": "x",
//		"ad":   "x",
//		"dad":  "xasd",
//		"dfr":  "ded",
//		"geg":  "dwf",
//		"greg":"asda",
//		"hyrh":"ger",
//		"kuk":"Ger",
//		"Fgry":"x",
//	}
//	for k, v := range m {
//		fmt.Println(k, " --- ", v)
//	}
//	fmt.Println("---------------")
//
//	for k, v := range m {
//		fmt.Println(k, " --- ", v)
//	}
//}

type test struct {
	a int64
	b int64
	c int64
	d int64
	e int64
}

type test1 struct {
	t  test
	t1 test
	t2 test
	t3 test
}

func main() {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}
	i := 0
	for key, value := range m {
		if i == 0 {
			m[5] = 5
		}
		fmt.Println(key, value)
		i++
	}
}
