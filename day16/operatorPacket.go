package day16

import (
	"math"

	"github.com/Workiva/go-datastructures/bitarray"
	log "github.com/sirupsen/logrus"
)

type operatorPacket struct {
	version    uint64
	typeId     uint64
	subPackets []Packet
}

func (packet operatorPacket) Version() uint64 {
	return packet.version
}

func (packet operatorPacket) TypeId() uint64 {
	return packet.typeId
}

func (packet operatorPacket) SubPackets() []Packet {
	return packet.subPackets
}

func (packet operatorPacket) Value() uint64 {
	switch packet.typeId {
	case 0: // sum
		sum := uint64(0)
		for _, subPacket := range packet.subPackets {
			sum += subPacket.Value()
		}
		return sum
	case 1: // product
		product := uint64(1)
		for _, subPacket := range packet.subPackets {
			product *= subPacket.Value()
		}
		return product
	case 2: // minimum
		minimum := uint64(math.MaxUint64)
		for _, subPacket := range packet.subPackets {
			value := subPacket.Value()
			if value < minimum {
				minimum = value
			}
		}
		return minimum
	case 3: // maximum
		maximum := uint64(0)
		for _, subPacket := range packet.subPackets {
			value := subPacket.Value()
			if value > maximum {
				maximum = value
			}
		}
		return maximum
	case 5: // greater than
		if packet.subPackets[0].Value() > packet.subPackets[1].Value() {
			return 1
		} else {
			return 0
		}
	case 6: // less than
		if packet.subPackets[0].Value() < packet.subPackets[1].Value() {
			return 1
		} else {
			return 0
		}
	case 7: // equal to
		if packet.subPackets[0].Value() == packet.subPackets[1].Value() {
			return 1
		} else {
			return 0
		}
	default:
		log.Panicf("%d is not a valid length type Id.", packet.typeId)
		return math.MaxUint64
	}
}

func (packet operatorPacket) VersionSum() uint64 {
	sum := uint64(packet.version)

	for _, subPacket := range packet.subPackets {
		sum += subPacket.VersionSum()
	}

	return sum
}

func parseOperatorPacket(version uint64, typeId uint64, payload bitarray.BitArray) (Packet, bitarray.BitArray) {
	log.Debug("Parsing oerator packet.")
	log.Tracef("version = %d", version)
	log.Tracef("typeId = %d", typeId)

	var lengthTypeId uint64
	lengthTypeId, payload = popBits(payload, 1)

	switch lengthTypeId {
	case 0: // total length of subpacket bits
		var totalLength uint64
		totalLength, payload = popBits(payload, 15)

		var subPacketsPayload bitarray.BitArray
		subPacketsPayload, payload = popBitsAsBitArray(payload, totalLength)

		subPackets := []Packet{}
		for !subPacketsPayload.IsEmpty() {
			var subPacket Packet
			subPacket, subPacketsPayload = ParsePacket(subPacketsPayload)
			subPackets = append(subPackets, subPacket)
		}

		packet := operatorPacket{
			version:    version,
			typeId:     typeId,
			subPackets: subPackets,
		}

		return packet, payload
	case 1: // number of subpackets
		var subPacketCount uint64
		subPacketCount, payload = popBits(payload, 11)

		subPackets := []Packet{}
		for i := uint64(0); i < subPacketCount; i++ {
			var subPacket Packet
			subPacket, payload = ParsePacket(payload)
			subPackets = append(subPackets, subPacket)
		}

		packet := operatorPacket{
			version:    version,
			typeId:     typeId,
			subPackets: subPackets,
		}

		return packet, payload
	default:
		log.Panicf("%d is not a valid length type Id.", lengthTypeId)
		return literalPacket{}, bitarray.NewBitArray(0)
	}
}
