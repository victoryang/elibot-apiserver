package alarm

import (
	"strings"
	"strconv"
	"fmt"
)

type Record struct {
	Time 			uint32			`json:"time"`
	ErrNo 			[]string 		`json:"errno"`
	Args			[]string 		`json:"args"`
	Msg 			string 			`json:"msg"`
}

func parseLineMessage(line string) (Record, error) {
	list := strings.Split(line, "\x03")
	if len(list) < 6 {
		return Record{}, fmt.Errorf("error errlog data")
	}

	t, err := strconv.ParseUint(list[0], 10, 32)
	if err !=nil {
		return Record{}, err
	}

	last := len(list) - 1
	return Record{
		Time:	uint32(t),
		ErrNo:	list[2:5],
		Args:	list[5:last],
		Msg:	list[last],
	}, nil
}