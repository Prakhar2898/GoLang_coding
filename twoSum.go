package main

import "sort"

func twoSum(n []int, target int) (int, int) {
	i := 0
	j := len(n) - 1
	sort.Ints(n)

	for i < j {
		if n[i]+n[j] < target {
			i = i + 1
			continue
		}
		if n[i]+n[j] > target {
			j = j - 1
			continue
		}
		if n[i]+n[j] == target {
			return n[i], n[j]
		}
	}
	return -1, -1

}
