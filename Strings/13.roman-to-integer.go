func romanToInt(s string) int {
    romanNumbers := map[string]int{
        "I":1,
        "V":5,
        "X":10,
        "L":50,
        "C":100,
        "D":500,
        "M":1000,
    }
    
    var sum int
    
    for i := 0; i<len(s); i++{
        if i == len(s)-1{
            sum += romanNumbers[string(s[i])]    
            break
        }
        if romanNumbers[string(s[i])] >= romanNumbers[string(s[i+1])]{
            sum += romanNumbers[string(s[i])]
        }else{
            sum += romanNumbers[string(s[i+1])] - romanNumbers[string(s[i])]
            i++
        }
    }

    return sum
}