import "strings"
var interboard = map[int]string {
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkll",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func makeboard() map[int][]string {
    board := make(map[int][]string)
	for k, v := range(interboard) {
		board[k] = strings.Split(v, "")
	}

	return board
}

func letterCombinations(digits string) []string {
		//out := []string{}
		board := makeboard()
		//remove 0,1
		new_digits := strings.Replace(digits, "0", "", -1)
		new_digits = strings.Replace(new_digits, "1", "", -1)

		//check if nil
		if len(new_digits) == 0 {
			return out
		}

		//head
		num_set := strings.Split(new_digits, "")
		//head := []string{"", "", "", ""}
		head := board[num_set[0]]

		//multiply
		if len(num_set) == 1 {
			return head
		}

		for i := 1; i < len(num_set); i++ {
				tail := board[num_set[i]]
				for m := 0; m < len(head); m++ {
					for n := 0; n < len(tail); n++ {
						head = append(head, head[m] + tail[n])
						}
					}
				head = head[m:]
		}
		return head
}
