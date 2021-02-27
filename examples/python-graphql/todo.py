from typing import List
import typing
import strawberry
from strawberry.scalars import ID

@strawberry.type
class Todo:
    id: ID
    name: str
    done: bool

@strawberry.input
class TodoInput:
    name: str
    done: typing.Optional[bool]

todoID = 2

todoDict = {
    0: Todo(id=0, name="Todo #0", done=False),
    1: Todo(id=1, name="Todo #1", done=False),
    2: Todo(id=2, name="Todo #3", done=True)
} 

@strawberry.type
class Query:
    @strawberry.field
    def todos(self, info, done: bool = None) -> List[Todo]:
        if done is not None:
            return filter(lambda todo: todo.done == done, todoDict.values())
        else:
            return todoDict.values()

@strawberry.type
class Mutation:
    @strawberry.mutation
    def createTodo(self, info, name: str, done: bool = False ) -> Todo:
        global todoID
        todoID += 1
        todo = Todo(
            id = todoID,
            name=name,
            done=done
        )
        todoDict[todoID]=todo
        return todo

    @strawberry.mutation
    def updateTodo(self, info, id: int, todo: TodoInput) -> Todo:
        old_todo = todoDict[id]
        old_todo.name = todo.name
        old_todo.done = todo.done
        return old_todo
    
    @strawberry.mutation
    def deleteTodo(self, info, id: int) -> Todo:
        old_todo = todoDict[id]
        del todoDict[id]
        return old_todo

schema = strawberry.Schema(query=Query,mutation=Mutation)
print (schema)