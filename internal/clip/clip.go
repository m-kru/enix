package clip

import (
	"github.com/zyedidia/clipper"
)

var clip clipper.Clipboard

var clipboards = []clipper.Clipboard{
	&clipper.Wayland{},
	&clipper.Xclip{},
	&clipper.Xsel{},
	&clipper.Wsl{},
	&clipper.Termux{},
	&clipper.Internal{},
}

func init() {
	var err error
	clip, err = clipper.GetClipboard(clipboards...)
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
