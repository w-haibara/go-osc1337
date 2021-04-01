package osc1337

import (
	"encoding/base64"
	"fmt"
	"io"
)

type Encoder struct {
	W                   io.Writer
	Size                int
	Inline              bool
	Name                []byte
	Width               string
	Height              string
	PreserveAspectRatio bool
	Align               string
	Type                string
}

func NewEncoder() Encoder {
	return Encoder{
		Inline:              false,
		Width:               "auto",
		Height:              "auto",
		PreserveAspectRatio: true,
		Align:               "left",
	}
}

func (e Encoder) Encode(img []byte) {
	if e.W == nil || img == nil {
		return
	}

	inline := 0
	if e.Inline {
		inline = 1
	}

	preserveAspectRatio := 0
	if e.PreserveAspectRatio {
		preserveAspectRatio = 1
	}

	fmt.Fprintf(e.W,
		"\033]1337;File=name=%s;inline=%d;width=%s;height=%s;preserveAspectRatio=%d;align=%s;type=%s:%s\a\n",
		base64.StdEncoding.EncodeToString(e.Name),
		inline, e.Width, e.Height, preserveAspectRatio, e.Align, e.Type,
		base64.StdEncoding.EncodeToString(img))
}
