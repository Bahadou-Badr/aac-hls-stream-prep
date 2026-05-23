
# Audio Stream Preparation Service

A modern Go backend service that prepares uploaded audio files for adaptive streaming using AAC, HLS, and CMAF/fMP4 packaging.

Built as a simplified media ingestion pipeline inspired by modern streaming platforms.

## Features

- Audio upload API
- AAC transcoding with FFmpeg
- Adaptive bitrate ladder generation
  - 96 kbps AAC
  - 160 kbps AAC
- HLS packaging
- CMAF/fMP4 segmented streaming
- Clean Go architecture
- Service / Repository pattern
- Modular media processing pipeline

---

# Architecture

```text
Upload Audio
      ↓
Store Original File
      ↓
AAC Transcoding
      ↓
Bitrate Ladder Generation
      ↓
HLS + CMAF Packaging
      ↓
Streaming-Ready Output
```
## Tech Stack
- Go
- FFmpeg
- HLS
- CMAF / fMP4
- AAC Codec


## API Endpoints
Upload Audio
```
POST /tracks/upload
```

Example
```
curl -X POST http://localhost:8080/tracks/upload \
  -F "file=@sample.wav"
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
Generated Output
```
storage/
└── tracks/
    └── {track_id}/
        ├── original.wav
        ├── track_96k.m4a
        ├── track_160k.m4a
        │
        └── hls/
            ├── 96k/
            │   ├── init.mp4
            │   ├── playlist.m3u8
            │   ├── segment_000.m4s
            │   └── segment_001.m4s
            │
            └── 160k/
                ├── init.mp4
                ├── playlist.m3u8
                ├── segment_000.m4s
                └── segment_001.m4s
```

Streaming Technologies
This project uses:

- AAC for modern audio compression
- HLS for adaptive HTTP streaming
- CMAF/fMP4 fragmented media segments

Instead of legacy MPEG-TS segments, the pipeline generates modern fragmented MP4 (.m4s) outputs with initialization segments (init.mp4).

