# 015. 3Sum

## 📌 Problem Link
[LeetCode 15](https://leetcode.com/problems/3sum/)

## 📝 Problem Description
Given an integer array `nums`, return all the triplets `[nums[i], nums[left], nums[right]]` such that `i != left`, `i != right`, and `left != right`, and the sum of the elements is zero. The solution set must not contain duplicate triplets.

## 🛠 Strategic Pattern: Sorting + Two Pointers (Meeting in the Middle)
This problem is a classic extension of the Two-Sum problem. By sorting the array, we can use a directional search to find triplets efficiently.

### Key Logic:
1.  **Sort the Array:** Essential for the two-pointer approach and duplicate management.
2.  **Fix an Anchor:** Iterate through the array, treating the current element as the first part of the triplet.
3.  **Two-Pointer Search:** Initialize `left` at `i + 1` and `right` at the end of the array. Move them toward each other based on the current sum.
4.  **Handle Duplicates:** Skip identical values for the anchor, the left pointer, and the right pointer to ensure all triplets in the result set are unique.

## 🚀 Complexity Analysis
- **Time Complexity:** $O(N^2)$
  - Sorting takes $O(N \log N)$.
  - The nested loops (one for the anchor and one for the two pointers) take $O(N^2)$.
- **Space Complexity:** $O(1)$ or $O(N)$
  - Depending on the implementation of the sorting algorithm. The space used for the result set is generally not counted toward complexity.
