#### convert mp4 to mp3 with ffmpeg
ffmpeg -i input.mp4 output.mp3
#### convert files to flac:
sox input.mp3 --rate 16k --bits 16 --channels 1 output.flac
