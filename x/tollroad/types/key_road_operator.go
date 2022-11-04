package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RoadOperatorKeyPrefix is the prefix to retrieve all RoadOperator
	RoadOperatorKeyPrefix = "RoadOperator/value/"
)

// RoadOperatorKey returns the store key to retrieve a RoadOperator from the index fields
func RoadOperatorKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
