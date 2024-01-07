func isValid(s string) bool {

	size := len(s)

	if size == 1 {
		return false
	}

	brackets := []uint8{s[0]}

	for i := 1; i <= size-1; i++ {
		switch string(s[i]) {
		case ")":
			if len(brackets) == 0 {
				return false
			}
			if brackets[len(brackets)-1] != 40 {
				return false
			} else {
				brackets = brackets[:len(brackets)-1]
			}
		case "}":
			if len(brackets) == 0 {
				return false
			}
			if brackets[len(brackets)-1] != 123 {
				return false
			} else {
				brackets = brackets[:len(brackets)-1]
			}
		case "]":
			if len(brackets) == 0 {
				return false
			}
			if brackets[len(brackets)-1] != 91 {
				return false
			} else {
				brackets = brackets[:len(brackets)-1]
			}
		default:
			{
				brackets = append(brackets, s[i])
			}
		}
	}

	if len(brackets) != 0 {
		return false
	}

	return true

}