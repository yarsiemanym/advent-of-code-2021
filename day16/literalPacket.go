package day16

import (
	"github.com/Workiva/go-datastructures/bitarray"
	log "github.com/sirupsen/logrus"
)

type literalPacket struct {
	version uint64
	typeId  uint64
	value   uint64
}

func (packet literalPacket) Version() uint64 {
	return packet.version
}

func (packet literalPacket) TypeId() uint64 {
	return packet.typeId
}

func (packet literalPacket) Value() uint64 {
	return packet.value
}

func (packet literalPacket) VersionSum() uint64 {
	return packet.version
}

func parseLiteralPacket(version uint64, typeId uint64, payload bitarray.BitArray) (Packet, bitarray.BitArray) {
	log.Debug("Parsing literal packet.")
	log.Tracef("version = %d", version)
	log.Tracef("typeId = %d", typeId)

	isLastGroup := false
	groupValues := []uint64{}
	for !isLastGroup {
		log.Debug("Parsing value group.")
		var bits uint64
		var groupValue uint64

		bits, payload = popBits(payload, 1)
		isLastGroup = bits == 0
		log.Tracef("isLastGroup = %v", isLastGroup)

		groupValue, payload = popBits(payload, 4)
		log.Tracef("groupValue = %d", groupValue)
		groupValues = append(groupValues, groupValue)
	}

	literalValue := uint64(0)
	for _, groupValue := range groupValues {
		literalValue = (literalValue << 4) + groupValue
	}

	literalPacket := &literalPacket{
		version: version,
		typeId:  typeId,
		value:   literalValue,
	}

	return literalPacket, payload
}
