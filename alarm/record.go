package alarm

import (
	"strings"
	"strconv"

	Log "elibot-apiserver/log"
)

type Record struct {
	Time 			uint32			`json:"time"`
	ErrNo 			[]string 		`json:"errno"`
	Msg 			string 			`json:"msg"`
	Args			[]string 		`json:"args"`
}

func parseLineMessage(line string) Record {
	list := strings.Split(line, "\x03")
	if len(list) < 6 {
		Log.Error("error errlog data")
		return nil
	}

	t, err := strconv.ParseUint(list[0], 10, 32)
	if err !=nil {
		Log.Error("error parse errlog data: time")
		return nil
	}

	return Record{
		Time:	uint32(t),
		ErrNo:	list[2:5],
		Msg:	list[5],
		Args:	list[6:],
	}
}