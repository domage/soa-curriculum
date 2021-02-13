# Клиент, который:
# 1. Устанавливает соединение и gRPC канал с сервером
# 2. Для каждой строки в массиве Sentences (асинхронно):
#   2.1. Вызывает gRPC метод QueueProcess, передавая строку для анализа
#   2.2. В результате вызова получает идентификатор, по которому будет получен ответ
#   2.3. Пока не получен валидный ответ, раз в секунду опрашивает сервер 
#        о полученном результате методом QueueGetResult
#   2.4. Как только получен результат - выводит его на консоль
#
# Для работы с gRPC из Python необходимо выполнить: 
# pip install grpcio-tools
#
# При изменении протокола обмена данными (файла ../pkg/proto/count/count.proto),
# для сборки сериализатора/десериализатора (count_pb2.py) и заглушек gRPC (count_pb2_grpc.py) 
# необходимо из текущей папки выполнить команду:
# python -m grpc_tools.protoc -I..\pkg\proto\count \
#        --python_out=. --grpc_python_out=. ..\pkg\proto\count\count.proto

# Подключаем сгенерированные gRPC-библиотеки
import count_pb2_grpc
import count_pb2

import grpc
import time
import concurrent.futures

# Набор строк, которые мы хотим обработать
Sentences = ["This is a test",
             "Suppose we had to create a large number of threads",
             "Declaring a queue",
             "That's it for our publisher.",
             ".NET is a free and open-source"]


def submitRequest(sentence, stub):
    # Формируем gRPC запрос, передаем в нем сообщение для обработки
    req = count_pb2.QRequest(wordsToCount=sentence)

    # Вызываем удаленную функцию. 
    # Результат вызова - идентификатор ответа, 
    # по которому будем проверять доступность результатов обработки
    resp = stub.QueueProcess(req)

    print(resp.resultId)

    hasResult = False
    count = 0
    
    # Пока не получим обработанный результат, раз в секунду
    # передаем серверу идентификатор ответа при вызове QueueGetResult 
    # и проверяем, обработан он, или нет
    while hasResult == False:
        time.sleep(1)
        result = stub.QueueGetResult(resp)
        hasResult = result.hasResult
        count = result.wordsCount

    # Как только обработка завершилась - печатаем результат
    print(sentence + ": " + str(count))


# Подключаемся к gRPC серверу
channel = grpc.insecure_channel('localhost:9000')

# Создаем клиентскую заглушку
stub = count_pb2_grpc.CountStub(channel)

# Создаем 5 потоков, в каждом из которых отправляем запрос на подсчет слов в предложении
# из списка Sentences, после чего начинаем опрос сервера раз в секунду на получение
# результатов обработки
with concurrent.futures.ThreadPoolExecutor(max_workers=5) as executor:
   future_to_count = {executor.submit(
       submitRequest, sentence, stub): sentence for sentence in Sentences}
