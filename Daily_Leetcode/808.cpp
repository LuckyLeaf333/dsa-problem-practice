double soups[5000][5000];

class Solution {
    /*
    https://leetcode.com/problems/soup-servings/description/

    Need to calculate the probability that bowl A runs out of soup before B
    And probability A and B run out of soup at the same time

    Observations:
    - As n approaches infinity, the probability A is consumed completely before B approaches 1.
    - As a function of n, soupServings(n) is monotone increasing
    - For a certain cutoff of n, just return 1
    - Otherwise, do DP to find the probability of A consumed before B
    */
public:
    const int MAX_N = 5000;

    double soupsProbability(int i, int j) {
        return 0.25*(soups[max(i-4, 0)][j] + soups[max(i-3, 0)][max(j-1, 0)] + soups[max(i-2, 0)][max(j-2, 0)] + soups[max(i-1, 0)][max(j-3, 0)]);
    }

    double soupServings(int n) {
        if (n > MAX_N) {
            return 1.0;
        }
        int m = ((n + 25 - 1)/25)+1;
        for(int i = 0; i < m; i++) {
            for(int j = 0; j < m; j++) {
                soups[i][j] = 0;
            }
        }
        soups[0][0] = 0.5;
        for(int i = 1; i < m; i++) {
            soups[0][i] = 1;
        }

        /*
        DP state:
        soups[i][j] contains the answer if the amount of soups are A=i and B=j
        soups[i][j] = soups[i-4][j] + soups[i-3][j-1] + soups[i-2][j-2] + soups[i-1][j-3]
        */
        for(int i = 0; i < m; i++) {
            for(int j = 0; j < m; j++) {
                if (i == 0 || j == 0) {
                    continue;
                }
                soups[i][j] = soupsProbability(i, j);
            }
        }

        return soups[m-1][m-1];
    }
};