func threeSum(nums []int) [][]int {
    var result [][]int
    // 1. sort the num
    sort.Ints(nums)

    for i:= 0; i< len(nums)-2; i++{
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        left := i + 1
        right := len(nums) - 1
        for right > left {
            sum := nums[i] + nums[left] + nums[right]

            if sum == 0 {
                // store to result
                result = append(result, []int{nums[i], nums[left], nums[right]})
                //move pointer
                left++
                right--
                // if left equal last number the skip
                for left< right && nums[left] == nums[left-1]{
                    left++
                }
                // if right equal to last right skip it
                for left< right && nums[right] == nums[right+1]{
                    right--
                }
            } else if sum < 0 {
                // sum is too low - increase the number to sum
                left ++
            } else {
                //sum > 0
                // sum is too high - decrease the number to sum
                right --
            }
        }
    }
    return result
}
