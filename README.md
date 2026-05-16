
Test Upload

```curl -X POST http://localhost:8080/tracks/upload \
  -F "file=@sample.wav"
{"id":"4b758cb4-b62d-4237-8e40-484a9bdb028e","status":"uploaded","file_path":"storage\\tracks\\4b758cb4-b62d-4237-8e40-484a9bdb028e\\sample.wav"}
```

```
$ curl "http://localhost:8080/tracks?id=4b758cb4-b62d-4237-8e40-484a9bdb028e"
{"id":"4b758cb4-b62d-4237-8e40-484a9bdb028e","status":"uploaded"}
```

```
{
  "id": "e5e1f1c4-9815-401b-8a63-f76a66ad1139",
  "status": "ready",
  "file_path": "storage\\tracks\\e5e1f1c4-9815-401b-8a63-f76a66ad1139\\sample.wav",
  "variants": [
    {
      "bitrate": "96k",
      "path": "./storage/tracks/e5e1f1c4-9815-401b-8a63-f76a66ad1139/track_96k.m4a"
    },
    {
      "bitrate": "160k",
      "path": "./storage/tracks/e5e1f1c4-9815-401b-8a63-f76a66ad1139/track_160k.m4a"
    }
  ]
}
```