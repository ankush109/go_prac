package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("welcome to our shop..")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	num2, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num2 + 1)
	}

	presentTime := time.Now()
	fmt.Println("the current time is ", presentTime.Format("01-02-2006"))

}
