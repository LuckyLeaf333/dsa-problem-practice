class Solution {
    /*
    https://leetcode.com/problems/power-of-two/

    Determine if a number is a power of 2.
    Simply use bitwise left shift
    Use long datatype to avoid integer overflow (max int is 2^31-1)
    */
public:
    bool isPowerOfTwo(int n) {
        long powerOfTwo = 1;
        while (powerOfTwo < long(n)) {
            powerOfTwo = powerOfTwo << 1;
        }
        return powerOfTwo == long(n);
    }
};