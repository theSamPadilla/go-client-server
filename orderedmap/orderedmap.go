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
// @Returns: (linkedlist.Node.Value, nil) if succesful, (nil, error) otherwise
func (om *OrderedMap) RemoveItemByKey(k interface{}) (interface{}, error) {
	node, exists := om.dict[k]
	if !exists {
		return nil, fmt.Errorf("provided key %s does not exist in the map", k)
	}

	//Else remove node from ll, and remove k from om.dict
	v := node.Value
	om.ll.Remove(node)
	delete(om.dict, k)
	return v, nil
}

// Gets the linkedlist.Node matching the provided k in OrderedMap.
// @Args: k - any type
// @Returns: (*linkedlist.Node, nil) if key exists in om.dict, (nil, error) otherwise.
func (om *OrderedMap) GetItemByKey(k interface{}) (*linkedlist.Node, error) {
	node, exists := om.dict[k]
	if exists {
		return node, nil
	}
	return nil, fmt.Errorf("provided key %s does not exist in the map", k)
}

// Gets the linkedlist.Node.Value matching the provided index i (0-indexed) in OrderedMap.
// @Args: i - unsigned int64
// @Returns: (linkedlist.Node.Value, nil) if om[i] exists, (nil, error) otherwise.
func (om *OrderedMap) GetItemByIndex(i uint64) (*linkedlist.Node, error) {
	//Check for index greater than or equal to length
	if len(om.dict) <= int(i) {
		return nil, fmt.Errorf("provided index %d is greater than or equal to map length %d", i, len(om.dict))
	}

	//Else count fromt the root.next until finding index
	n := om.ll.GetFirstNode()
	for count := 0; count < len(om.dict); count++ {
		if count == int(i) {
			return n, nil
		}
		n = n.GetNextNode()
	}
	return nil, fmt.Errorf("could not find index %d in map", i)
}

// Gets all items in the map
// @Returns: []map[interface{}]interface{} matching linkedlist.Node.Value if om[i] exists, nil otherwise.
func (om *OrderedMap) GetAllItemsInOrder() []map[interface{}]interface{} {
	var r []map[interface{}]interface{}

	//!TODO: BUILD CATCH FOR EMPTY QUERY

	//Iterate through the linkedlist and add key, and value
	n := om.ll.GetFirstNode()
	for count := 0; count < len(om.dict); count++ {
		r = append(r, map[interface{}]interface{}{n.Key: n.Value})

		//Check for last node and break if found
		if om.ll.IsLastNode(n) {
			break
		}
		n = n.GetNextNode()
	}
	return r
}
