# LiveTv

[![Buffalo](https://img.shields.io/badge/powered%20by-buffalo-blue.svg?style=flat-square)](http://gobuffalo.io)
[![Travis branch](https://img.shields.io/travis/AUTProjects/LiveTv/master.svg?style=flat-square)](https://travis-ci.org/AUTProjects/LiveTv)

## Introduction
IRIB TV Live on Your browser based on dash

## Demo
Download [sample video](https://www.sample-videos.com/video/mp4/720/big_buck_bunny_720p_10mb.mp4) and store it,
then creates it's ts segments and store them into `public/cdn` and serve them.

## Let's Learn ffmpeg

### Installation

```sh
apt install ffmpeg
```

### Live with webcam

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
