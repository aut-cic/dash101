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
	m := mpd.NewMPD(mpd.DASH_PROFILE_ONDEMAND, "PT30S", "PT1.97S")

	// Audio
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

	// Video
	videoAS, err := m.AddNewAdaptationSetVideo(mpd.DASH_MIME_TYPE_VIDEO_MP4, "progressive", true, 1)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	videoRep1, err := videoAS.AddNewRepresentationVideo(1100690, "avc1.4d401e", "800k/video-1", "30000/1001", 640, 360)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	videoRep1.SetNewBaseURL("/cdn/800k/output-video-1.mp4")
	videoRep1.AddNewSegmentBase("686-813", "0-685")

	videoRep2, err := videoAS.AddNewRepresentationVideo(1633516, "avc1.4d401f", "1200k/video-1", "30000/1001", 960, 540)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	videoRep2.SetNewBaseURL("/cdn/1200k/output-video-1.mp4")
	videoRep2.AddNewSegmentBase("686-813", "0-685")

	// Subtitle
	subtitleAS, err := m.AddNewAdaptationSetSubtitle(mpd.DASH_MIME_TYPE_SUBTITLE_VTT, "en")
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	subtitleRep, err := subtitleAS.AddNewRepresentationSubtitle(256, "captions_en")
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	subtitleRep.SetNewBaseURL("http://example.com/content/sintel/subtitles/subtitles_en.vtt")

	return c.Render(200, r.XML(m))
}
