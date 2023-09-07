package linklist

type Node struct {
	Data int
	Next *Node
}

type SingleLinkList struct {
	head *Node
}

func NewLinkList() *SingleLinkList {
	return &SingleLinkList{
		head: nil,
	}
}

func (list *SingleLinkList) Count() int {
	p := list.head
	len := 0
	for p != nil {
		len++
		p = p.Next
	}
	return len
}

func (list *SingleLinkList) Append(v int) {
	n := &Node{
		Data: v,
		Next: nil,
	}
	if list.head == nil {
		list.head = n
	} else {
		p := list.head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = n
	}
}

func (list *SingleLinkList) Find(v int) *Node {
	p := list.head
	for p != nil && p.Data != v {
		p = p.Next
	}

	return p
}
