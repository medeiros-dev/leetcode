func longestCommonPrefix(strs []string) string {
    
	prefix := strs[0]

    if(len(prefix) == 0){
        return ""
    }

	for _, s := range strs{
		
		for len(prefix) > len(s) || prefix[:] != s[:len(prefix)]{
			prefix = prefix[:len(prefix)-1]
		} 

		if len(prefix) == 0{
			break
		}
        
	}

	return prefix
}