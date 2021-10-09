package reverse_double_linked_list

type DoubleLinkedList struct {
	root *Node
}

type Node struct {
	value    int
	previous *Node
	next     *Node
}

func NewDoubleLinkedList(values ...int) *DoubleLinkedList {
	list := &DoubleLinkedList{}
	var node *Node
	for index, value := range values {
		if index == 0 {
			root := &Node{value: value}
			list.root = root
			node = root
		} else {
			newNode := &Node{value: value}
			node.next = newNode
			newNode.previous = node
			node = newNode
		}
	}
	return list
}

func (l *DoubleLinkedList) Values() []int {
	var values = make([]int, 0)

	for node := l.root; node != nil; {
		values = append(values, node.value)
		node = node.next
	}

	return values
}

func (l *DoubleLinkedList) Reverse() *DoubleLinkedList {
	node := l.lastNode()
	if node == nil {
		return &DoubleLinkedList{}
	}
	currentNode := &Node{value: node.value}
	list := &DoubleLinkedList{currentNode}
	for node.previous != nil {
		node = node.previous
		newNode := &Node{value: node.value}
		currentNode.next = newNode
		newNode.previous = currentNode
		currentNode = newNode
	}

	return list
}

func (l *DoubleLinkedList) lastNode() *Node {
	node := l.root
	if node == nil {
		return nil
	}
	for node.next != nil {
		node = node.next
	}
	return node
}
