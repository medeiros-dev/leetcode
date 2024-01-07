import(
    "strconv"
)

func fizzBuzz(n int) []string {
    
    var result []string
    for x := 1; x <= n; x++{
        if x%3 == 0 && x%5 == 0{
            result = append(result, "FizzBuzz")
        }else if x%3 == 0{
            result = append(result, "Fizz")
        }else if x%5 == 0{
            result = append(result, "Buzz")
        }else{
            result = append(result, strconv.Itoa(x))
        }
    }

    return result

}