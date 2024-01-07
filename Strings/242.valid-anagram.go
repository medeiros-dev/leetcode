func isAnagram(s string, t string) bool {
    
    if len(s) != len(t){
        return false
    }

	aux := make(map[rune]int)

	for _, value := range s{
		_, exist := aux[value]
		if(exist){
			aux[value]++
		}else{
			aux[value] = 1
		}
	}

	for _, value := range t{
		_, exist := aux[value]
		if(exist){
			aux[value]--
			if aux[value] < 0{
				return false
			}
		}else{
			return false
		}
	}
	result := 0
	for _, value := range aux{
		result += value
        if result < 0 {
            return false
        }
	}

	if result == 0{
		return true
	}else{
		return false
	}

}