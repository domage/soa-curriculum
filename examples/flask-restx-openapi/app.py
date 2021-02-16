# Подробнее узнать про flask_restx можно тут: https://flask-restx.readthedocs.io/en/latest/index.html

import helper
from flask import Flask, abort, request
from flask_restx import Api, Resource, fields

flask_app = Flask(__name__)
api = Api(app = flask_app, version='1.0', title='ToDo API',
    description='A simple ToDo API')

name_space = api.namespace('todos', description='ToDo Operations')

# Модель задачи. Используется для сериализации.
todo = api.model('Todo', {
    'todo_id': fields.Integer(required=False, readonly=True, description='The task id'),
    'description': fields.String(required=True, description='The task details'),
    'status': fields.String(required=True, description='The task status')
})

# Модель задачи, но вместо id указываем URL. Эта модель возвращается клиенту.
todoUrl = api.model('Todo URL', {
    'todo_url': fields.Url('todo_ep', readonly=True, absolute=True, description='The task URL'),
    'description': fields.String(required=True, description='The task details'),
    'status': fields.String(required=True, description='The task status')
})

# Получаем, редактируем и удаляем отдельную запись по индексу todo_id
# endpoint='todo_ep' позволяет нам в дальнейшем формировать URL для модели ToDo URL
@name_space.route('/todoapp/api/v1.0/todos/<int:todo_id>', endpoint='todo_ep')
@name_space.response(404, 'Todo not found')
@name_space.param('todo_id', 'The task identifier')
class ToDo(Resource):
    @name_space.doc('get_todo')
    # Возвращаемый результат сериализуем в JSON моделью todoUrl
    @name_space.marshal_with(todoUrl)
    def get(self, todo_id):
        # Получаем запись из базы данных
        result = helper.get_todo(todo_id)
        if result is None:
            abort(404)
        return result

    @name_space.doc('update_todo')
    # Проверяем тело входного запроса на соответствие с моделью Todo
    @name_space.expect(todo, validate=True)
    # Результат возвращаем моделью TodoURL
    @name_space.marshal_with(todoUrl)
    def put(self, todo_id):
        # Получаем запись из базы данных
        response = helper.get_todo(todo_id)
        # Если не найдено - ошибка 404
        if response is None:
            abort(404)
        response = helper.update_todo(todo_id, request.get_json()['description'], request.get_json()['status'])
        # Если не удачно - возвращаем ошибку 400
        if response is None:
            abort(400)
        # Возвращаем полное описание добавленного элемента
        return response
    
    @name_space.doc('delete_todo')
    @name_space.response(204, 'Todo deleted')
    def delete(self, todo_id):
        # Получаем запись из базы данных
        response = helper.get_todo(todo_id)
        # Если не найдено - ошибка 404
        if response is None:
            abort(404)
        helper.remove_todo(todo_id)
        return '', 204

# Получаем список всех записей и добавляем новую запись
@name_space.route('/todoapp/api/v1.0/todos')
class ToDoList(Resource):
    @name_space.doc('list_todos')
    @name_space.marshal_list_with(todoUrl)
    def get(self):
        return helper.get_all_todos()

    @name_space.doc('create_todo')
    @name_space.expect(todo, validate=True)
    @name_space.marshal_with(todoUrl, code=201)
    def post(self):
        # Добавляем элемент в базу данных
        response = helper.add_to_list(request.get_json()['description'], request.get_json()['status'])
        # Если не удачно - возвращаем ошибку 400
        if response is None:
            abort(400)
        # Возвращаем полное описание добавленного элемента
        return response, 201