package day16

import (
	"testing"

	"github.com/Workiva/go-datastructures/bitarray"
)

func Test_hexStringToBitArray(t *testing.T) {
	payload := "89AB" // 10001001 10101011
	expectedBitArray := bitarray.NewBitArray(16)
	expectedBitArray.SetBit(0)
	expectedBitArray.SetBit(4)
	expectedBitArray.SetBit(7)
	expectedBitArray.SetBit(8)
	expectedBitArray.SetBit(10)
	expectedBitArray.SetBit(12)
	expectedBitArray.SetBit(14)
	expectedBitArray.SetBit(15)
	actualBitArray := hexStringToBitArray(payload)

	if !actualBitArray.Equals(expectedBitArray) {
		t.Error("Bit array does not match.")
	}
}

func Test_nextByte(t *testing.T) {
	text := "D2FE28"
	bits := hexStringToBitArray(text)
	version, bits := popBits(bits, 3)
	typeId, bits := popBits(bits, 3)

	if version != 6 {
		t.Errorf("Expected 6 but got %d.", version)
	}

	if typeId != 4 {
		t.Errorf("Expected 4 but got %d.", typeId)
	}
}

func Test_ParsePacket_Literal(t *testing.T) {
	text := "D2FE28"
	bits := hexStringToBitArray(text)
	packet, _ := ParsePacket(bits)

	if packet.Version() != 6 {
		t.Errorf("Expected 6 but got %d.", packet.Version())
	}

	if packet.TypeId() != 4 {
		t.Errorf("Expected 4 but got %d.", packet.TypeId())
	}

	if packet.Value() != 2021 {
		t.Errorf("Expected 2021 but got %d.", packet.Value())
	}
}
