func findDifference(nums1 []int, nums2 []int) [][]int {


    allNumbers := make(map [int]int)

    var result1 []int
    var result2 []int 

    for _, v := range nums1{
        allNumbers[v] = 1
    }

    for _, v := range nums2{
        if allNumbers[v] == 0  || allNumbers[v] == 2{
            allNumbers[v] = 2
        }else{
            allNumbers[v] = 3
        }
    }

    for key, value := range allNumbers{
        if value == 1{
            result1 = append(result1, key)
        }else if value == 2{
            result2 = append(result2, key)
        }
    }

    return [][]int{result1, result2}

}