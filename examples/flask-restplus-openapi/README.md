# Upgraded ToDo REST-service using Python, Fask and Flask-RESTX
Это расширение учебного приложения-примера, демонстрирующее работы с расширением [Flask-RESTX](https://flask-restx.readthedocs.io/en/latest/index.html) для систематизации работы с REST-сервисами и автоматической генерации документации по API с использованием [OpenAPI](https://swagger.io/specification/).

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
- ``python -m pip install flask_restx``
- ``python -m flask run``

### Работа с сервисом
Доступа ToDo API организован из корня сервера и по-умолчанию доступен по адресу: [http://127.0.0.1:5000/](http://127.0.0.1:5000/)