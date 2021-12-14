import requests

resp = requests.post("http://127.0.0.1:5000/logs", json={"ip": "123.123.123.123", "timestamp": "xxx", "url": "yyy"})
resp = requests.post("http://127.0.0.1:5000/logs", json={"ip": "xxx", "timestamp": "xxx", "url": "yyy"})
resp = requests.post("http://127.0.0.1:5000/logs", json={"ip": "1yyxyx", "timestamp": "xxx", "url": "yyy"})
resp = requests.post("http://127.0.0.1:5000/logs", json={"ip": "1111111111111", "timestamp": "xxx", "url": "yyy"})
resp = requests.post("http://127.0.0.1:5000/logs", json={"ip": "1111111111111", "timestamp": "xxx", "url": "yyy"})

resp = requests.get("http://127.0.0.1:5000/visitors")
print(resp.text)