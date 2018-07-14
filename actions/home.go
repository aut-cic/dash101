package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/zencoder/go-dash/mpd"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

// MPDHandler returns Media Presentation Description
func MPDHandler(c buffalo.Context) error {
	// https://en.wikipedia.org/wiki/ISO_8601#Durations
	m := mpd.NewMPD(mpd.DASH_PROFILE_LIVE, "PT1M6.2S", "PT25.8S")

	// Audio
	/*
		audioAS, err := m.AddNewAdaptationSetAudio(mpd.DASH_MIME_TYPE_AUDIO_MP4, true, 1, "und")
		if err != nil {
			return c.Error(http.StatusInternalServerError, err)
		}

		audioRep, err := audioAS.AddNewRepresentationAudio(44100, 128558, "mp4a.40.5", "800k/audio-und")
		if err != nil {
			return c.Error(http.StatusInternalServerError, err)
		}
		audioRep.SetNewBaseURL("/cdn/800k/output-audio-und.mp4")
		audioRep.AddNewSegmentBase("629-756", "0-628")
	*/

	// Video
	videoAS, err := m.AddNewAdaptationSetVideo(mpd.DASH_MIME_TYPE_VIDEO_MP4, "progressive", true, 1)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	if _, err := videoAS.SetNewSegmentTemplate(107520, "/cdn/$RepresentationID$/init-stream0.m4s", "/cdn/$RepresentationID$/chunk-stream0-$Number%05d$.m4s", 1, 12800); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	// bandwidth - in Bits/s (i.e. 1518664).
	// codecs - codec string for Audio Only (in RFC6381, https://tools.ietf.org/html/rfc6381) (i.e. avc1.4d401f).
	// id - ID for this representation, will get used as $RepresentationID$ in template strings.
	// frameRate - video frame rate (as a fraction) (i.e. 30000/1001).
	// width - width of the video (i.e. 1280).
	// height - height of the video (i.e 720).
	if _, err := videoAS.AddNewRepresentationVideo(118119494, "avc1.64001f", "800k", "25/1", 1280, 720); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	if _, err := videoAS.AddNewRepresentationVideo(18260565, "avc1.64000d", "400k", "25/1", 320, 240); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	/*
		videoRep2, err := videoAS.AddNewRepresentationVideo(1633516, "avc1.4d401f", "1200k/video-1", "30000/1001", 960, 540)
		if err != nil {
			return c.Error(http.StatusInternalServerError, err)
		}
		videoRep2.SetNewBaseURL("/cdn/1200k/output-video-1.mp4")
	*/

	// Subtitle
	/*
		subtitleAS, err := m.AddNewAdaptationSetSubtitle(mpd.DASH_MIME_TYPE_SUBTITLE_VTT, "en")
		if err != nil {
			return c.Error(http.StatusInternalServerError, err)
		}

		subtitleRep, err := subtitleAS.AddNewRepresentationSubtitle(256, "captions_en")
		if err != nil {
			return c.Error(http.StatusInternalServerError, err)
		}
		subtitleRep.SetNewBaseURL("http://example.com/content/sintel/subtitles/subtitles_en.vtt")
	*/

	return c.Render(200, r.XML(m))
}
