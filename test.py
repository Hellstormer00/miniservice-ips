import requests

resp = requests.post("http://127.0.0.1:5000/logs", json={"ip": "123.123.123.123", "timestamp": "xxx", "url": "yyy"})
print(resp, resp.text)