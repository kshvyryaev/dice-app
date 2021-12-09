package main

import (
	"fmt"
)

func main() {
	resultMultiplier := 1000000
	dice := NewDice(6)

	onceResult := dice.RollOnce()
	fmt.Println("Once result:", onceResult*resultMultiplier)
	manyResult := dice.RollMany(resultMultiplier)
	fmt.Println("Many result:", manyResult)
}
