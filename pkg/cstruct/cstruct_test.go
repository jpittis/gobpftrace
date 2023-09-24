package cstruct

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type example struct {
	Bytes       []byte
	Int64       int64
	Bool        bool
	XXX_Private int64
}

const (
	converted = `struct example {
    uint8_t* Bytes;
    uint64_t BytesLen;
    uint64_t BytesCap;
    int64_t Int64;
    uint8_t Bool;
    int64_t XXX_Private;
}`
	convertedSkipProtobufPrivate = `struct example {
    uint8_t* Bytes;
    uint64_t BytesLen;
    uint64_t BytesCap;
    int64_t Int64;
    uint8_t Bool;
}`
)

func TestFromGoStruct(t *testing.T) {
	rt := reflect.TypeOf(example{})
	testFromGoStruct(t, rt, false, converted)
	testFromGoStruct(t, rt, true, convertedSkipPrivate)
}

func testFromGoStruct(t *testing.T, rt reflect.Type, skipProtobufPrivate bool, expected string) {
	result, err := FromGoStruct(rt, skipProtobufPrivate)
	require.NoError(t, err)
	require.Equal(t, expected, result)
}
