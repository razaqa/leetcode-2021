func twoSum(nums []int, target int) []int {
    result := make(map[int]int)
    for index, value := range nums {
        if _, ok := result[value]; ok {
            return []int{result[value], index}
        } else {
            result[target - value] = index
        }
    }
    return []int{}
}