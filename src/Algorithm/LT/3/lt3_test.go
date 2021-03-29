package lt3

import (
	"fmt"
	"testing"
)

func TestLt3(t *testing.T) {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
	s = " "
	fmt.Println(lengthOfLongestSubstring(s))
	s = ""
	fmt.Println(lengthOfLongestSubstring(s))
	s = "a"
	fmt.Println(lengthOfLongestSubstring(s))
}
