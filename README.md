# websites-status-checker

POST Request -
```
curl "http://localhost:4000/post" -d '{"websites":["www.google.com","www.github.com","www.medium.com","www.facebook.com","www.udemy.com","www.oreilly.com","www.gasycguy.com","www.fakewebsite1.com"]}'
```

GET Request -
```
curl "http://localhost:4000/get"

curl "http://localhost:4000/get?name=www.facebook.com&name=www.fakewebsite1.com&name=www.gmail.com"
```
--------------------------------------------------------------------------------------------------------------------------------------------------------

Http Endpoints Request -
```
POST http://localhost:4000/post
Content-Type: application/json
{"websites":["www.google.com","www.github.com","www.medium.com","www.facebook.com","www.udemy.com","www.oreilly.com","www.gasycguy.com","www.fakewebsite1.com"]}


GET http://localhost:4000/get
Accept: application/json


GET http://localhost:4000/get?name=www.facebook.com&name=www.fakewebsite1.com&name=www.gmail.com
Accept: application/json
```
