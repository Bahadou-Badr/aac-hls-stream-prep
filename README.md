
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
### Project Architecture
```
main.go        → bootstrap/infrastructure
router.go      → route registration
handler.go     → HTTP layer
service.go     → business logic
worker/        → async processing
transcoder/    → FFmpeg
hls/           → packaging
storage/       → filesystem abstraction
```

### Workflow 
```text
Upload Audio
      ↓
Store Original File
      ↓ 
worker queue
      ↓
FFmpeg pipeline
      ↓
AAC Transcoding
      ↓
Bitrate Ladder Generation
      ↓
HLS/CMAF packaging
      ↓
HTTP stream delivery
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

Modern platforms are moving from legacy MP3 and MPEG-TS
toward AAC + HLS + CMAF/fMP4 because of:

- better compression efficiency
- lower latency
- improved compatibility
- adaptive bitrate streaming
- CDN optimization
---
This project now demonstrates:
### Backend Engineering
- clean architecture
- service layer
- worker pools
- async processing
- REST APIs
### Media Engineering
- AAC encoding
- bitrate ladders
- HLS
- CMAF/fMP4

### Systems Thinking
- ingestion pipeline
- stream preparation
- scalable processing flow