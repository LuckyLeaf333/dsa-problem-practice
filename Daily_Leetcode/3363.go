/*
https://leetcode.com/problems/find-the-maximum-number-of-fruits-collected/

Problem: 3 children on a square nxn board start in 3 corners of the board and collect fruits from each square while they move to the bottom right corner of the board. Want to find the maximum number of fruits the can collect. Each room can have a different number of fruits.

The children can move diagonally but the one starting in (n-1, 0) must move right (j+1) on every step.
The child starting in (0, n-1) must move down (i+1) with every step.

Intuition: The (0, 0) child actually only has 1 path of length n-1 to the bottom right corner (n-1, n-1), which is taking the diagonal path across the board.

It is not possible for the (n-1, 0) child or (0, n-1) child to cross over the path of the (0, 0) child. As in, the (n-1, 0) child cannot make it to any tile with coordinate (x, x+1) for x < n-2. To do so, they would need to take at least 1 step down (not diagonal), which they are not allowed to do. This is because if you are at (x, x+1), even if you take as many down-diagonal steps as possible, you will keep going to coordinates of form (x, x+1), and can't reach (n-1, n-1) by taking only diagonal steps or right steps.

Therefore, each child has an optimal space of movement that is disjoint from the other child (in the optimal strategy, no child will cross paths)

This is a DP problem:
- (0, 0) child will always collect fruits on the diagonal
- (n-1, 0) child is confined to bottom left half of board
- (0, n-1) child is confined to top right half of board

Need to be careful of some unreachable tiles
Invariants:
- If the (n-1, 0) child has i-coordinate y, their j-coordinate is at most
- If the (n-1, 0) child is at (y, x), then y >= x
- If the (0, n-1) child has i-coordinate y, then their j-coordinate is at least n-y-1
- If the (0, n-1) child is at (y, x), then x >= y

Time complexity: O(n^2)
Space complexity: O(n^2)
*/

func maxCollectedFruits(fruits [][]int) int {
	maxFruits := 0
	boardDim := len(fruits)
	// First, count fruits on the diagonal for the (0, 0) child's path
	for i := 0; i < boardDim; i++ {
		maxFruits += fruits[i][i]
		fruits[i][i] = 0
	}

	// Second, use DP to find max fruits for (n-1, 0), bottom left child
	dp1 := make([][]int, boardDim)
	for i := 0; i < boardDim; i++ {
		dp1[i] = make([]int, boardDim)
	}

	for steps := 0; steps < boardDim; steps++ {
		j := steps // j-coordinate = step count
		for i := boardDim - 1; i >= boardDim-steps-1 && i >= j; i-- {
			if j-1 >= 0 {
				dp1[i][j] = max(dp1[i][j], dp1[i][j-1])
				if i+1 <= boardDim-1 {
					dp1[i][j] = max(dp1[i][j], dp1[i+1][j-1])
				}
				dp1[i][j] = max(dp1[i][j], dp1[i-1][j-1])
			}
			dp1[i][j] += fruits[i][j]
		}
	}
	maxFruits += dp1[boardDim-1][boardDim-1]

	dp2 := make([][]int, boardDim)
	for i := 0; i < boardDim; i++ {
		dp2[i] = make([]int, boardDim)
	}
	for steps := 0; steps < boardDim; steps++ {
		i := steps // i-coordinate = step count
		for j := boardDim - 1; j >= boardDim-steps-1 && j >= i; j-- {
			if i-1 >= 0 {
				dp2[i][j] = max(dp2[i][j], dp2[i-1][j])
				if j+1 <= boardDim-1 {
					dp2[i][j] = max(dp2[i][j], dp2[i-1][j+1])
				}
				dp2[i][j] = max(dp2[i][j], dp2[i-1][j-1])
			}
			dp2[i][j] += fruits[i][j]
		}
	}
	maxFruits += dp2[boardDim-1][boardDim-1]

	return maxFruits
}