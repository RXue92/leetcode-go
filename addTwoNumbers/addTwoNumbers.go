/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var out *ListNode
    out = getdigits(l1)
    return out
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