package str

import "regexp"

// RemoveSymbols function remove all symbols from string
func RemoveSymbols(str string) (string, error) {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		return "", err
	}
	str = re.ReplaceAllString(str, " ")
	return str, nil
}
