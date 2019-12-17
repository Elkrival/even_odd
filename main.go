package main

import ("fmt"
	"time"
	"math/rand"
)

type IntegerList []int
type EvenOrOdd []string
func main() {
	integerList := makeList()
	fmt.Println(integerList)
	 fmt.Println(isEvenOrOdd(integerList))
}
func makeList()IntegerList {
	N := 120
	list := IntegerList{}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	r.Seed(time.Now().UnixNano())
	for index := 0; index < N; index++ {
		number := r.Intn(1998)
		list = append(list, number)
	}
	return list
}

func isEvenOrOdd(integerList IntegerList) EvenOrOdd{
	statusList := EvenOrOdd{}
	for _, integer := range integerList {
		if integer % 2 == 0 {
			statusList = append(statusList, "even")
		}else {
			statusList = append(statusList, "odd")
		}
	}
	return statusList
}
func (i IntegerList) showInts() IntegerList {
	return i
}