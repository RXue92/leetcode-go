/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    s1 := getdigits(l1)
    s2 := getdigits(l2)
    s := adder(s1, s2)
    out := parser(s)
    return out
}

func parser(s []int) *ListNode {

}



func adder(s1 []int, s2 []int) []int {
    n1 := toNum(s1)
    n2 := toNum(s2)
    n := n1 + n2

    out := []int{0}
    for n > 0 {
        out = append(out, n%10)
        n = n / 10
    }

    if len(out) > 1 {
        return out[1:] 
    } else {
        return out
    }
}

func toNum(s []int) int {
    base := 1
    var sum int
    for _, v := range(s) {
        sum = sum + v*base
        base = base * 10
    }
    return sum
}

func getdigits(l *ListNode) []int {
    out := []int
    out = append(out, l.Val)
    temp := l.Next
    for l.Next != nil {
        out = append(out, temp.Val)
        if temp.Next != nil {
            temp = temp.Next
        } else {
            break
        }
    }
    return out
}