module github.com/micro/services/platform/api

go 1.13

require (
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro/v2 v2.7.1-0.20200527112433-192f548c8304
	github.com/micro/services/platform/service v0.0.0-20200313185528-4a795857eb73
	github.com/micro/services/users/service v0.0.0-20200402122209-bbd3453477a3
)

replace github.com/micro/services/platform/service => ../service
