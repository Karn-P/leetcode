# 1011. Capacity To Ship Packages Within D Days

## 🛠 Strategic Pattern: Binary Search on Solution Space
This problem uses Binary Search not to find an element in a list, but to find the **minimum valid integer** in a range of possible answers.

## 📝 Problem Summary
Find the **minimum** ship capacity that allows us to ship weights in the given order within a specific number of days.

---

## 🛠 Strategic Pattern: Binary Search on Solution Space (Elite Q4)
This is a **Decision Search**. We aren't searching for a number in an array; we are searching for the "sweet spot" capacity that is just big enough to work, but small enough to be the minimum.

### 🧠 The Mental Model: "Manager & Worker"
*   **The Manager (Binary Search):** Manages the range between `low` (max single weight) and `high` (total weight sum). The Manager makes a "guess" at the capacity (`mid`).
*   **The Worker (Greedy Helper):** Actually tries to pack the boxes using that `mid` capacity.
    *   If a box makes the ship overflow, the Worker starts a **new day**.
*   **The Feedback:**
    *   If Worker succeeds: The Manager tries a **smaller** ship (`high = mid`).
    *   If Worker fails: The Manager must try a **larger** ship (`low = mid + 1`).

### Key Logic:
1.  **Identify Search Range:** 
    - `low`: The maximum single weight (the ship must fit the largest item).
    - `high`: The sum of all weights (shipping everything in one day).
2.  **Binary Search (`low < high`):**
    - Calculate `mid` as a trial capacity.
    - Use a **Greedy helper function** to count how many days are required for that capacity.
    - If `daysRequired <= targetDays`, the capacity is feasible; try a smaller one (`high = mid`).
    - Otherwise, the capacity is too small; increase it (`low = mid + 1`).

## 🚀 Complexity Analysis
- **Time Complexity:** $O(N \log S)$, where $N$ is the number of packages and $S$ is the sum of weights.
- **Space Complexity:** $O(1)$.
