package day08

import (
	"strings"

	"github.com/Workiva/go-datastructures/bitarray"
	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type signalMapper struct {
	signalMap map[rune]rune
}

func NewSignalMapper(uniqueSignalPatterns []string) *signalMapper {
	return &signalMapper{
		signalMap: constructSignalMap(uniqueSignalPatterns),
	}
}

func constructSignalMap(uniqueSignalPatterns []string) map[rune]rune {
	var zeroBits bitarray.BitArray
	var oneBits bitarray.BitArray
	var twoBits bitarray.BitArray
	var threeBits bitarray.BitArray
	var fourBits bitarray.BitArray
	var fiveBits bitarray.BitArray
	var sixBits bitarray.BitArray
	var sevenBits bitarray.BitArray
	var eightBits bitarray.BitArray
	var nineBits bitarray.BitArray

	// First pass
	for _, uniqueSignalPattern := range uniqueSignalPatterns {
		length := len(uniqueSignalPattern)

		switch length {
		case 2: // 1
			log.Debug("'1' signal pattern deciphered.")
			log.Tracef("signalPattern = \"%v\"", uniqueSignalPattern)
			oneBits = signalPatternToBitArray(uniqueSignalPattern)
		case 3: // 7
			log.Debug("'7' signal pattern deciphered.")
			log.Tracef("signalPattern = \"%v\"", uniqueSignalPattern)
			sevenBits = signalPatternToBitArray(uniqueSignalPattern)
		case 4: // 4
			log.Debug("'4' signal pattern deciphered.")
			log.Tracef("signalPattern = \"%v\"", uniqueSignalPattern)
			fourBits = signalPatternToBitArray(uniqueSignalPattern)
		case 7: // 8
			log.Debug("'8' signal pattern deciphered.")
			log.Tracef("signalPattern = \"%v\"", uniqueSignalPattern)
			eightBits = signalPatternToBitArray(uniqueSignalPattern)
		}
	}
	// 1, 4, 7, and 8 are deciphered.

	// Second pass
	for _, uniqueSignalPattern := range uniqueSignalPatterns {
		length := len(uniqueSignalPattern)

		switch length {
		case 5: // 2, 3, 5
			unknownBits := signalPatternToBitArray(uniqueSignalPattern)

			if unknownBits.And(oneBits).Equals(oneBits) { // 3
				log.Debug("'3' signal pattern deciphered.")
				log.Tracef("signalPattern = \"%v\"", uniqueSignalPattern)
				threeBits = unknownBits
			}

		case 6: // 0, 6, 9
			unknownBits := signalPatternToBitArray(uniqueSignalPattern)

			if unknownBits.And(fourBits).Equals(fourBits) { // 9
				log.Debug("'9' signal pattern deciphered.")
				log.Tracef("signalPattern = \"%v\"", uniqueSignalPattern)
				nineBits = unknownBits
			} else if unknownBits.And(sevenBits).Equals(sevenBits) { // 0
				log.Debug("'0' signal pattern deciphered.")
				log.Tracef("signalPattern = \"%v\"", uniqueSignalPattern)
				zeroBits = unknownBits
			} else { // 6
				log.Debug("'6' signal pattern deciphered.")
				log.Tracef("signalPattern = \"%v\"", uniqueSignalPattern)
				sixBits = unknownBits
			}
		}
	}
	// 0, 1, 3, 4, 6, 7, 8, and 9 are deciphered.

	// Third pass
	for _, uniqueSignalPattern := range uniqueSignalPatterns {
		length := len(uniqueSignalPattern)

		switch length {
		case 5: // 2, 3, 5
			unknownBits := signalPatternToBitArray(uniqueSignalPattern)

			if !unknownBits.Equals(threeBits) {
				if unknownBits.And(nineBits).Equals(unknownBits) { // 5
					log.Debug("'5' signal pattern deciphered.")
					log.Tracef("signalPattern = \"%v\"", uniqueSignalPattern)
					fiveBits = unknownBits
				} else {
					log.Debug("'2' signal pattern deciphered.") // 2
					log.Tracef("signalPattern = \"%v\"", uniqueSignalPattern)
					twoBits = unknownBits
				}
			}
		}
	}
	// All digits are deciphered.

	aBit := xor(sevenBits, oneBits)
	bBit := xor(nineBits, threeBits)
	cBit := xor(eightBits, sixBits)
	dBit := xor(eightBits, zeroBits)
	eBit := xor(sixBits, fiveBits)
	fBit := xor(zeroBits, twoBits).And(oneBits)
	gBit := xor(xor(nineBits, fourBits), aBit)

	// Create map
	signalMap := map[rune]rune{
		rune(bitArrayToSignalPattern(aBit)[0]): 'a',
		rune(bitArrayToSignalPattern(bBit)[0]): 'b',
		rune(bitArrayToSignalPattern(cBit)[0]): 'c',
		rune(bitArrayToSignalPattern(dBit)[0]): 'd',
		rune(bitArrayToSignalPattern(eBit)[0]): 'e',
		rune(bitArrayToSignalPattern(fBit)[0]): 'f',
		rune(bitArrayToSignalPattern(gBit)[0]): 'g',
	}

	for _, signal := range "abcdefg" {
		log.Tracef("Signal '%c' maps to signal '%c'.", signal, signalMap[signal])
	}

	return signalMap
}

func xor(a bitarray.BitArray, b bitarray.BitArray) bitarray.BitArray {
	max := bitarray.NewBitArray(7, true)
	notA := max.Nand(a)
	notB := max.Nand(b)
	return (a.And(notB)).Or(b.And(notA))
}

func signalPatternToBitArray(signalPattern string) bitarray.BitArray {
	bits := bitarray.NewBitArray(7)

	for index, signal := range "abcdefg" {
		if strings.ContainsRune(signalPattern, signal) {
			bits.SetBit(uint64(7 - index))
		}
	}

	return bits
}

func bitArrayToSignalPattern(bits bitarray.BitArray) string {
	signalPattern := ""

	for index, signal := range "abcdefg" {
		bit, err := bits.GetBit(uint64(7 - index))
		common.Check(err)

		if bit {
			signalPattern += string(signal)
		}
	}

	return signalPattern
}

func (signalMapper *signalMapper) MapSignals(scambledSignals string) string {
	unscrabledSignals := ""

	for _, scambledSignal := range scambledSignals {
		unscrabledSignal, exists := signalMapper.signalMap[scambledSignal]

		if exists {
			unscrabledSignals += string(unscrabledSignal)
		} else {
			log.Fatalf("No mapping exists for signal '%c'.", scambledSignal)
		}
	}

	return unscrabledSignals
}
