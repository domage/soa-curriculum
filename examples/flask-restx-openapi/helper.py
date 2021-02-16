# Данные хранятся в SQLite базе данных todo.db в таблице todos
# CREATE TABLE "todos" "todo_id" INTEGER PRIMARY KEY, "description" TEXT NOT NULL, "status" TEXT NOT NULL );

import sqlite3
from flask import jsonify
import app

DB_PATH = './todo.db'   # Update this path accordingly
NOTSTARTED = 'Not Started'

# Получить все элементы в таблице
def get_all_todos():
    try:
        conn = sqlite3.connect(DB_PATH)
        # Обеспечивает работу с названиями колонок в таблице
        conn.row_factory = sqlite3.Row
        c = conn.cursor()
        c.execute('select * from todos')
        # Получаем список строк в перечислимом формате
        rows = c.fetchall()
        return rows
    except Exception as e:
        print('Error: ', e)
        return None

# Получить отдельный элемент
def get_todo(todo_id):
    try:
        conn = sqlite3.connect(DB_PATH)
        # Обеспечивает работу с названиями колонок в таблице
        conn.row_factory = sqlite3.Row
        c = conn.cursor()
        c.execute("select * from todos where todo_id=?;" , [todo_id])
        r = c.fetchone()
        return r
    except Exception as e:
        print('Error: ', e)
    return None

# Добавить элемент в таблицу
def add_to_list(description, status):
    try:
        conn = sqlite3.connect(DB_PATH)
        c = conn.cursor()
        c.execute('insert into todos(description, status) values(?,?)', (description, status))
        conn.commit()
        result = get_todo(c.lastrowid)
        return result
    except Exception as e:
        print('Error: ', e)
        return None

# Обновить элемент с todo_id в таблице
def update_todo(todo_id, description, status):
    try:
        conn = sqlite3.connect(DB_PATH)
        c = conn.cursor()
        c.execute('update todos set description=?, status=? where todo_id=?', (description, status, todo_id))
        conn.commit()
        result = get_todo(todo_id)
        return result
    except Exception as e:
        print('Error: ', e)
        return None

# Удалить элемент из таблицы по индексу
def remove_todo(todo_id):
    try:
        conn = sqlite3.connect(DB_PATH)
        c = conn.cursor()
        c.execute('DELETE FROM todos WHERE todo_id=?', [todo_id])
        conn.commit()
        return jsonify( { 'result': True } )
    except Exception as e:
        print('Error: ', e)
        return None

