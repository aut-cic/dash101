# LiveTv
IRIB TV Live on Your browser based on dash

## Let's Learn ffmpeg

1. List input stream devices

```sh
ffmpeg -f avfoundation -list_devices true -i ""
```

2. Stream FaceTime HD Camera into `out.mpg`

```sh
ffmpeg -f avfoundation -r 30 -s 1280x720 -i "0" out.mpg
```

3. Multi quality dash from static video

```sh
ffmpeg -i SampleVideo_1280x720_5mb.mp4 -map 0:v:0 -map 0:a:0 -map 0:v:0 -map 0:a:0 -b:v:0 250k -filter:v:0 "scale=-2:240" -profile:v:0 baseline -filter:v:1 "scale=-2:720" -profile:v:1 main -use_timeline 1 -use_template 1 -window_size 5 -adaptation_sets "id=0,streams=v id=1,streams=a" -f dash hello.mpd
```
