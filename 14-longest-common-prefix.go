func longestCommonPrefix(strs []string) string {
    prefix := strs[0]

	for _, s := range strs{

		toLarge:
		if len(prefix) > len(s){
			prefix = prefix[:len(prefix)-1]
			goto toLarge
		}

		findPrefix:
		if len(prefix) == 0{
			break
		}
		if prefix[:] == s[:len(prefix)]{
			continue
		}else{
			prefix = prefix[:len(prefix)-1]
			goto findPrefix
		}

	}

	return prefix
}