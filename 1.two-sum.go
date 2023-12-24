func twoSum(nums []int, target int) []int {
    var numMap = make(map[int]int)
    for x, v := range nums{
      p := target - v
      if index, found := numMap[v]; found{
          return []int{x, index}
      }
      numMap[p] = x
    }

    return nil
}