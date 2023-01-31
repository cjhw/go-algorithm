package Greed

import "sort"

func findContentChildren(g []int, s []int) int {

	sort.Ints(g)
	sort.Ints(s)

	child:=0

	for i:=0;i<len(s)&&child<len(g);i++ {
		if s[i]>=g[child] {
			child++
		}
	}

	return child
}