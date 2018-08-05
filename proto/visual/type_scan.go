package fb_visual

import (
	"fmt"
)

func (receiver *Type) Scan(value interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	const minUint32 = 0
	const maxUint32 = 4294967295

	switch casted := value.(type) {
	case []byte:
		switch {
		case "FB_VISUAL_MONO01"             == string(casted): // Starting from Go 1.5 the compiler optimizes this type of string versus []byte comparison.
			receiver.value = CodeMono01
		case "FB_VISUAL_MONO10"             == string(casted): // Starting from Go 1.5 the compiler optimizes this type of string versus []byte comparison.
			receiver.value = CodeMono10
		case "FB_VISUAL_TRUECOLOR"          == string(casted): // Starting from Go 1.5 the compiler optimizes this type of string versus []byte comparison.
			receiver.value = CodeTrueColor
		case "FB_VISUAL_PSEUDOCOLOR"        == string(casted): // Starting from Go 1.5 the compiler optimizes this type of string versus []byte comparison.
			receiver.value = CodePseudoColor
		case "FB_VISUAL_DIRECTCOLOR"        == string(casted): // Starting from Go 1.5 the compiler optimizes this type of string versus []byte comparison.
			receiver.value = CodeDirectColor
		case "FB_VISUAL_STATIC_PSEUDOCOLOR" == string(casted): // Starting from Go 1.5 the compiler optimizes this type of string versus []byte comparison.
			receiver.value = CodeStaticPseudoColor
		case "FB_VISUAL_FOURCC"             == string(casted): // Starting from Go 1.5 the compiler optimizes this type of string versus []byte comparison.
			receiver.value = CodeFourCC
		default:
			return fmt.Errorf("fbdev: scanning from type %T is supported, but value unrecognized: %q", value, casted)
		}

	case int8:
		if casted < minUint32 {
			return fmt.Errorf("bdev: scanning from type %T is supported, but value is not in range: %d", value, casted)
		}
		receiver.value = uint32(casted)
	case int16:
		if casted < minUint32 {
			return fmt.Errorf("bdev: scanning from type %T is supported, but value is not in range: %d", value, casted)
		}
		receiver.value = uint32(casted)
	case int32:
		if casted < minUint32 {
			return fmt.Errorf("bdev: scanning from type %T is supported, but value is not in range: %d", value, casted)
		}
		receiver.value = uint32(casted)
	case int64:
		if casted < minUint32 || maxUint32 < casted {
			return fmt.Errorf("bdev: scanning from type %T is supported, but value is not in range: %d", value, casted)
		}
		receiver.value = uint32(casted)

	case string:
		switch casted {
		case "FB_VISUAL_MONO01":
			receiver.value = CodeMono01
		case "FB_VISUAL_MONO10":
			receiver.value = CodeMono10
		case "FB_VISUAL_TRUECOLOR":
			receiver.value = CodeTrueColor
		case "FB_VISUAL_PSEUDOCOLOR":
			receiver.value = CodePseudoColor
		case "FB_VISUAL_DIRECTCOLOR":
			receiver.value = CodeDirectColor
		case "FB_VISUAL_STATIC_PSEUDOCOLOR":
			receiver.value = CodeStaticPseudoColor
		case "FB_VISUAL_FOURCC":
			receiver.value = CodeFourCC
		default:
			return fmt.Errorf("fbdev: scanning from type %T is supported, but value unrecognized: %q", value, casted)
		}

	case uint8:
		receiver.value = uint32(casted)
	case uint16:
		receiver.value = uint32(casted)
	case uint32:
		receiver.value = casted
	case uint64:
		if maxUint32 < casted {
			return fmt.Errorf("bdev: scanning from type %T is supported, but value is not in range: %d", value, casted)
		}
		receiver.value = uint32(casted)

	default:
		return fmt.Errorf("fbdev: scanning from type %T not supported.", value)
	}

	return nil
}
