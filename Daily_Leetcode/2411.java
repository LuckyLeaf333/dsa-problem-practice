class Solution {
    /**
     * https://leetcode.com/problems/smallest-subarrays-with-maximum-bitwise-or
     * 
     * Problem: Find the length of the shortest continous subarray with maximum bitwise or
     * for every value in the nums array
     * 
     * Intuition: This is a sliding window problem.
     * Two key observations:
     * 
     * 1. For indices i < j <= k and we know the shortest continous subarray starting at j ends at k
     * then the shortest continous subarray starting a i cannot end any further to the left than index k.
     * 
     * 2. For indices i < k, if we know:
     * - The shortest continous subarray starting a i cannot end any further to the left than index k
     * - For each binary digit, the index of the value closest to the right of i that has a 1 for that digit
     * We can determine if we can remove k from the subarray without decreasing the value of the OR
     * 
     * Complexity:
     * Let C be the maximum ge of elements in 
     * Space: O()
     */

    /**
     * Maximum number of binary digits in a number
     * (It is given in the problem what the max value is; we can set it a bit higher just to be safe)
     */
    private int maxDigits = 40;

    /**
     * Deconstruct num into binary digits
     */
    private void binaryDecomposition(int[] positions, int num, int newPos) {
        int index = 0;
        while(num > 0) {
            if (num % 2 == 1) {
                positions[index] = newPos;
            }
            index++;
            num = num / 2;
        }
    }

    /**
     * Check if an array contains an element
     */
    private boolean arrayContains(int[] arr, int elem) {
        for(int i = 0; i < arr.length; i++) {
            if(arr[i] == elem) {
                return true;
            }
        }
        return false;
    }

    public int[] smallestSubarrays(int[] nums) {
        int[] ans = new int[nums.length];

        ans[nums.length-1] = 1;
        int end = nums.length-1;
        int[] positions = new int[maxDigits];
        binaryDecomposition(positions, nums[end], end);
        for(int start = nums.length-2; start >= 0; start--) {
            binaryDecomposition(positions, nums[start], start);
            while (!arrayContains(positions, end) && end > start) {
                end--;
            }
            ans[start] = end - start + 1;
        }
        return ans;
    }
}