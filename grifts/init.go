package grifts

import (
	"github.com/AUTProjects/livetv/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
