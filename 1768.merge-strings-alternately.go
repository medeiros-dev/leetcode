func mergeAlternately(word1 string, word2 string) string {
  
    var result string

	for i := 0; i < len(word1) + len(word2); i++{

		if i/2 > len(word1) - 1{
			for x := i/2; x < len(word2); x++{
				result = result + string(word2[x])
			}
			break
		}else if i/2 > len(word2) - 1{
			for x := i/2; x < len(word1); x++{
				result = result + string(word1[x])
			}
			break
		}

		if (i%2 == 0 && i/2 <= len(word1) - 1){
			result = result + string(word1[i/2])
		}else{
			result = result + string(word2[i/2])
		}
		
	}

    return result
}