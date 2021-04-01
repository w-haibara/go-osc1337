package osc1337

import (
	"encoding/base64"
	"fmt"
	"io"
)

type Encoder struct {
	W                   io.Writer
	Name                []byte
	Size                int
	Width               string
	Height              string
	PreserveAspectRatio bool
	Inline              bool
	Align               string
	Type                string
}

func NewEncoder() Encoder {
	return Encoder{
		Width:               "auto",
		Height:              "auto",
		PreserveAspectRatio: true,
		Align:               "left",
		Inline:              false,
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
		"\x1b]1337;File=name=%s;width=%s;height=%s;preserveAspectRatio=%d;inline=%d:align=%s;type=%s;:%s",
		base64.StdEncoding.EncodeToString(e.Name),
		e.Width, e.Height, preserveAspectRatio, inline, e.Align, e.Type,
		base64.StdEncoding.EncodeToString(img))
}
