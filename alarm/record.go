package alarm

import (
	"strings"
	"strconv"
	"fmt"

	Log "elibot-apiserver/log"
)

type Record struct {
	Time 			uint32			`json:"time"`
	ErrNo 			[]string 		`json:"errno"`
	Msg 			string 			`json:"msg"`
	Args			[]string 		`json:"args"`
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

	return Record{
		Time:	uint32(t),
		ErrNo:	list[2:5],
		Msg:	list[5],
		Args:	list[6:],
	}, nil
}