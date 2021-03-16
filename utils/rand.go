package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)
func init(){
	fmt.Println("initialize rand seed using system time")
	rand.Seed(time.Now().UnixNano())
}

func RandNumbers(length int) (numbers string) {
	if length < 1 {
		return
	}
	for i := 0; i < length; i++ {
		numbers += strconv.Itoa(rand.Intn(10))
	}
	return numbers
}

