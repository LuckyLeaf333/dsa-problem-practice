class Solution {
    public int numOfUnplacedFruits(int[] fruits, int[] baskets) {
        boolean[] usedBaskets = new boolean[baskets.length];

        int unplacedCounter = 0;
        for(int i = 0; i < fruits.length; i++) {
            boolean placedFlag = false;
            for(int j = 0; j < baskets.length; j++) {
                if(baskets[j] >= fruits[i] && !usedBaskets[j]) {
                    usedBaskets[j] = true;
                    placedFlag = true;
                    break;
                }
            }
            if (!placedFlag) {
                unplacedCounter++;
            }
        }

        return unplacedCounter;
    }
}