# 78. Subsets

## 📌 Problem Link
[LeetCode 78](https://leetcode.com/problems/subsets/)

## 📝 Problem Summary
Given an integer array `nums` of unique elements, return all possible **subsets** (the power set). The solution must not contain duplicate subsets.

---

## 🛠 Strategic Pattern: Backtracking (DFS)
Backtracking is an exhaustive search algorithm. It is used when you need to explore **all possibilities** and can "turn back" when a path is finished.

### 🧠 The Mental Model: "The Suitcase"
1. **The Goal:** Fill a suitcase (`current`) with different combinations of items.
2. **The Rule:** Once you decide to look at items starting from index `i`, you never look back at items before `i` (to avoid duplicates like `[1,2]` and `[2,1]`).
3. **The Rhythm:**
    *   **Pack:** Put an item in the suitcase (`append`).
    *   **Explore:** See what else fits by calling the function again (`backtrack(i + 1)`).
    *   **Unpack:** Take that specific item out (`slice = slice[:len-1]`) so you can try the next available item in the buffet line.



### ⚠️ Go-Specific Warning: Slice Mutability
In Go, slices are headers pointing to an underlying array. 
* **The Trap:** If you `append(result, current)`, you are appending a reference. As `current` changes during backtracking, all sets in `result` will change too!
* **The Fix:** You must create a "snapshot" of the current state using `copy()`.

---

## 🚀 Complexity Analysis
- **Time Complexity:** $O(N \cdot 2^N)$
    - There are $2^N$ possible subsets.
    - We spend $O(N)$ time creating a copy for each subset.
- **Space Complexity:** $O(N)$
    - This is the maximum depth of the recursion stack (how deep the tree goes).
