package orderedmap

import (
	"client-server/linkedlist"
	"fmt"
)

// An OrderedMap is a map of keys k to *linkedlist.Node instances.
// Each linkedlinst.Node instance in turn has a key and value, and previous and next pointers.
// Nodes are appended to linkedlist.LinkedList to track order
type OrderedMap struct {
	dict map[interface{}]*linkedlist.Node
	ll   linkedlist.LinkedList
}

// Constructs an empty ordered map
func Constructor() *OrderedMap {
	return &OrderedMap{ //No need to instantiate ll, nil type works
		dict: make(map[interface{}]*linkedlist.Node),
	}
}

// Sets item in ordered map. If key exists, it resets the value.
// @Args: k, v - any type.
// @Returns: true if a new k->v pair was set, false if k[v] was reset
func (om *OrderedMap) SetItem(k interface{}, v interface{}) bool {
	_, exists := om.dict[k]
	if exists {
		om.dict[k].Value = v
		return false
	}

	//Else add k, v to linked list, and append that element to map
	newNode := om.ll.Append(k, v)
	om.dict[k] = newNode
	return true
}

// Removes key from ordered map.
// @Args: k - any type
// @Returns: nil if remove was succesful, error otherwise
func (om *OrderedMap) RemoveItemByKey(k interface{}) error {
	node, exists := om.dict[k]
	if !exists {
		return fmt.Errorf("provided key %s does not exist in the map", k)
	}

	//Else remove node from ll, and remove k from om.dict
	om.ll.Remove(node)
	delete(om.dict, k)
	return nil
}

// Gets the linkedlist.Node matching the provided k in OrderedMap.
// @Args: k - any type
// @Returns: *linkedlist.Node if key exists in om.dict, nil otherwise.
func (om *OrderedMap) GetItemByKey(k interface{}) *linkedlist.Node {
	node, exists := om.dict[k]
	if exists {
		return node
	}

	return nil
}

// Gets the linkedlist.Node.Value matching the provided index i (0-indexed) in OrderedMap.
// @Args: i - unsigned int32
// @Returns: interface{} matching linkedlist.Node.Value if om[i] exists, nil otherwise.
func (om *OrderedMap) GetItemByIndex(i uint32) interface{} {
	//Check for index greater than or equal to length
	if len(om.dict) >= int(i) {
		return nil
	}

	//Else count fromt the root.next until finding index
	n := *om.ll.GetFirstNode()
	for count := 0; count < len(om.dict); count++ {
		if count == int(i) {
			return n.Value
		}
		n = *n.GetNextNode()
	}
	return nil
}
