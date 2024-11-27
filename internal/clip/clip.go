package clip

import (
	"github.com/zyedidia/clipper"
)

var clip clipper.Clipboard

func init() {
	var err error
	clip, err = clipper.GetClipboard(clipper.Clipboards...)
	if err != nil {
		return
	}
}

func Read() string {
	data, err := clip.ReadAll(clipper.RegClipboard)
	if err != nil {
		return ""
	}

	return string(data)
}

func Write(str string) {
	_ = clip.WriteAll(clipper.RegClipboard, []byte(str))
}
