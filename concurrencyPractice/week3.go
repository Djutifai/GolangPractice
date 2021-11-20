package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortThis(arr []int, c chan []int) {
	fmt.Println("Im going to sort this: ", arr)
	sort.Ints(arr)
	c <- arr
}

func splitIntoFourParts(arr []int) [][]int {
	var div, mod int
	var n = int(math.Min(float64(4), float64(len(arr))))
	res := make([][]int, 0)
	div, mod = len(arr)/4, len(arr)%4
	for i := 0; i < n; i++ {
		res = append(res, arr[i*div+int(math.Min(float64(i), float64(mod))):(i+1)*div+int(math.Min(float64(i+1), float64(mod)))])
	}
	return res
}

func main() {
	fmt.Println("Hello, please enter down below any number of integers so it can be sorted by different goroutines :)")
	reader := bufio.NewReader(os.Stdin)
	arr := make([]int, 0)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	strArr := strings.Fields(str)
	for _, part := range strArr {
		val, err := strconv.Atoi(part)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, val)
	}
	c := make(chan []int, 4)
	myArrays := splitIntoFourParts(arr)
	for i := 0; i < len(myArrays); i++ {
		go sortThis(myArrays[i], c)
	}
	finalArray := make([]int, 0)
	for i := 0; i < len(myArrays); i++ {
		finalArray = append(finalArray, <-c...)
	}
	sort.Ints(finalArray)
	fmt.Println(finalArray)
}
