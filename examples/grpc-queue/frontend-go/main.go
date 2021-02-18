// Сервер который:
// 1. Принимает запросы QueueProcess по gRPC на обработку строк
//    1.1. В ответ возвращает QResultId: идентификатор ответа для получения результата в дальнейшем
// 2. Передет запросы на обработку строк в виде protobuf-сообщений QProcessMessage в очередь rabbitMQ
// 3. Получает из очереди результаты обработки QResultMessage и сохраняет их в словаре
// 4. По запросу QueueGetResult от клиента, содержащим QResultId, возвращает результат
//
// Для работы с protobuf и RabbitMQ в Go необходимо:
// 1. Установить protoc
// 2. Инициализировать текущую директорию для работы с модулями:
//    go mod init github.com/domage/grpc-queue/frontend-go
// 3. Для того, чтобы сгенерированные protoc модули были доступны, добавить в
//    файл go.mod ссылку на папку ./count/
//    replace github.com/domage/grpc-queue/count => ./count
//
// При изменении протокола обмена данными (файла ../pkg/proto/count/count.proto),
// для сборки сериализатора/десериализатора (располагаются в папке count)
// необходимо из текущей папки выполнить команду:
// protoc  -I ../pkg/proto/count --go_out=./count \
//         --go_opt=paths=source_relative --go-grpc_out=./count \
//         --go-grpc_opt=paths=source_relative ../pkg/proto/count/count.proto
//
// Запуск: go run main.go
//
// Дополнительная информация:
// - Использование protobuf и gRPC в Go:
//		https://grpc.io/docs/languages/go/basics/
//
// - Работа с RabbitMQ:
//      https://www.rabbitmq.com/tutorials/tutorial-one-go.html
//      https://www.rabbitmq.com/tutorials/tutorial-two-go.html
//		https://medium.com/@masnun/work-queue-with-go-and-rabbitmq-b8c295cde861

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/domage/grpc-queue/count"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type server struct {
	count.UnimplementedCountServer
}

// Счетчик для индексации полученных запросов
var respCounter int32

// Словарь для выдачи результатов обработки запросов
var responses map[int32]count.QResultMessage

// Очереди для отправки запросов на обработку данных и получение результатов
var qReq, qResp amqp.Queue

// Канал подключения к очереди сообщений
var ch *amqp.Channel

// Соединение с очередью сообщений
var conn *amqp.Connection

// gRPC-функция, получающая текст, в котором надо подсчитать количество слов.
// Направляет сообщение в очередь отправки запросов в очереди сообщений.
// Возвращает идентификатор ответа по которому можно получить результат обработки.
func (s *server) QueueProcess(ctx context.Context, request *count.QRequest) (*count.QResultId, error) {
	log.Println(fmt.Sprintf("QueueProcess: %s", request.GetWordsToCount()))

	// Генерация уникального идентификатора
	respCounter = respCounter + 1

	// Формируем сообщение QProcessMessage для отправки в очередь
	body := &count.QProcessMessage{WordsCount: request.WordsToCount, ResultId: respCounter}
	bytes, err := proto.Marshal(body)

	// Отправляем сообщение в очередь
	rmqPublish(qReq.Name, bytes)
	failOnError(err, "Failed to publish a message")

	// Возвращаем QResultId, содержащий идентификатор, по которому можно будет получить результат
	return &count.QResultId{ResultId: respCounter}, nil
}

// gRPC-функция, возвращающая результат обработки по идентификатору ответа.
// Если результат не доступен (еще не подсчитан, или в принципе отсутствует), то hasResult=false
// Если результат доступен, то hasResult=true и wordsCount содержит число слов в ранее переданной строке
func (s *server) QueueGetResult(ctx context.Context, request *count.QResultId) (*count.QResult, error) {
	log.Println(fmt.Sprintf("QueueGetResult: %d", request.GetResultId()))

	// По-умолчанию считаем, что результата еще нет
	res := count.QResult{}

	// Если мы уже получили результат обработки и сохранили его в локальном словаре
	if val, ok := responses[request.ResultId]; ok {
		// Наполняем структуру результатом обработки
		res.WordsCount = val.WordsCount
		log.Println(fmt.Sprintf("Return result: %d", val.WordsCount))
		return &res, nil
	}

	// Если ответ не готов, то возвращаем пустую структуру со статусом "Идентификатор не найден"
	return &res, status.Error(codes.NotFound, "ResultId was not found")
}

// Подключение к очереди сообщений RabbitMQ и создание канала для взаимодействия с ней.
// Для подключение передается строка подключения формата "amqp://guest:guest@localhost:5672/".
func rmqConnect(connString string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(connString)
	ch, err := conn.Channel()
	return conn, ch, err
}

// Создание очереди с заданным именем.
// Либо подтверждение наличия такой очереди если она уже создана.
// В работе используется глобальая переменная - канал ch.
func rmqQueueDeclare(queueName string) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	return q, err
}

// Отправка сообщения body в очередь с именем queueName.
// В работе используется глобальая переменная - канал ch.
func rmqPublish(queueName string, body []byte) error {
	err := ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/binary",
			Body:        body,
		})
	return err
}

// Получение потока сообщений из очереди с именем queueName.
// В работе используется глобальая переменная - канал ch.
func rmqConsume(queueName string) (<-chan amqp.Delivery, error) {
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	return msgs, err
}

func main() {
	var err error
	// Инициализация соединения с очередью
	// Инициализация глобальных переменных
	conn, ch, err = rmqConnect("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	defer ch.Close()

	// Формируем очереди на отправку запросов на обработку данных
	qReq, err = rmqQueueDeclare("Request")
	// и получение обратных сообщений
	qResp, err = rmqQueueDeclare("Response")
	failOnError(err, "Failed to declare a queue")

	// Формируем поток на получение сообщений об обработке данных
	msgs, err := rmqConsume(qResp.Name)
	failOnError(err, "Failed to register a consumer")

	// Инициализация хранилища для результатов ответов
	responses = make(map[int32]count.QResultMessage)
	// Идентификаторы результатов resultId - это целые числа
	respCounter = 0

	// Инициализация gRPC сервера
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	count.RegisterCountServer(srv, &server{})

	// Запускаем gRPC сервер в отдельной горутине
	go srv.Serve(lis)

	// Стартуем обработку сообщений из очереди
	forever := make(chan bool)

	// Горутина обработки входящих сообщений
	go func() {
		// Как только получаем входящее сообщение d
		for d := range msgs {
			// Логгируем
			log.Printf("Received a message: %s", d.Body)
			// Для получения resultId десериализуем сообщение
			resp := &count.QResultMessage{}
			proto.Unmarshal(d.Body, resp)
			// Сохраняем сообщение в словаре под индексом resultId
			responses[resp.ResultId] = *resp
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
