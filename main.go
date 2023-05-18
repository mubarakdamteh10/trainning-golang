package main

import "fmt"

func PassDataList(scores []int) []int {

	return scores
}

func main() {
	scores := []int{30, 40, 50, 60}
	fmt.Println(PassDataList(scores))
}
