# LeetCode 54: Spiral Matrix

## 📌 Overview
This repository contains a Go implementation of the Spiral Matrix problem. The goal is to return all elements of an `m x n` matrix in spiral order.

## 🚀 Strategy: Boundary Collapse
Instead of using complex coordinate math or tracking directions with state variables, this solution uses the **Four-Wall Boundary** approach.

### How it works:
1. Define four boundaries: `top`, `bottom`, `left`, and `right`.
2. Traverse from `left` to `right` along the `top` row, then increment `top`.
3. Traverse from `top` to `bottom` along the `right` column, then decrement `right`.
4. Traverse from `right` to `left` along the `bottom` row (if still valid), then decrement `bottom`.
5. Traverse from `bottom` to `top` along the `left` column (if still valid), then increment `left`.
6. Repeat until the boundaries cross.

## 🛠️ Implementation Details
- **Language:** Go (Golang)
- **Time Complexity:** O(M × N) — Every element is visited exactly once.
- **Space Complexity:** O(1) — Auxiliary space is constant (excluding the result slice).

### Handling Edge Cases
The solution includes internal guard checks (`if top <= bottom` and `if left <= right`). These are critical for handling **non-square matrices** (e.g., 3x1 or 1x5) to prevent the algorithm from re-traversing rows or columns that have already been moved inward.

## 📝 Key Lessons
* **Pointer Discipline:** Moving boundaries inward is much cleaner than managing a `direction` variable.
* **Non-Square Matrices:** Always test with a single row or single column matrix to ensure the loop exit conditions are robust.
* **Go Idioms:** Using `make([]int, 0, rows*cols)` for the result slice is more efficient as it pre-allocates memory for the expected number of elements.
