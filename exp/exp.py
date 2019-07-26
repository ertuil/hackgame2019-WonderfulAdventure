from websocket import create_connection
ws = create_connection("ws://localhost:8080/ws")

result = ws.recv()
print(result)

ws.send("0")

result = ws.recv()
print(result)
result = ws.recv()
print(result)

ws.send("0")
result = ws.recv()
print(result)

result = ws.recv()
print(result)

ws.send("0")
result = ws.recv()
print(result)

ws.send("-2000000000000000000") # 一个简单的溢出
result = ws.recv()
print(result)

result = ws.recv()
print(result)

ws.send("2")
result = ws.recv()
print(result)

ws.send("0")
result = ws.recv()
print(result)

result = ws.recv()
print(result)

result = ws.recv()
print(result)

ws.send("2")
result = ws.recv()
print(result)

result = ws.recv()
print(result)

ws.close()