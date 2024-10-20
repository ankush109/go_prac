package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var age = 20

type Person struct {
	Name string
	age  int
}
type gasEngine struct {
	mpg     uint8
	gallons uint8
}

func (e gasEngine) milesLeft() uint {
	return uint(e.gallons)
}
func main() {
	fmt.Println("enter a string ..")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thannks for reading.", input)
	var name = "golamng"
	fmt.Print(name)
	i := 1
	fmt.Print(i)
	var myarray [3]int = [3]int{1, 2, 3}
	fmt.Println(myarray)
	ans := add(2, 3)
	print(ans)
	mapdata := map[string]int{"a": 1, "b": 1}
	fmt.Println(mapdata["a"])

	newPerson := Person{Name: "Ankush Banerjee\n", age: 21}
	fmt.Print(newPerson.Name)
	k, b, _ := multi()
	println("k is", k, "b is ", b)
	if k > 1 {
		print("oh no..")
	} else {
		print("all good!\n")
	}
	if rank := getRank(); rank == 1 {
		print("Congo you are rank 1!\n")

	} else {
		print("Study More !")
	}
	for i := 0; i < 10; i++ {
		fmt.Print("the val is ", i)
		print("\n")
	}
	//var myengine gasEngine = gasEngine{20, 10}
	//fmt.Printf("total left in tank %v", myengine.milesLeft())
	slice := []string{"first", "second"}
	var fruitlist = []string{}
	fmt.Printf("type of fruit listi is %T\n", fruitlist)
	//print(slice)
	fruitlist = append(fruitlist, "mango", "banana")
	newslice := append(slice, "fourth", "fifth")
	fmt.Println(newslice)
	ageMap := make(map[string]int)
	ageMap["apple"] = 20

	days := []string{"a", "b", "c", "d"}

	// loops
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	// for Each

	for _, day := range days {
		fmt.Printf("lafda is %v\n", day)
	}

	// making a slice
	highscore := make([]int, 4)
	highscore[0] = 1

	// removing a value from the slice given index

	var courses = []string{"python", "js", "java", "c", "c++"}
	courses = append(courses[:2], courses[2+1:]...)

	fmt.Println(courses)
	numbers := []int{1, 2, 3}
	for i, num := range numbers {
		fmt.Printf("%d ---> %d\n", i, num)
	}
	url := "https://jsonplaceholder.typicode.com/todos/"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

	// pointers in golang
	myorder := newOrder("1", 40.40, "recieved")
	fmt.Println(myorder)

}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

// returns pointer to the struct instance
func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

func add(a, b int) int {
	return a + b
}
func multi() (int, int, int) {
	return 1, 2, 3
}
func getRank() int {
	return 1
}
