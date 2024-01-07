func isPalindrome(x int) bool {
    
    if x < 0{
        return false
    }

    if x < 10{
        return true
    }

    aux := x
    reverse := 0

    for aux > 0{
        reverse = reverse*10 + aux%10
        aux = aux/10
    }

	return x == reverse
}