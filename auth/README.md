curl -X POST -d '{"authority": 0}' http://192.168.1.253:9000/v1/users/admin?pwd=YWRtaW4=
admin
admin

## create md5 for password
	package main

	import (
	        "fmt"
	        "encoding/base64"
	        "crypto/md5"
	)

	func main() {
	        pw := "333333"

	        sli := md5.Sum([]byte(pw))
	        encypt := base64.StdEncoding.EncodeToString(sli[:])
	        fmt.Println("encoded is: ", encypt)
	}

	result: GhANLA2rGcRDDn1zdis0Iw==

	curl -X POST -d '{"authority": 2}' --header "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QxIn0.tyRrkjOmz-mbdFHBFef_0NIXk_mN5cF9tJ65ShOEvxk" http://192.168.1.253:9000/v1/users/test9?pwd=MTIz