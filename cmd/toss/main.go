package main

import (
	"fmt"

	coin "github.com/hyrmn/cointoss/pkg/cointoss"
)

func main() {
	var result = coin.Toss()
	fmt.Println(result)
}
