# На основе материалов туториала https://blog.miguelgrinberg.com/post/designing-a-restful-api-with-python-and-flask
# Исходный код туториала: https://github.com/geonaut/flask-todo-rest-api
# Настройка VS Code: https://code.visualstudio.com/docs/python/tutorial-flask
# TODO: доделать авторизацию https://blog.miguelgrinberg.com/post/restful-authentication-with-flask
# Авторизация JWT: https://realpython.com/token-based-authentication-with-flask/

import helper
from flask import Flask, abort, request
from datetime import datetime
import re

app = Flask(__name__)

# Получаем отдельную запись по индексу todo_id
@app.route('/todoapp/api/v1.0/todos/<int:todo_id>', methods=['GET'])
def get_todo(todo_id):
    # Получаем запись из базы данных
    response = helper.get_todo(todo_id)

    # Если не найдено - ошибка 404
    if response is None:
        abort(404)
    
    return response

# Получаем список всех элементов коллекции
@app.route('/todoapp/api/v1.0/todos', methods=['GET'])
def get_all_todos():
    return helper.get_all_todos()

# Добавить элемент в коллекцию. В теле запроса должен быть передан JSON с полем 'description'
@app.route('/todoapp/api/v1.0/todos', methods=['POST'])
def add_todo():
    
    # Если в параметрах запроса нет тела, либо нет поля 'description' - отбой 
    if not request.json or not 'description' in request.json:
        abort(400)
    
    # Получаем поле из запроса
    description = request.get_json()['description']

    # Добавляем элемент в базу данных
    response = helper.add_to_list(description)

    # Если не удачно - возвращаем ошибку 400
    if response is None:
        abort(400)

    # Возвращаем полное описание добавленного элемента
    return response

# Добавить элемент в коллекцию. В теле запроса должен быть передан JSON с полем 'description'
@app.route('/todoapp/api/v1.0/todos/<int:todo_id>', methods=['PUT'])
def update_todo(todo_id):
    
    # Получаем запись из базы данных
    response = helper.get_todo(todo_id)

    # Если не найдено - ошибка 404
    if response is None:
        abort(404)

    # Если в параметрах запроса нет тела, либо нет поля 'description' - отбой 
    if not request.json:
        abort(400)
    if not 'description' in request.json:
        abort(400)
    if not 'status' in request.json:
        abort(400)
    
    # Добавляем элемент в базу данных
    response = helper.update_todo(todo_id, request.get_json()['description'], request.get_json()['status'])

    # Если не удачно - возвращаем ошибку 500
    if response is None:
        abort(400)

    # Возвращаем полное описание добавленного элемента
    return response

# Удалить элемент из коллекции
@app.route('/todoapp/api/v1.0/todos/<int:todo_id>', methods = ['DELETE'])
def delete_task(todo_id):

    # Получаем запись из базы данных
    response = helper.get_todo(todo_id)

    # Если не найдено - ошибка 404
    if response is None:
        abort(404)

    response = helper.remove_todo(todo_id)

    return response