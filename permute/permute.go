package permute

import (
	"fmt"
	"math/rand"
)

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func filterPermutations(a [][]int) [][]int {
	filteredPermutations := [][]int{}
	for _, arrangement := range a {
		if !containsAnySelfMap(arrangement) {
			filteredPermutations = append(filteredPermutations, arrangement)
		}
	}
	return filteredPermutations
}

func GetDerangements(playerCount int) []int {
	list := []int{}
	for i := 0; i < playerCount; i++ {
		list = append(list, i+1)
	}
	allPermutations := permutations(list)
	fmt.Printf("Total number of permutations: %d\n", len(allPermutations))
	filteredPermutations := filterPermutations(allPermutations)
	fmt.Printf("Total number of derangements: %d\n", len(filteredPermutations))

	return getRandomDerangement(filteredPermutations)
}

func getRandomDerangement(filteredPermutations [][]int) []int {
	position := rand.Intn(len(filteredPermutations))
	return filteredPermutations[position]
}

func containsAnySelfMap(a []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == i+1 {
			return true 
		}
	}
	return false 
}
