//beat 15%
func convert(s string, numRows int) string {
	strs := strings.Split(s, "")
	mymap := make(map[int][]string)
	for i := 0; i < len(strs); i++ {
		j := getrow(i, numRows)
		_, ok := mymap[j]
		if !ok {
			mymap[j] = []string{strs[i]}
		} else {
			mymap[j] = append(mymap[j], strs[i])
		}
	}
	out := []string{}
	for k := 1; k <= numRows; k++ {
		out = append(out, mymap[k]...)
	}
	return strings.Join(out, "")
}

func getrow(i int, numRows int) int {
	group := 2*numRows - 2
	if group == 0 {
		return 1
	}
	i = i % group
	var out int

	if i < numRows {
		out = i + 1
	} else {
		out = 2*numRows - 1 - i
	}
	return out
}
