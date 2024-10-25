package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	input := []int{1, 2, 3, 4, 5}
	result := []int{}
	for _, data := range input {
		wg.Add(1)
		go processData(&wg, &result, data)
	}
	wg.Wait()
	fmt.Println("the result is ", result)
}
func processData(wg *sync.WaitGroup, result *[]int, data int) {
	defer wg.Done()
	*result = append(*result, data*2)
}
