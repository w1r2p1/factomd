// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package interfaces

import (
	"time"
)

//A structure for handling timestamps for messages
type ITimestamp interface {
	SetTimeNow()
	SetTime(miliseconds uint64)
	GetTime() time.Time
	UnmarshalBinaryData(data []byte) (newData []byte, err error)
	UnmarshalBinary(data []byte) error
	GetTimeSeconds() int64
	GetTimeMinutesUInt32() uint32
	GetTimeMilli() int64
	GetTimeMilliUInt64() uint64
	GetTimeSecondsUInt32() uint32
	MarshalBinary() ([]byte, error)
	String() string
}
