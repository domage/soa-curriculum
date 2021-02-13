module github.com/domage/grpc-queue/frontend-go

go 1.15

require (
	github.com/domage/grpc-queue/count v0.0.0-00010101000000-000000000000
	github.com/streadway/amqp v1.0.0
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)

// Для того, чтобы go мог добраться до библиотек, сгенерированных protoc,
// располагающихся в папке "./count" необходимо указать ему, где искать
// модуль, который формируется на основе .proto файла (см. директиву 
// option go_package в "../pkg/proto/count"
replace github.com/domage/grpc-queue/count => ./count
