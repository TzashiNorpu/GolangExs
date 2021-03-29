package lt3

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	lastOccured := make(map[rune]int)
	maxLen := 0
	start := 0
	for i, ch := range runes {
		if v, ok := lastOccured[ch]; ok && v >= start {
			start = lastOccured[ch] + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLen
}
