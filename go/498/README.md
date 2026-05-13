# 498. Diagonal Traverse

## 📌 Problem Link
[LeetCode 498](https://leetcode.com/problems/diagonal-traverse/)

## 📝 Problem Description
Given an `m x n` matrix `mat`, return an array of all the elements of the array in a diagonal order.

## 🛠 Strategic Pattern: State-Based Simulation
This is a classic "Index Management" problem commonly found in Q3 of technical assessments. The core challenge is handling the "bounces" when the traversal hits the boundaries of the matrix.

### Key Logic:
1.  **Direction Toggling:** Maintain a `dir` variable to switch between **Up-Right** (row--, col++) and **Down-Left** (row++, col--).
2.  **Boundary Priority:**
    *   **Moving Up-Right:** If you hit the top edge, move right. If you hit the right edge, move down. *Priority: Check the right edge last to handle the top-right corner safely.*
    *   **Moving Down-Left:** If you hit the left edge, move down. If you hit the bottom edge, move right. *Priority: Check the bottom edge last to handle the bottom-left corner safely.*
3.  **Simulation:** Continue until all $M \times N$ elements are visited.

## 🚀 Complexity Analysis
- **Time Complexity:** $O(M \times N)$ - Each element is visited exactly once.
- **Space Complexity:** $O(1)$ - Excluding the output array.
