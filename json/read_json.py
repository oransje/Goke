from json import loads

with open ('data.json') as file:
    data = loads(file.read())

print(data)