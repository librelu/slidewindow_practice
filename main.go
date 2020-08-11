package main

import (
	"fmt"

	"github.com/xlab/treeprint"
)

var totalResult = map[string][]int{}

func main() {
	tree := treeprint.New()
	slice := []int{2, 2, 2, 2, 3, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 7, 8}
	if result, left := getOrdering(slice); result != nil {
		coordinate := "0"
		t1 := tree.AddBranch(fmt.Sprintf("+%v", result))
		totalResult[coordinate] = result
		findPossible(left, coordinate, t1)
	}

	if result, left := getThreeSames(slice); result != nil {
		coordinate := "1"
		totalResult[coordinate] = result
		t3 := tree.AddBranch(fmt.Sprintf("+%v", result))
		findPossible(left, coordinate, t3)
	}

	if result, left := getFourSames(slice); result != nil {
		coordinate := "2"
		totalResult[coordinate] = result
		t2 := tree.AddBranch(fmt.Sprintf("+%v", result))
		findPossible(left, coordinate, t2)
	}
	fmt.Println(tree.String())
	fmt.Println(totalResult)

}
func findPossible(slice []int, coordinate string, tree treeprint.Tree) {
	if result, left := getOrdering(slice); result != nil {
		coordinate1 := coordinate + "-0"
		totalResult[coordinate1] = result
		t1 := tree.AddBranch(fmt.Sprintf("+%v", result))
		findPossible(left, coordinate1, t1)
	}
	if result, left := getThreeSames(slice); result != nil {
		coordinate2 := coordinate + "-1"
		totalResult[coordinate2] = result
		t2 := tree.AddBranch(fmt.Sprintf("+%v", result))
		findPossible(left, coordinate2, t2)
	}
	if result, left := getFourSames(slice); result != nil {
		coordinate3 := coordinate + "-2"
		totalResult[coordinate3] = result
		t3 := tree.AddBranch(fmt.Sprintf("+%v", result))
		findPossible(left, coordinate3, t3)
	}
}
func getThreeSames(slice []int) ([]int, []int) {
	for p1 := 0; p1 < len(slice); p1++ {
		p2 := p1 + 1
		for p2 < len(slice) {
			if slice[p1] == slice[p2] {
				if (p2 - p1) == 2 {
					return slice[p1 : p2+1], slice[p2+1:]
				}
			} else {
				break
			}
			p2++
		}
	}
	return nil, slice
}
func getFourSames(slice []int) ([]int, []int) {
	for p1 := 0; p1 < len(slice); p1++ {
		p2 := p1 + 1
		for p2 < len(slice) {
			if slice[p1] == slice[p2] {
				if (p2 - p1) == 3 {
					return slice[p1 : p2+1], slice[p2+1:]
				}
			} else {
				break
			}
			p2++
		}
	}
	return nil, slice
}

func getOrdering(slice []int) ([]int, []int) {
	continueCount := 0
	for p1 := 0; p1 < len(slice); p1++ {
		p2 := p1 + 1
		if p2+1 > len(slice) {
			break
		}
		if slice[p2]-slice[p1] == 1 {
			continueCount++
		} else {
			continueCount = 0
		}
		if continueCount == 2 {
			return slice[p2-continueCount : p2+1], slice[p2+1:]
		}
	}
	return nil, slice
}
