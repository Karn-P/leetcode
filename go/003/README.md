# 003. Longest Substring Without Repeating Characters

## 📌 Problem Link
[LeetCode 3](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## 📝 Problem Description
Given a string `s`, find the length of the **longest substring** without repeating characters.

**Example:**
- Input: `s = "abcabcbb"`
- Output: `3`
- Explanation: The answer is "abc", with the length of 3.

## 🛠 Strategic Pattern: Sliding Window (Optimized Jump)
This problem represents the **Dynamic Window** pattern. The Sliding Window manages a range that expands to explore new data and contracts to resolve conflicts.

### Key Logic:
1.  **Expansion:** Move the `right` pointer to include new characters.
2.  **Conflict Detection:** Use a `map` or `array` to store the last seen index of each character.
3.  **The "Jump" Contraction:** When a duplicate is found, instead of incrementing the `left` pointer slowly, we jump `left` to `last_seen_index + 1`.
4.  **Safety Guard:** Ensure `left` only moves forward (`idx >= left`) to avoid jumping back into "stale" duplicates outside the current window.

## 🚀 Complexity Analysis
- **Time Complexity:** $O(N)$
  - We traverse the string exactly once. Map/Array operations are $O(1)$.
- **Space Complexity:** $O(min(m, n))$
  - We need space for the hash map to store characters. $m$ is the size of the charset (e.g., 128 for ASCII).
