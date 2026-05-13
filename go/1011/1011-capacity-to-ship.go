func shipWithinDays(weights []int, days int) int {
    // 1. find max weights to be low, sum of the weight to be high.
    low, high := 0,0
    for _, w := range weights {
        if w >low {
            low = w
        }
        high += w
    }
    // 2. cal mid to be the capacity of the helper function then loop until the low = high then return mid value

    for high > low {
        mid := low + (high - low) /2
        if canShip(mid, weights, days) {
            high = mid
        } else {
            low = mid + 1
        }
    }
    return low
}

func canShip(capacity int, weights []int, days int) bool {
    dayCount := 1
    currLoad := 0
    for _, w := range weights {
        if currLoad + w > capacity {
            dayCount ++
            currLoad = w
        } else {
            currLoad += w
        }
    }
    return dayCount <= days
}
