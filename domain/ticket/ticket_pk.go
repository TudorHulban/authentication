package ticket

import "strconv"

type PrimaryKeyTicket uint64

func (p PrimaryKeyTicket) String() string {
	return strconv.FormatUint(uint64(p), 10)
}
