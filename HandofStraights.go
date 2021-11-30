//All we need to do is maintain counters of all avaliable numbers
// Then we can sort, because the smallest and biggest numbers has priority. Eg: in [6,....,7,....,8,....,9] 6 can only be used with 7 and 9 can only be used with 8; when "middle" values can be used by both upper and lower bounds. To skip that consideration we just move from the beginning by sorted array. Then all smallest will be resolved first, and by reaching the biggest the only possible option will be left.
//Lastly, we will iterate over an array, make a group and decrease counters. You will see if it is no possible to make a group.

//time complexity: O(nlogn)
//space complexity: O(n)


func isNStraightHand(hand []int, groupSize int) bool {
   if len(hand) < groupSize {
		return false
	}
	m := make(map[int]int)
	for _, v := range hand {
		m[v]++
	}
	sort.Ints(hand)
	for _, v := range hand {
		if m[v] == 0 {
			continue
		}
		for i := 1; i < groupSize; i++ {
			if m[v+i] == 0 {
				return false
			}
		}
		for i := 0; i < groupSize; i++ {
			m[v+i]--
		}
	}
	return true
}


/*
Success
Details 
Runtime: 48 ms, faster than 90.00% of Go online submissions for Hand of Straights.
Memory Usage: 6.9 MB, less than 90.00% of Go online submissions for Hand of Straights.
*/
