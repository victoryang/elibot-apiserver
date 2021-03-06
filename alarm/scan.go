package alarm

import (
	"os"
	"bufio"
	Log "elibot-apiserver/log"
)

func ScanAndParse(input string) []Record {
	fd, err := os.Open(input)
	if err!=nil {
		Log.Error("Read error: ", err)
		return nil
	}
	defer fd.Close()

	var rec []Record
	scanner := bufio.NewScanner(fd)
    for scanner.Scan() {
        if r,err := parseLineMessage(scanner.Text()); err==nil {
        	rec = append(rec, r)
        }
    }
    return rec
}