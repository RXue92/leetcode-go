//beat 0.89%
func lengthOfLongestSubstring(s string) int {
	var longest int
	var out string
	for s != "" {
		substring, length := lengthOfLongestSubstringFromHead(s)
		if length > longest {
			longest = length
			out = substring
		}
		s = s[1:]
	}

	fmt.Println(out)
	return longest
}


func lengthOfLongestSubstringFromHead(s string) (string, int) {
	//assume s != ""


	comp := make(map[string]int)
	var out []string
	var length int
	var head string
	l := len(s)
	for i := 0; i < l; i++ {
		head = s[:1]
		if _, ok := comp[head]; !ok {
			comp[head] = i
			out = append(out, head)
			length = i + 1
			s = s[1:]
		} else {
			break
		}
	}
	str := strings.Join(out, "")
	return str, length
}
