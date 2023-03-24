// Package linkedlist is a custom implementation of a linkedlist for use
// by orderedmap [client-server/orderedmap]
//
// The linkedlist package implements a two-sided pointer linked list where
// the `previous` pointer of root points to the last node of the linked list.
package linkedlist

// Value of the ordered map and a node of a linked list to keep order of elements
type Node struct {
	//Exported values
	Key   interface{}
	Value interface{}

	previous, next *Node //Internal pointers
}

// Gets next node of the linked list
func (n *Node) GetNextNode() *Node {
	return n.next
}

// A LinkedList instantiates the root of the LinkedList on a `nil` Node.
// The Zero value of LinkedList `nil` is good to use as the root.
type LinkedList struct {
	Root Node
}

// Gets last node of the linked list
func (ll *LinkedList) GetLastNode() *Node {
	return ll.Root.previous
}

// Gets first node of the linked list
func (ll *LinkedList) GetFirstNode() *Node {
	return ll.Root.next
}

// Returns ture if the node is the last in the
// linked list, false otherwise
func (ll *LinkedList) IsLastNode(n *Node) bool {
	return ll.Root.previous == n
}

// Pushes an element to the end of the linked list
// @Args: key, value - any type.
// @Returns: new Node n *Node with key and value, n.previous pointing to l.root.prev and
// l.root.prev pointing to n.
func (ll *LinkedList) Append(key interface{}, value interface{}) *Node {
	n := &Node{Key: key, Value: value}
	//First appending -> previous and next to n
	if ll.Root.previous == nil {
		ll.Root.previous = n
		ll.Root.next = n
		return n
	}

	// Point previous of n to last, next of last to n, and previous of root to n
	last := ll.GetLastNode()
	n.previous = last
	last.next = n
	ll.Root.previous = n
	return n

}

// Removes an element e from the linked list
// @Args: n of typte *Node
func (ll *LinkedList) Remove(n *Node) {
	switch {
	//Removing the first element (after root) -> Point root.next to n.next
	case n.previous == nil:
		ll.Root.next = n.next
	//Removing tail -> Repoint root.previous to n.previous and n.previous.next to nil
	case n.next == nil:
		ll.Root.previous = n.previous
		n.previous.next = nil
	//Middle node -> n.previous.next to n.next, and n.next.previous to n.previous
	default:
		n.previous.next = n.next
		n.next.previous = n.previous
	}

	//Clean memory
	n.previous = nil
	n.next = nil
	n = nil
}
