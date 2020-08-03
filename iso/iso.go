package iso

import "regexp"

// Val ...
func Val(code string, min, max int, re *regexp.Regexp) bool {
	r := len(code) >= min && len(code) <= max
	if re != nil {
		return r && re.MatchString(code)
	}
	return r
}
