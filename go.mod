module apollo

go 1.19

require (
	github.com/julienschmidt/httprouter v1.3.0
	github.com/zmb3/spotify/v2 v2.3.1
	golang.org/x/oauth2 v0.4.0
	handlers v0.0.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.5.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace handlers => ./handlers
