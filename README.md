
Test Upload

```curl -X POST http://localhost:8080/tracks/upload \
  -F "file=@sample.wav"
{"id":"4b758cb4-b62d-4237-8e40-484a9bdb028e","status":"uploaded","file_path":"storage\\tracks\\4b758cb4-b62d-4237-8e40-484a9bdb028e\\sample.wav"}
```
```
$ curl "http://localhost:8080/tracks?id=4b758cb4-b62d-4237-8e40-484a9bdb028e"
{"id":"4b758cb4-b62d-4237-8e40-484a9bdb028e","status":"uploaded"}```

