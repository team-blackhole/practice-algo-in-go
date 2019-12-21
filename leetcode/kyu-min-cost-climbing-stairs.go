package leetcode

/*
	https://leetcode.com/problems/min-cost-climbing-stairs/
*/

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	cost1 := cost[length-1]
	cost2 := cost[length-2]
	temp := 0
	for i := length - 3; i >= 0; i-- {
		temp = cost2
		cost2 = min(cost1, cost2) + cost[i]
		cost1 = temp
	}
	return min(cost2, cost1)
}

func min(a, b int) (k int) {
	if a >= b {
		k = b
	}
	k = a
	return

}
