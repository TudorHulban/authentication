package helpers

import "strconv"

type PrimaryKey uint64

func (pk PrimaryKey) String() string {
	return strconv.FormatUint(
		uint64(pk),
		10,
	)
}

const PrimaryKeyZero = PrimaryKey(0)
