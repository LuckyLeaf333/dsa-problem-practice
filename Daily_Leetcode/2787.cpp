#include <bits/stdc++.h>
using namespace std;
#define ll long long

class Solution {
    /*
    https://leetcode.com/problems/ways-to-express-an-integer-as-sum-of-powers/

    DP problem where you choose the largest number in the sum
    And then use memoization to store the number of ways to make the
    remaining amount of the sum
    Use 2d dp to store number of sets that only use numbers up to a certain value
    Like subset sum problem
    */
public:
    int numberOfWays(int n, int x) {
        ll dp[301][301];
        ll possible[n];
        int numPos = 0;

        memset(dp, 0, sizeof(dp));
        
        // dp[i][j] = number of ways to make sum equal i using first j numbers in the set, inclusive
        for (int i = 1; i <= n; i++) {
            if (pow(i, x) <= n) {
                possible[numPos] = (int)pow(i, x);
                for (int j = numPos; j <= n; j++) {
                    dp[(int)pow(i, x)][j] = 1;
                }
                numPos++;
            }
        }

        for (int i = 1; i <= n; i++) {
            for (int j = 1; j < numPos; j++) {
                if (possible[j] > i) {
                    continue;
                }
                for (int k = j; k < numPos; k++) {
                    dp[i][k] += dp[i-possible[j]][j-1];
                }
            }
        }
        return dp[n][numPos-1] % 1000000007;
    }
};