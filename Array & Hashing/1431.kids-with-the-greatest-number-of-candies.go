func kidsWithCandies(candies []int, extraCandies int) []bool {
    
    var greatest int
    
    for _, v := range candies{
        if v > greatest{
            greatest = v
        }
    }

    var result []bool

    for _, v := range candies{
        if v + extraCandies >= greatest{
            result = append(result, true)
        }else{
            result = append(result, false)
        }
    }

    return result

}