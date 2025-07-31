import java.util.HashSet;
import java.util.Set;

class Solution {
    /**
    https://leetcode.com/problems/bitwise-ors-of-subarrays/

    Problem: Find the number of distinct bitwise ORs of all non-empty subarrays of arr

    Intuition:
    - Asking for distinct values likely means we should use Hashset
    - O(n^2) solution is too slow; times out

    - Suppose we try subarrays strictly from left to right, with the most leftmost-starting subarray first
    - Suppose for i < j, we scan a subarray starting from j and reach an index k where the subarray OR is the same as one encountered from scanning the subarray starting from i up to index k
    Then we don't need to scan further
     */

    public int subarrayBitwiseORs(int[] arr) {
        ArrayList<HashSet<Integer>> uniqueVals = new ArrayList<HashSet<Integer>>();

        for (int i = 0; i < arr.length; i++) {
            uniqueVals.add(new HashSet<Integer>());
        }

        Set<Integer> ans = new HashSet<>();

        for(int i = 0; i < arr.length; i++) {
            int val = 0;
            for(int j = i; j < arr.length; j++) {
                val = val | arr[j];
                if(uniqueVals.get(j).contains(val)) {
                    break;
                }
                else{
                    ans.add(val);
                    uniqueVals.get(j).add(val);
                }
            }
        }
        return ans.size();
    }
}