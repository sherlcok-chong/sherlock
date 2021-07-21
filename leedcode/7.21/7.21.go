package __21

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	rea, reb := headA, headB
	for rea != reb {
		if rea == nil {
			rea = headB
		} else {
			rea = rea.Next
		}
		if reb == nil {
			reb = headA
		} else {
			reb= reb.Next
		}
	}
	return reb
}
