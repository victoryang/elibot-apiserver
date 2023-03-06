module elibot-apiserver

go 1.12

require (
	github.com/dgrijalva/jwt-go v0.0.0-20180308231308-06ea1031745c
	github.com/fsnotify/fsnotify v1.4.7
	github.com/golang/protobuf v1.2.0
	github.com/gorilla/mux v1.6.2
	github.com/gorilla/websocket v1.4.1
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mattn/go-sqlite3 v0.0.0-20180719091609-b3511bfdd742
	github.com/rs/cors v1.6.0
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/cobra v0.0.0-20180829160745-99dc12355885
	github.com/spf13/pflag v0.0.0-20180831151432-298182f68c66 // indirect
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/urfave/negroni v0.0.0-20180130044549-22c5532ea862
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v1.19.0
	gopkg.in/yaml.v2 v2.2.1
)

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190402142545-baf5eb976a8c

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190325154230-a5d413f7728c

replace golang.org/x/net => github.com/golang/net v0.0.0-20190328230028-74de082e2cca

replace golang.org/x/text => github.com/golang/text v0.3.0

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.19.1

replace golang.org/x/tools => github.com/golang/tools v0.0.0-20190402200628-202502a5a924

replace golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a

replace golang.org/x/lint => github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3

replace golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6

replace google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190401181712-f467c93bbac2

replace golang.org/x/exp => github.com/golang/exp v0.0.0-20190402192236-7fd597ecf556

replace cloud.google.com/go => github.com/googleapis/google-cloud-go v0.37.2

replace google.golang.org/appengine => github.com/golang/appengine v1.5.0

replace golang.org/x/image => github.com/golang/image v0.0.0-20190321063152-3fc05d484e9f

replace golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190327163128-167ebed0ec6d

replace golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4

replace golang.org/x/build => github.com/golang/build v0.0.0-20190403045414-85a73d7451e7

replace google.golang.org/api => github.com/googleapis/google-api-go-client v0.3.0

replace golang.org/x/perf => github.com/golang/perf v0.0.0-20190312170614-0655857e383f
