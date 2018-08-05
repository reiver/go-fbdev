package fb_visual

import (
	"fmt"
)

type Type struct {
	value uint32
}

func (receiver Type) Code() uint32 {
	return receiver.value
}

func (receiver Type) String() string {
	switch receiver.value {
	case CodeMono01:
		return "FB_VISUAL_MONO01"
	case CodeMono10:
		return "FB_VISUAL_MONO10"
	case CodeTrueColor:
		return "FB_VISUAL_TRUECOLOR"
	case CodePseudoColor:
		return "FB_VISUAL_PSEUDOCOLOR"
	case CodeDirectColor:
		return "FB_VISUAL_DIRECTCOLOR"
	case CodeStaticPseudoColor:
		return "FB_VISUAL_STATIC_PSEUDOCOLOR"
	case CodeFourCC:
		return "FB_VISUAL_FOURCC"
	default:
		return fmt.Sprintf("FB_VISUAL_UNKNOWN(%d)", receiver.value)
	}
}
