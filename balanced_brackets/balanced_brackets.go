package balanced_brackets

/*
 * Complete the 'isBalanced' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */
func isBalanced(s string) string {
	matchFound := true
	for matchFound == true {
		matchFound = false
		for i := 0; i < len(s)-1; i++ {
			//check if there exists an opening and closing neighboring pair and delete it from the string
			if (s[i] == '{' && s[i+1] == '}') || (s[i] == '[' && s[i+1] == ']') || (s[i] == '(' && s[i+1] == ')') {
				test := []rune(s)
				s = string(append(test[0:i], test[i+2:]...))
				matchFound = true
			}
		}
	}
	if len(s) == 0 {
		return "YES"
	}

	return "NO"
}
