package main

import (
	"fmt"

	"github.com/google/uuid"
)

//uuidで固定長36文字
func main() {
	fmt.Println(uuid.New().String())
}
