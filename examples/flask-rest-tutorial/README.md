# ToDo REST-service using Python and Fask 
Это учебное приложение-пример, демонстрирующее основы создания REST веб-сервисов на базе языка Python и веб-фреймворка Flask. Веб-сервис обеспечивает работу со списком дел "ToDo". При помощи REST API обеспечивается возможность взаимодействия со списком дел путем обмена информацией в формате JSON, включая:
- по запросу ``GET`` к списку ToDo: получение списка ToDo;
- по запросу ``GET`` отдельного ToDo: получение отдельного ToDo по его индексу;
- по запросу ``POST`` к списку ToDo: добавление нового ToDo;
- по запросу ``PUT`` отдельного ToDo: обновление записи отдельного ToDo;
- по запросу ``DELETE`` отдельного ToDo: удаление отдельного ToDo.

Хранение данных производится в файле ``todo.db``, представляющем собой базу данных SQLite.

## Подготовка и настройка
### Подготовка
Перед запуском системы, необходимо:
- Установить [python3](https://www.python.org/download/releases/3.0/)
- При использовании Visual Studio Code, можно подготовить среду для разработки посредством материалов туториола: https://code.visualstudio.com/docs/python/tutorial-flask

### Запуск приложения
В корневой директории выполнить 
- ``pip install virtualenv``
- ``python -m venv flask-todo``
- ``virtualenv flask-todo``
- ``python -m pip install flask``
- ``python -m flask run``

### Работа с сервисом
Для отправки запросов можно использовать утилиту ``curl`` либо [postman](https://www.postman.com/downloads/). Пример запросов (при условии, что сервер развернут по адресу ``http://127.0.0.1:5000``):
- получение списка ToDo: ``curl --location --request GET 'http://127.0.0.1:5000/todoapp/api/v1.0/todos'``
- получение отдельного элемента ToDo: ``curl --location --request GET 'http://127.0.0.1:5000/todoapp/api/v1.0/todos'``
- получение отдельного элемента ToDo: ``curl --location --request GET 'http://127.0.0.1:5000/todoapp/api/v1.0/todos'``
- добавление нового элемента ToDo: ``curl --location --request POST 'http://127.0.0.1:5000/todoapp/api/v1.0/todos' --header 'Content-Type: application/json' --data-raw '{"description": "To test POST"}'``
- обновление записи отдельного элемента ToDo: ``curl --location --request PUT 'http://127.0.0.1:5000/todoapp/api/v1.0/todos/3' --header 'Content-Type: application/json' --data-raw '{ "description": "To test POST", "status": "Started","uri": "http://127.0.0.1:5000/todoapp/api/v1.0/todos/3"}'``
- удаление отдельного ToDo: ``curl --location --request DELETE 'http://127.0.0.1:5000/todoapp/api/v1.0/todos/3'``

## Ссылки и благодарности
Приложение подготовлено на основе материалов туториала [Designing a RESTful API with Python and Flask (Miguel Grinberg)](https://blog.miguelgrinberg.com/post/designing-a-restful-api-with-python-and-flask). Исходный код материалов туториала доступен в репозитории: https://github.com/geonaut/flask-todo-rest-api