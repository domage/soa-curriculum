# Overcomplicated and incorrect word counting application
Это учебное приложение-пример, демонстрирующий возможности по организации асинхронной обработки запросов с использованием очередей сообщений. 
Архитектура приложения:
- Клиент на языке Python, который по gRPC связывается с Frontend-сервером, передает ему запрос на обработку данных, получает идентификатор, по которому в дальнейшем сможет получить результат обработки.
- Фронтэнд-сервер на языке Go, получающий по gRPC запросы от клиента и перенаправляющий их в очередь RabbitMQ, а также получающий из очереди результаты обработки, и, по запросу, передающий их клиенту
- Сервера обработки данных на языке C#, получающие из очереди сообщений запросы на обработку данных и отправляющие результаты обработки обратно в очередь.

## Подготовка и настройка
### Подготовка
Перед запуском системы, необходимо:
- Установить [python3](https://www.python.org/download/releases/3.0/), [golang](https://golang.org/dl/) и [.NET SDK (бывший .NET Core)](https://dotnet.microsoft.com/download/dotnet/thank-you/sdk-5.0.102-windows-x64-installer).
- Установить [protoc](https://developers.google.com/protocol-buffers/docs/downloads).
- Развернуть на локальной системе экземпляр [RabbitMQ](https://www.rabbitmq.com/). Это можно сделать с использованием docker командой  
  ``docker run -d --hostname my-rabbit -p 5672:5672 --name some-rabbit rabbitmq:3``

### Фронтэнд-сервер
В директории ``/frontend-go`` выполнить 
- ``go run main.go``

### Клиент
В директории ``/client-python`` выполнить 
- ``pip install grpcio-tools``
- ``python client.py``

### Сервера обработки данных
В директории ``backend-c-sharp\Worker`` выполнить 
- ``dotnet add package RabbitMQ.Client``
- ``dotnet add package Grpc``
- ``dotnet add package Grpc.Core``
- ``dotnet add package Grpc.Tools``
- ``dotnet add package Google.Protobuf``
- ``dotnet restore``
- ``dotnet run``