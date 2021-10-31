func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    nextL1 := l1.Next
    nextL2 := l2.Next
    
    result := (l1.Val + l2.Val) % 10
    remain := int((l1.Val + l2.Val) / 10)
    
    currentNode := &ListNode{Val: result}
    head := currentNode
    
    for nextL1 != nil || nextL2 != nil {
        if nextL1 == nil {
            nextL1 = &ListNode{Val: 0}            
        }
        if nextL2 == nil {
            nextL2 = &ListNode{Val: 0}            
        }

        result = (remain + nextL1.Val + nextL2.Val) % 10
        remain = int((remain + nextL1.Val + nextL2.Val) / 10)

        nextNode := &ListNode{Val: result}
        currentNode.Next = nextNode
        currentNode = nextNode

        nextL1 = nextL1.Next
        nextL2 = nextL2.Next
    }
    
    if remain != 0 {
        nextNode := &ListNode{Val: remain}
        currentNode.Next = nextNode
        currentNode = nextNode
    }
    return head
}