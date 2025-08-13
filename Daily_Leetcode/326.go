/*
https://leetcode.com/problems/power-of-three/

In this problem, we are limited to signed positive integers (32-bit)
Therefore, the max power of 3 in this number range is 3^19=1162261467
Any number that perfectly divides 3^19 with no remainder must have a prime
factorization composed of only 3, aka, a power of 3.
Therefore, we can complete this in a 1-liner
(rarely possible for non-Python solutions!)

Of course, you also can't be a power of 3 if non-positive
Take advantage of short-circuiting to avoid divide by zero during % operator
*/
func isPowerOfThree(n int) bool {
	return (n > 0) && (1162261467%n == 0)
}