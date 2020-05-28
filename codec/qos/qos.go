package qos

import "fmt"

// Qos Qos
type Qos uint8

// Qoss
const (
	AtMostOnceDelivery  Qos = 0
	AtLeastOnceDelivery Qos = 1
	ExactlyOnceDelivery Qos = 2
	Reserved                = 3
)

// ParseQos ParseQos
func ParseQos(src uint8) Qos {
	switch Qos(src) {
	case AtMostOnceDelivery:
		return AtMostOnceDelivery
	case AtLeastOnceDelivery:
		return AtLeastOnceDelivery
	case ExactlyOnceDelivery:
		return ExactlyOnceDelivery
	case Reserved:
		panic("reserved qos can't be used")
	default:
		panic(fmt.Sprintf("unknow qos:%d", src))
	}
}

// ToUint8 ToUint8
func (q Qos) ToUint8() uint8 {
	return uint8(q)
}
