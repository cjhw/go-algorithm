package Greed

func canCompleteCircuit(gas []int, cost []int) int {
	curSum := 0
	totalSum := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 {
			curSum = 0
			start = i + 1
		}
	}

	if totalSum < 0 {
		return -1
	}

	return start
}