from websocket import create_connection
a = -1900000000000000000

def ta(a):
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
    # -1900000000000000000 ~ -3600000000000000000
    #ws.send("-2000000000000000000") # 一个简单的溢出
    # a = -1900000000000000000 + 1844674407370955161
    # a = -1900000000000000000 - 1844674407370955161 * 201
    ws.send(str(a))
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

ta(a)