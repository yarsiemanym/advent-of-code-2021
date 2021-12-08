package day08

import (
	"reflect"
	"testing"

	"github.com/Workiva/go-datastructures/bitarray"
)

func Test_signalMapper_Xor(t *testing.T) {
	a := bitarray.NewBitArray(8)
	a.SetBit(1)
	b := bitarray.NewBitArray(8)
	b.SetBit(1)
	b.SetBit(2)

	expectedBits := bitarray.NewBitArray(8)
	expectedBits.SetBit(2)

	actualBits := xor(a, b)

	if !actualBits.Equals(expectedBits) {
		t.Errorf("Expected %v but got %v.", expectedBits, actualBits)
	}
}

func Test_signalMapper_signalPatternToBitArray(t *testing.T) {
	expectedBits := bitarray.NewBitArray(7)
	expectedBits.SetBit(7)
	expectedBits.SetBit(6)
	expectedBits.SetBit(5)
	expectedBits.SetBit(3)
	expectedBits.SetBit(2)
	expectedBits.SetBit(1)

	actualBits := signalPatternToBitArray(zero)

	if !expectedBits.Equals(actualBits) {
		t.Errorf("Expected %v but got %v.", expectedBits, actualBits)
	}
}

func Test_signalMapper_bitArrayToSignalPattern(t *testing.T) {
	bits := bitarray.NewBitArray(7)
	bits.SetBit(7)
	bits.SetBit(6)
	bits.SetBit(5)
	bits.SetBit(3)
	bits.SetBit(2)
	bits.SetBit(1)

	actualSignalPattern := bitArrayToSignalPattern(bits)

	if actualSignalPattern != zero {
		t.Errorf("Expected \"%v\" but got \"%v\".", zero, actualSignalPattern)
	}
}

func Test_signalMapper_constructSignalMap(t *testing.T) {
	signalPatterns := []string{
		"acedgfb",
		"cdfbe",
		"gcdfa",
		"fbcad",
		"dab",
		"cefabd",
		"cdfgeb",
		"eafb",
		"cagedb",
		"ab",
	}

	expectedSignalMap := map[rune]rune{
		'd': 'a',
		'e': 'b',
		'a': 'c',
		'f': 'd',
		'g': 'e',
		'b': 'f',
		'c': 'g',
	}

	actualSignalMap := constructSignalMap(signalPatterns)

	if !reflect.DeepEqual(actualSignalMap, expectedSignalMap) {
		t.Errorf("Expected %v but got %v.", expectedSignalMap, actualSignalMap)
	}
}

func Test_signalMapper_MapSignalPattern(t *testing.T) {
	signalPatterns := []string{
		"acedgfb",
		"cdfbe",
		"gcdfa",
		"fbcad",
		"dab",
		"cefabd",
		"cdfgeb",
		"eafb",
		"cagedb",
		"ab",
	}

	signalMapper := NewSignalMapper(signalPatterns)

	signal := signalMapper.MapSignals("cdfeb")

	if signal != "gadbf" {
		t.Errorf("Expected \"gadbf\" but got \"%v\".", signal)
	}

	signal = signalMapper.MapSignals("fcadb")

	if signal != "dgcaf" {
		t.Errorf("Expected \"dgcaf\" but got \"%v\".", signal)
	}
}
