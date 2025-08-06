/**
  https://leetcode.com/problems/fruits-into-baskets-iii/

  Problem:
  Put fruits in baskets such that we have a list of n fruit types and baskets and each fruit type must go in a basket with capacity at least as large as the number of fruits of that type. Each basket can only hold one fruit type.

  Scanning the fruit array from left to right, the next unplaced fruit type must go in the leftmost available basekt it can fit in. Otherwise, the fruit type is unplaced

  Goal: Find the number of unplaced fruits

  Intuition: To place a fruit type, we must find the leftmost available basket it can fit in. Unfortunately, the "leftmost" condition means we cannot just sort the baskets

  Brute force solution is to scan left->right, but in worst case, this is a linear-time operation for every fruit type

  Use square root decomposition to speed things up

  Split basket array into consecutive chunks of sqrt(n) size
  Find the maximum capacity basket in each chunk

  Now, when you scan the basket array for a suitable basket, you can skip chunks that don't fit the fruit type. Once you find a chunk that has a basket that can fit the fruit type, do linear search within the chunk. Once you fill a basket in the chunk, you need to update the local maximum capacity.

  Searching chunks this way takes 3sqrt(n) accesses in the basket array in the worst case

  Time complexity is O(n*sqrt(n))
*/

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	unplacedFruits := 0

	// Trick to compute ceil(n/m) is (n+m-1)/m
	chunkSize := int(math.Sqrt(float64(len(baskets))))
	numChunks := (len(baskets) + chunkSize - 1) / chunkSize
	chunkLocalMax := make([]int, numChunks)
	for i := 0; i < numChunks; i++ {
		// Find the maximum capacity within the ith basket
		for j := i * chunkSize; j < len(baskets) && j < i*chunkSize+chunkSize; j++ {
			chunkLocalMax[i] = max(chunkLocalMax[i], baskets[j])
		}
	}

	for _, fruitQuantity := range fruits {
		placedFlag := false
		for i := 0; i < numChunks; i++ {
			if chunkLocalMax[i] < fruitQuantity {
				// The fruit cannot fit in any basket within the chunk
				continue
			}
			// Linear search through the chunk
			// Also compute the local maximum again while iterating through the chunk because needs to be updated
			newMax := 0
			for j := i * chunkSize; j < len(baskets) && j < i*chunkSize+chunkSize; j++ {
				if baskets[j] >= fruitQuantity && !placedFlag {
					placedFlag = true
					baskets[j] = 0
				}
				newMax = max(newMax, baskets[j])
			}
			chunkLocalMax[i] = newMax
			break
		}
		if !placedFlag {
			unplacedFruits++
		}
	}

	return unplacedFruits
}