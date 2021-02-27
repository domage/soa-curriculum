import typing
import strawberry

@strawberry.type
class Book:
    title: str
    author: str

@strawberry.type
class Query:
    books: typing.List[Book]

def get_books():
    return [
        Book(
            title='The Great Gatsby',
            author='F. Scott Fitzgerald',
        ),
    ]

@strawberry.type
class Query:
    books: typing.List[Book] = strawberry.field(resolver=get_books)


@strawberry.type
class Mutation:
    @strawberry.mutation
    def update_book(self, email: str) -> bool:
        print(f'sending email to {email}')
        return True

schema = strawberry.Schema(query=Query)
print (schema)