class Solution {
    /**
     * https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and
     * 
     * Problem: Find maximum-length continous subarray with maximum bitwise AND
     * 
     * Intuition: If a < b, then a & b < b. 
     * Therefore, we just need to find max-length subarray with maximum value
     * 
     * Space: O(1)--Only need to store a constant number of variables, of type int and boolean
     * Time Complexity: O(N)--We only need to loop through the array once
     */

    public int longestSubarray(int[] nums) {
        int maxSeen = nums[0]; // Maximum number observed
        int longest = 1; // Longest subarray observed with all elements equal to maxSeen so far
        int counter = 1; // Counts the length of the subarray ending at the current index
        boolean continueCount = true; // True if the previous index in nums was equal to maxSeen

        for(int i = 1; i < nums.length; i++) {
            if(nums[i] > maxSeen) {
                // Start counting length of new subarray
                counter = 1;
                maxSeen = nums[i];
                continueCount = true;
                longest = 1;
            }
            else if(nums[i] == maxSeen) {
                // Have a candidate for possible longest answer
                if(continueCount) {
                    counter++; // Previous index was also equal to maxSeen
                }
                else{
                    // Previous index not equal to maxSeen 
                    counter = 1;
                    continueCount = true;
                }
            }
            else {
                // The value of nums[i] is lower than what we've seen before
                continueCount = false;
                counter = 0;
            }
            if (continueCount && counter > longest) {
                // Our current count indicates a new longest max AND subarray
                longest = counter;
            }
        }
        // Return length of longest subarray
        return longest;
    }
}