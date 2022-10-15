from json import loads

with open ('todo_DUMP_10-15.json') as file:
    data = loads(file.read())

print(data)