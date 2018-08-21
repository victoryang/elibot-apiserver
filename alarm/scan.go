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
        r := parseLineMessage(scanner.Text())
        rec = append(rec, r)
    }
    return rec
}