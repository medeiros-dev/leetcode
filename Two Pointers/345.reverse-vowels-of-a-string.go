func reverseVowels(s string) string {
    
    type vowels struct{
        value rune
        position int
    }

    var inverse []vowels
	var result []rune

	for i, v := range s{
		if v == 65 || v == 69 || v == 73 || v == 79 || v == 85 || v == 65 + 32 || v == 69 + 32 || v == 73 + 32 || v == 79 + 32 || v == 85 + 32{
			inverse = append(inverse, vowels{v, i})
		}
		result = append(result, v)
	} 

	for i := 0; i <= (len(inverse)/2) -1; i++{
		result[inverse[len(inverse)-i-1].position], result[inverse[i].position] = inverse[i].value, inverse[len(inverse)-i-1].value
	}

    return string(result)
}