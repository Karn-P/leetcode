package main

func lengthOfLongestSubstring(s string) int {
    maxLen := 0
    // two pointer
    // [a b |c d b| x]
    left := 0
    charMap := make(map[byte]int)

    // increase right pointer
    for right := 0; right < len(s); right++ {
        // get current character
        char := s[right]

        // check if the current char is in the map and should not move before left pointer
        if idx, exists := charMap[char]; exists && idx >= left {
            // if yes update the left position to last seen + 1
            left = idx + 1
        }

        // update last seen position
        charMap[char] = right

        // check if the current window is more than max len
        if (right - left + 1) > maxLen {
            maxLen = right - left + 1
        }
    }
    return maxLen

}
