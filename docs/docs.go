package main

import (
	"fmt"
	// "net/http"
)

func main() {

	fmt.Println("Lets get start it")

	//1-  variables declaration
	/***************************************
	const x = 5
	var y = 3
	z := 5
	var a string = "ads"
	var b bool = true
	y = 7
	fmt.Printf("y =  %d \n", y)
	fmt.Printf("x =  %d \n", x)
	fmt.Printf("z =  %d \n", z)
	fmt.Printf("a =  %s \n", a)
	fmt.Printf("b =  %t \n", b)

	fmt.Printf("y type=  %T \n", y)
	fmt.Println("Hello GO")
	*************************************/
	// 2- Arrays
	// var b [2]bool

	// b[0] = true
	// b[1] = false

	// a := [3]string{"asd", "nmr", "7ot"}
	// var c = [1]int{1}

	// fmt.Println(a)
	// fmt.Println(c)

	// println(add(2, 3))

	// 3- conditions
	/***************************
	x := 3
	y := 5

	if x > y {
		fmt.Print("x greater than y")
	} else if x == y {
		fmt.Print("x equals y")
	} else {
		fmt.Print("x less than y")
	}
	**********************/

	// 4- Loops, switch , defer

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// sum := 2

	// for sum < 10 {
	// 	fmt.Println(sum)
	// 	sum += sum
	// }

	// var grade float64 = 40

	// switch grade {
	// case 30:
	// 	fmt.Println("bad")
	// 	break
	// case 60:
	// 	fmt.Println("good")
	// 	break
	// case 90:
	// 	fmt.Println("verygood")
	// 	break
	// case 100:
	// 	fmt.Println("excellent")
	// 	break
	// default:
	// 	fmt.Println("nothing")
	// }

	// switch {
	// case grade < 30:
	// 	fmt.Println("bad")
	// 	break
	// case grade < 60:
	// 	fmt.Println("good")
	// 	break
	// case grade < 85:
	// 	fmt.Println("verygood")
	// 	break
	// case grade < 100:
	// 	fmt.Println("excellenet")
	// 	break
	// default:
	// 	fmt.Println("nothing")
	// }

	// defer fmt.Println(grade)

	// fmt.Println("hello")

	// http.HandleFunc("/", handlerFun)
	// http.ListenAndServe(":3000", nil)
	// fmt.Print("Start Server at port 3000")
}

/*********************
// 3- Functions
func add(x, y int) int {
	return x + y
}
**********************/

// func handlerFun(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "<h1> Hello Golang </h1>")
// }
