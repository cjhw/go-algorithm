package Recall

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var (
		res  [][]int
		path []int
		used []bool
	)

	res, path = make([][]int, 0), make([]int, 0, len(candidates))
	used = make([]bool, len(candidates))
	sort.Ints(candidates)

	var dfs func(candidates []int,start int,target int) 

	dfs=func(candidates []int, start, target int) {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i:=start;i<len(candidates);i++ {
            if candidates[i] > target {  
                break
            }
			if i > 0 && candidates[i] == candidates[i-1]  && used[i-1] == false { 
            	continue
        	}
			used[i] = true
			path = append(path, candidates[i])
			dfs(candidates,i+1,target-candidates[i])
			used[i] = false
			path=path[:len(path)-1]
		}
	}

	dfs(candidates, 0, target)
    return res
}