class Solution {
    /*
    https://leetcode.com/problems/largest-3-same-digit-number-in-string/

    Very easy problem, just find the 3-char substring with the largest character
    such that each character in the substring is the same
     */
    public String largestGoodInteger(String num) {
        char bestDigit = 0;
        char[] chars = num.toCharArray();
        for (int i = 2; i < chars.length; i++) {
            if (chars[i-2] == chars[i-1] && chars[i-1] == chars[i]) {
                if (chars[i] > bestDigit) {
                    bestDigit = chars[i];
                }
            }
        }
        if (bestDigit == 0) {
            return "";
        }
        // Java quirk: cannot concat char with +, but can if you have a string in it
        return "" + bestDigit + bestDigit + bestDigit;
    }
}