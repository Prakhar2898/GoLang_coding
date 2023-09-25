package main

func checkPr(n []string) bool {
	tempStack := []string{}
	// tempStack = tempStack[:0]
	// fmt.Println(tempStack)
	// fmt.Println(len(tempStack))

	for _, v := range n {

		if v == "(" || v == "{" || v == "[" {
			tempStack = append(tempStack, v)
		}
		if v == ")" || v == "}" || v == "]" {
			l := len(tempStack)
			le := tempStack[l-1]
			if l == 0 {
				return false
			}

			if (v == ")" && le == "(") || (v == "}" && le == "{") || (v == "]" && le == "[") {
				tempStack = tempStack[0 : l-1]
			} else {
				return false
			}
		}
	}

	return true
}
