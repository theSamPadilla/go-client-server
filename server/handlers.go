// Implements all the handlers of server package
package server

import (
	"client-server/common"
	"client-server/orderedmap"

	"github.com/rabbitmq/amqp091-go"
)

// Adds an item to the orderedmap
func AddItem(k interface{}, v interface{}, om *orderedmap.OrderedMap, r *amqp091.Delivery, nonSequential bool, sleep bool) {
	sleepIfOn(sleep)

	//Handle request
	common.INFO.Printf("Goroutine ID: %d\n", getGID())
	isNew := om.SetItem(k, v)
	if isNew {
		common.RESULT.Printf("New orderedmap entry: %s->%s\n", k, v)
	} else {
		common.RESULT.Printf("Value of %s reset to %s\n", k, v)
	}

	//Log Goroutine ID and ack
	ackIfSequential(!nonSequential, r)
}

func RemoveItem(k interface{}, om *orderedmap.OrderedMap, r *amqp091.Delivery, nonSequential bool, sleep bool) {
	sleepIfOn(sleep)

	//Handle request
	common.INFO.Printf("Goroutine ID: %d\n", getGID())
	v, err := om.RemoveItemByKey(k)
	if err != nil {
		common.WARN.Printf("error: %s\n", err)
	} else {
		common.RESULT.Printf("Removed %s->%s from ordered map\n", k, v)
	}

	//Log Goroutine ID and ack
	ackIfSequential(!nonSequential, r)
}

func GetItem(k interface{}, om *orderedmap.OrderedMap, r *amqp091.Delivery, nonSequential bool, sleep bool) {
	sleepIfOn(sleep)

	//Handle request
	common.INFO.Printf("Goroutine ID: %d\n", getGID())
	n, err := om.GetItemByKey(k)
	if err != nil {
		common.WARN.Printf("error: %s\n", err)
	} else {
		common.RESULT.Printf("Got %s->%s from ordered map\n", k, n.Value)
	}

	//Log Goroutine ID and ack
	ackIfSequential(!nonSequential, r)
}

func GetItemByIndex(i interface{}, om *orderedmap.OrderedMap, r *amqp091.Delivery, nonSequential bool, sleep bool) {
	sleepIfOn(sleep)

	//Handle request
	common.INFO.Printf("Goroutine ID: %d\n", getGID())
	index, ok := i.(uint64)
	if ok {
		common.WARN.Printf("error processing provided index %s", i)
	}
	n, err := om.GetItemByIndex(index)
	if err != nil { //Message gets acknowledged even on error
		common.WARN.Printf("error: %s\n", err)
	} else {
		common.RESULT.Printf("Got %s from index %s in ordered map\n", n.Value, i)
	}

	//Log Goroutine ID and ack
	ackIfSequential(!nonSequential, r)
}

func GetAllItems(om *orderedmap.OrderedMap, r *amqp091.Delivery, nonSequential bool) {
	//Handle request, never wait
	common.INFO.Printf("Goroutine ID: %d\n", getGID())
	result := om.GetAllItemsInOrder()
	common.RESULT.Printf("Got All Items: %+v\n", result)

	//Log Goroutine ID and ack
	ackIfSequential(!nonSequential, r)
}
