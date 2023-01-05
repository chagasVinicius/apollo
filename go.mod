module apollo

go 1.19

require (
	github.com/julienschmidt/httprouter v1.3.0
	handlers v0.0.1
)

replace (
	handlers => ./handlers
)
