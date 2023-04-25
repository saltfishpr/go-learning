# go-whisper

go-whisper 使用[whisper.cpp](https://github.com/ggerganov/whisper.cpp)将语音转为文本.

### Requirement

- [go](https://go.dev/dl/)
- [ffmpeg](https://ffmpeg.org/download.html)

### Build

Linux:

```shell
LIB_FOLDER="$(pwd)/lib/$(uname -s)-$(uname -m)" && C_INCLUDE_PATH=$LIB_FOLDER LIBRARY_PATH=$LIB_FOLDER go build -o go-whisper
```

Windows:

TODO

### Usage

```shell
/path/to/go-whisper '<audio_file>'
```
