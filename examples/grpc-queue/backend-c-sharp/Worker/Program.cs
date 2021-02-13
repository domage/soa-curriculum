// Обработчик, который 
//  1. Получает из очереди RabbitMQ сообщения в формате protobuf
//  2. Выполняет длительную обработку сообщений
//  3. Отправляет результаты обработки в очередь результатов в формате protobuf
//  4. Отправляет в канал RabbitMQ подтверждение об успешной обработке сообщения
//   
// Для работы с protobuf и RabbitMQ необходимо:
//  1. Установить protoc
//  2. Выполнить:
//      dotnet add package RabbitMQ.Client
//      dotnet add package Grpc
//      dotnet add package Grpc.Core
//      dotnet add package Grpc.Tools
//      dotnet add package Google.Protobuf
//      dotnet restore
//  3. Для сборки сериализатора/десериализатора (Count.cs) необходимо из текущей папки выполнить команду
//      protoc -I ../../pkg/proto/count --csharp_out=. ../../pkg/proto/count/count.proto
//
// Запуск:
// dotnet run
//
// Дополнительная информация:
// - Использование protobuf и gRPC в C#: 
//      https://medium.com/@nikhilajayk/creating-your-first-grpc-net-core-client-and-server-app-using-visual-studio-or-visual-studio-code-293a6a5a5f7
//      https://developers.google.com/protocol-buffers/docs/csharptutorial
// - Работа с RabbitMQ:
//      https://www.rabbitmq.com/tutorials/tutorial-one-dotnet.html
//      https://www.rabbitmq.com/tutorials/tutorial-two-dotnet.html

using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System;
using System.Text;
using System.Threading;
using System.IO;
using Google.Protobuf;
using Count;

namespace Worker
{
    class Program
    {
        static void Main(string[] args)
        {
            // Подключаемся к очереди RabbitMQ, развернутой на локальном хосте
            // Используется порт по-умолчанию 5672
            var factory = new ConnectionFactory() { HostName = "localhost" };
            using (var connection = factory.CreateConnection())
            using (var channel = connection.CreateModel())
            {
                // Формируем очереди на отправку запросов на обработку данных
                channel.QueueDeclare(queue: "Request",
                                     durable: false,
                                     exclusive: false,
                                     autoDelete: false,
                                     arguments: null);

                // и получение обратных сообщений
                channel.QueueDeclare(queue: "Response",
                                     durable: false,
                                     exclusive: false,
                                     autoDelete: false,
                                     arguments: null);

                // Устанавливаем, что обработчик получает необработанные сообщения по-одному
                // Это возможно только в случае, если при потреблении установлен флаг "autoAck: false" (см. ниже)
                // Иначе очередь не может понять, что обработчик занят
                channel.BasicQos(prefetchSize: 0, prefetchCount: 1, global: false);

                // Определяем канал-обработчик сообщений
                var consumer = new EventingBasicConsumer(channel);

                // Как только в канал поступает сообщение - начинаем его обработку
                consumer.Received += (model, ea) =>
                {
                    // Для отладки печатаем полученное сообщение в виде текста
                    // Так как мы работаем с protobuf, то будут появляться непечатаемые символы
                    var body = ea.Body.ToArray();
                    var message = Encoding.UTF8.GetString(body);
                    Console.WriteLine(" [x] Received {0}", message);

                    // Десериализация сообщения
                    Count.QProcessMessage iMessage = Count.QProcessMessage.Parser.ParseFrom(body);

                    // Подсчет слов в предложении (но это не точно)
                    int words = iMessage.WordsCount.Split(' ').Length;
                    // Симуляция бурной деятельности, процесс очень занят длительное время
                    Thread.Sleep(words * 1000);

                    Console.WriteLine(" [x] Done. Words count = " + words.ToString());
                    Console.WriteLine(" [x] Words = " + iMessage.WordsCount);

                    // Формируем сообщение с результатом обработки
                    Count.QResultMessage oMessage = new Count.QResultMessage { WordsCount = words, ResultId = iMessage.ResultId };
                    using (var stream = new MemoryStream(512))
                    {
                        // Сохраняем его
                        oMessage.WriteTo(stream);

                        Console.WriteLine(" [x] Message prepared: {0}", Encoding.UTF8.GetString(stream.ToArray()));

                        // Оптправляем сообщение в канал результатов
                        channel.BasicPublish(exchange: "",
                            routingKey: "Response",
                            basicProperties: null,
                            body: stream.ToArray());
                        
                        // Отладочная информация
                        Console.WriteLine(" [x] Sent {0}", Encoding.UTF8.GetString(stream.ToArray()));

                    }

                    // По завершению обработки сообщения, высылаем в очередь подтверждение, что обработка завершена
                    // После этого мы сможем получить очередное сообщение, если в очереди еще что-то осталось
                    channel.BasicAck(deliveryTag: ea.DeliveryTag, multiple: false);

                };

                // Получить сообщения из очереди "Request"
                // Если "autoAck: true" то все оставшиеся сообщения будут розданы "поровну" каждому из подключенных на момент подключения обработчиков.
                // Если "autoAck: false" и "prefetchCount: 1" у соединения установлен в 1 (см. выше) то сообщение передается обработчику только после того, 
                // как он подтвердил обработку предыдущего сообщения. См. https://www.rabbitmq.com/tutorials/tutorial-two-dotnet.html
                channel.BasicConsume(queue: "Request",
                                     autoAck: false,
                                     consumer: consumer);

                Console.WriteLine(" Press [enter] to exit.");
                Console.ReadLine();
            }
        }
    }
}