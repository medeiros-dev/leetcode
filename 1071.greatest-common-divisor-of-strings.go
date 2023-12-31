func isValidDivision(str string, divisor int) bool{
	
	if len(str) == divisor{
		return true
	}

	if len(str) % divisor != 0{
		return false
	}
	
	for i:= 0; i < len(str) - divisor; i = i + divisor{
		if str[i:i+divisor] != str[divisor:divisor*2]{
			return false
		}
	}

	return true
}


func gcdOfStrings(str1 string, str2 string) string {

    if str1 == str2{
        return str1
    }

	var divisors []int

	for i := 1; i <= len(str1); i++{
		if len(str1)%i == 0{
			divisors = append(divisors, i)
		}
	}

    if len(divisors) > 1{
	    divisors = divisors[:len(divisors)-1] 
    }
	
	for i := len(divisors) -1; i >= 0; i--{
		if isValidDivision(str1, divisors[i]) && isValidDivision(str2, divisors[i]) && str1[:divisors[i]] == str2[:divisors[i]]{
			return str1[:divisors[i]]
		}
	}

	return ""

}
