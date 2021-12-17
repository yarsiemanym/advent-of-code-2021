package day16

import (
	"strconv"
	"strings"

	"github.com/Workiva/go-datastructures/bitarray"
	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type Packet interface {
	Version() uint64
	TypeId() uint64
	Value() uint64
	CheckSum() uint64
}

func ParsePacket(payload bitarray.BitArray) (Packet, bitarray.BitArray) {
	log.Debug("Parsing packet.")

	version, payload := popBits(payload, 3)
	log.Tracef("version = %d", version)

	typeId, payload := popBits(payload, 3)
	log.Tracef("typeId = %d", typeId)

	switch typeId {
	case 4:

		return parseLiteralPacket(version, typeId, payload)
	default:
		return parseOperatorPacket(version, typeId, payload)
	}
}

func hexStringToBitArray(payload string) bitarray.BitArray {
	payload = strings.Trim(payload, " \n")
	size := uint64(len(payload) * 4)
	bitArray := bitarray.NewBitArray(size)

	for index, character := range payload {
		bits, err := strconv.ParseInt(string(character), 16, 8)

		if err != nil {
			log.Fatalf("'%c' is not a valid hexidecimal digit.", character)
		}

		if bits&0b00001000 != 0 {
			position := uint64((4 * index))
			bitArray.SetBit(position)
		}

		if bits&0b00000100 != 0 {
			position := uint64((4 * index) + 1)
			bitArray.SetBit(position)
		}

		if bits&0b00000010 != 0 {
			position := uint64((4 * index) + 2)
			bitArray.SetBit(position)
		}

		if bits&0b00000001 != 0 {
			position := uint64((4 * index) + 3)
			bitArray.SetBit(position)
		}
	}

	return bitArray
}

func popBits(bits bitarray.BitArray, length uint64) (uint64, bitarray.BitArray) {
	frontBits := uint64(0)
	mask := uint64(1)

	for i := uint64(0); i < length; i++ {
		isSet, _ := bits.GetBit(i)
		if isSet {
			frontBits |= mask << (length - i - 1)
		}
	}

	remainingBits := bitarray.NewBitArray(bits.Capacity() - length)

	for i := length; i < bits.Capacity(); i++ {
		isSet, err := bits.GetBit(i)
		common.Check(err)

		if isSet {
			remainingBits.SetBit(i - length)
		}
	}

	return frontBits, remainingBits
}

func popBitsAsBitArray(bits bitarray.BitArray, length uint64) (bitarray.BitArray, bitarray.BitArray) {
	newBits := bitarray.NewBitArray(length)

	for i := uint64(0); i < length; i++ {
		var bit uint64
		bit, bits = popBits(bits, 1)

		if bit != 0 {
			newBits.SetBit(i)
		}
	}

	return newBits, bits
}
