package fb_visual

import (
	"testing"
)

func TestTypeScan(t *testing.T) {

	tests := []struct{
		Value    interface{}
		Expected uint32
	}{
		{
			Value:    []byte("FB_VISUAL_MONO01"),
			Expected: CodeMono01,
		},
		{
			Value:    []byte("FB_VISUAL_MONO10"),
			Expected: CodeMono10,
		},
		{
			Value:    []byte("FB_VISUAL_TRUECOLOR"),
			Expected: CodeTrueColor,
		},
		{
			Value:    []byte("FB_VISUAL_PSEUDOCOLOR"),
			Expected: CodePseudoColor,
		},
		{
			Value:    []byte("FB_VISUAL_DIRECTCOLOR"),
			Expected: CodeDirectColor,
		},
		{
			Value:    []byte("FB_VISUAL_STATIC_PSEUDOCOLOR"),
			Expected: CodeStaticPseudoColor,
		},
		{
			Value:    []byte("FB_VISUAL_FOURCC"),
			Expected: CodeFourCC,
		},









		{
			Value:    "FB_VISUAL_MONO01",
			Expected: CodeMono01,
		},
		{
			Value:    "FB_VISUAL_MONO10",
			Expected: CodeMono10,
		},
		{
			Value:    "FB_VISUAL_TRUECOLOR",
			Expected: CodeTrueColor,
		},
		{
			Value:    "FB_VISUAL_PSEUDOCOLOR",
			Expected: CodePseudoColor,
		},
		{
			Value:    "FB_VISUAL_DIRECTCOLOR",
			Expected: CodeDirectColor,
		},
		{
			Value:    "FB_VISUAL_STATIC_PSEUDOCOLOR",
			Expected: CodeStaticPseudoColor,
		},
		{
			Value:    "FB_VISUAL_FOURCC",
			Expected: CodeFourCC,
		},









		{
			Value:      int8(0),
			Expected: uint32(0),
		},
		{
			Value:      int8(1),
			Expected: uint32(1),
		},
		{
			Value:      int8(2),
			Expected: uint32(2),
		},
		{
			Value:      int8(3),
			Expected: uint32(3),
		},
		{
			Value:      int8(4),
			Expected: uint32(4),
		},
		{
			Value:      int8(6),
			Expected: uint32(6),
		},
		{
			Value:      int8(7),
			Expected: uint32(7),
		},
		{
			Value:      int8(8),
			Expected: uint32(8),
		},
		{
			Value:      int8(9),
			Expected: uint32(9),
		},
		{
			Value:      int8(10),
			Expected: uint32(10),
		},

		{
			Value:      int8(119),
			Expected: uint32(119),
		},
		{
			Value:      int8(120),
			Expected: uint32(120),
		},
		{
			Value:      int8(121),
			Expected: uint32(121),
		},
		{
			Value:      int8(122),
			Expected: uint32(122),
		},
		{
			Value:      int8(123),
			Expected: uint32(123),
		},
		{
			Value:      int8(124),
			Expected: uint32(124),
		},
		{
			Value:      int8(125),
			Expected: uint32(125),
		},
		{
			Value:      int8(126),
			Expected: uint32(126),
		},
		{
			Value:      int8(127),
			Expected: uint32(127),
		},









		{
			Value:     int16(0),
			Expected: uint32(0),
		},
		{
			Value:     int16(1),
			Expected: uint32(1),
		},
		{
			Value:     int16(2),
			Expected: uint32(2),
		},
		{
			Value:     int16(3),
			Expected: uint32(3),
		},
		{
			Value:     int16(4),
			Expected: uint32(4),
		},
		{
			Value:     int16(6),
			Expected: uint32(6),
		},
		{
			Value:     int16(7),
			Expected: uint32(7),
		},
		{
			Value:     int16(8),
			Expected: uint32(8),
		},
		{
			Value:     int16(9),
			Expected: uint32(9),
		},
		{
			Value:     int16(10),
			Expected: uint32(10),
		},

		{
			Value:     int16(119),
			Expected: uint32(119),
		},
		{
			Value:     int16(120),
			Expected: uint32(120),
		},
		{
			Value:     int16(121),
			Expected: uint32(121),
		},
		{
			Value:     int16(122),
			Expected: uint32(122),
		},
		{
			Value:     int16(123),
			Expected: uint32(123),
		},
		{
			Value:     int16(124),
			Expected: uint32(124),
		},
		{
			Value:     int16(125),
			Expected: uint32(125),
		},
		{
			Value:     int16(126),
			Expected: uint32(126),
		},
		{
			Value:     int16(127),
			Expected: uint32(127),
		},
		{
			Value:     int16(128),
			Expected: uint32(128),
		},
		{
			Value:     int16(129),
			Expected: uint32(129),
		},
		{
			Value:     int16(130),
			Expected: uint32(130),
		},
		{
			Value:     int16(131),
			Expected: uint32(131),
		},

		{
			Value:     int16(32759),
			Expected: uint32(32759),
		},
		{
			Value:     int16(32760),
			Expected: uint32(32760),
		},
		{
			Value:     int16(32761),
			Expected: uint32(32761),
		},
		{
			Value:     int16(32762),
			Expected: uint32(32762),
		},
		{
			Value:     int16(32763),
			Expected: uint32(32763),
		},
		{
			Value:     int16(32764),
			Expected: uint32(32764),
		},
		{
			Value:     int16(32765),
			Expected: uint32(32765),
		},
		{
			Value:     int16(32766),
			Expected: uint32(32766),
		},
		{
			Value:     int16(32767),
			Expected: uint32(32767),
		},









		{
			Value:     int32(0),
			Expected: uint32(0),
		},
		{
			Value:     int32(1),
			Expected: uint32(1),
		},
		{
			Value:     int32(2),
			Expected: uint32(2),
		},
		{
			Value:     int32(3),
			Expected: uint32(3),
		},
		{
			Value:     int32(4),
			Expected: uint32(4),
		},
		{
			Value:     int32(6),
			Expected: uint32(6),
		},
		{
			Value:     int32(7),
			Expected: uint32(7),
		},
		{
			Value:     int32(8),
			Expected: uint32(8),
		},
		{
			Value:     int32(9),
			Expected: uint32(9),
		},
		{
			Value:     int32(10),
			Expected: uint32(10),
		},

		{
			Value:     int32(119),
			Expected: uint32(119),
		},
		{
			Value:     int32(120),
			Expected: uint32(120),
		},
		{
			Value:     int32(121),
			Expected: uint32(121),
		},
		{
			Value:     int32(122),
			Expected: uint32(122),
		},
		{
			Value:     int32(123),
			Expected: uint32(123),
		},
		{
			Value:     int32(124),
			Expected: uint32(124),
		},
		{
			Value:     int32(125),
			Expected: uint32(125),
		},
		{
			Value:     int32(126),
			Expected: uint32(126),
		},
		{
			Value:     int32(127),
			Expected: uint32(127),
		},
		{
			Value:     int32(128),
			Expected: uint32(128),
		},
		{
			Value:     int32(129),
			Expected: uint32(129),
		},
		{
			Value:     int32(130),
			Expected: uint32(130),
		},
		{
			Value:     int32(131),
			Expected: uint32(131),
		},

		{
			Value:     int32(32759),
			Expected: uint32(32759),
		},
		{
			Value:     int32(32760),
			Expected: uint32(32760),
		},
		{
			Value:     int32(32761),
			Expected: uint32(32761),
		},
		{
			Value:     int32(32762),
			Expected: uint32(32762),
		},
		{
			Value:     int32(32763),
			Expected: uint32(32763),
		},
		{
			Value:     int32(32764),
			Expected: uint32(32764),
		},
		{
			Value:     int32(32765),
			Expected: uint32(32765),
		},
		{
			Value:     int32(32766),
			Expected: uint32(32766),
		},
		{
			Value:     int32(32767),
			Expected: uint32(32767),
		},

		{
			Value:     int32(2147483639),
			Expected: uint32(2147483639),
		},
		{
			Value:     int32(2147483640),
			Expected: uint32(2147483640),
		},
		{
			Value:     int32(2147483641),
			Expected: uint32(2147483641),
		},
		{
			Value:     int32(2147483642),
			Expected: uint32(2147483642),
		},
		{
			Value:     int32(2147483643),
			Expected: uint32(2147483643),
		},
		{
			Value:     int32(2147483644),
			Expected: uint32(2147483644),
		},
		{
			Value:     int32(2147483645),
			Expected: uint32(2147483645),
		},
		{
			Value:     int32(2147483646),
			Expected: uint32(2147483646),
		},
		{
			Value:     int32(2147483647),
			Expected: uint32(2147483647),
		},









		{
			Value:     int64(0),
			Expected: uint32(0),
		},
		{
			Value:     int64(1),
			Expected: uint32(1),
		},
		{
			Value:     int64(2),
			Expected: uint32(2),
		},
		{
			Value:     int64(3),
			Expected: uint32(3),
		},
		{
			Value:     int64(4),
			Expected: uint32(4),
		},
		{
			Value:     int64(6),
			Expected: uint32(6),
		},
		{
			Value:     int64(7),
			Expected: uint32(7),
		},
		{
			Value:     int64(8),
			Expected: uint32(8),
		},
		{
			Value:     int64(9),
			Expected: uint32(9),
		},
		{
			Value:     int64(10),
			Expected: uint32(10),
		},

		{
			Value:     int64(119),
			Expected: uint32(119),
		},
		{
			Value:     int64(120),
			Expected: uint32(120),
		},
		{
			Value:     int64(121),
			Expected: uint32(121),
		},
		{
			Value:     int64(122),
			Expected: uint32(122),
		},
		{
			Value:     int64(123),
			Expected: uint32(123),
		},
		{
			Value:     int64(124),
			Expected: uint32(124),
		},
		{
			Value:     int64(125),
			Expected: uint32(125),
		},
		{
			Value:     int64(126),
			Expected: uint32(126),
		},
		{
			Value:     int64(127),
			Expected: uint32(127),
		},
		{
			Value:     int64(128),
			Expected: uint32(128),
		},
		{
			Value:     int64(129),
			Expected: uint32(129),
		},
		{
			Value:     int64(130),
			Expected: uint32(130),
		},
		{
			Value:     int64(131),
			Expected: uint32(131),
		},

		{
			Value:     int64(32759),
			Expected: uint32(32759),
		},
		{
			Value:     int64(32760),
			Expected: uint32(32760),
		},
		{
			Value:     int64(32761),
			Expected: uint32(32761),
		},
		{
			Value:     int64(32762),
			Expected: uint32(32762),
		},
		{
			Value:     int64(32763),
			Expected: uint32(32763),
		},
		{
			Value:     int64(32764),
			Expected: uint32(32764),
		},
		{
			Value:     int64(32765),
			Expected: uint32(32765),
		},
		{
			Value:     int64(32766),
			Expected: uint32(32766),
		},
		{
			Value:     int64(32767),
			Expected: uint32(32767),
		},

		{
			Value:     int64(2147483639),
			Expected: uint32(2147483639),
		},
		{
			Value:     int64(2147483640),
			Expected: uint32(2147483640),
		},
		{
			Value:     int64(2147483641),
			Expected: uint32(2147483641),
		},
		{
			Value:     int64(2147483642),
			Expected: uint32(2147483642),
		},
		{
			Value:     int64(2147483643),
			Expected: uint32(2147483643),
		},
		{
			Value:     int64(2147483644),
			Expected: uint32(2147483644),
		},
		{
			Value:     int64(2147483645),
			Expected: uint32(2147483645),
		},
		{
			Value:     int64(2147483646),
			Expected: uint32(2147483646),
		},
		{
			Value:     int64(2147483647),
			Expected: uint32(2147483647),
		},
		{
			Value:     int64(2147483648),
			Expected: uint32(2147483648),
		},
		{
			Value:     int64(2147483649),
			Expected: uint32(2147483649),
		},
		{
			Value:     int64(2147483650),
			Expected: uint32(2147483650),
		},
		{
			Value:     int64(2147483651),
			Expected: uint32(2147483651),
		},

		{
			Value:     int64(4294967289),
			Expected: uint32(4294967289),
		},
		{
			Value:     int64(4294967290),
			Expected: uint32(4294967290),
		},
		{
			Value:     int64(4294967291),
			Expected: uint32(4294967291),
		},
		{
			Value:     int64(4294967292),
			Expected: uint32(4294967292),
		},
		{
			Value:     int64(4294967293),
			Expected: uint32(4294967293),
		},
		{
			Value:     int64(4294967294),
			Expected: uint32(4294967294),
		},
		{
			Value:     int64(4294967295),
			Expected: uint32(4294967295),
		},









		{
			Value:     uint8(0),
			Expected: uint32(0),
		},
		{
			Value:     uint8(1),
			Expected: uint32(1),
		},
		{
			Value:     uint8(2),
			Expected: uint32(2),
		},
		{
			Value:     uint8(3),
			Expected: uint32(3),
		},
		{
			Value:     uint8(4),
			Expected: uint32(4),
		},
		{
			Value:     uint8(6),
			Expected: uint32(6),
		},
		{
			Value:     uint8(7),
			Expected: uint32(7),
		},
		{
			Value:     uint8(8),
			Expected: uint32(8),
		},
		{
			Value:     uint8(9),
			Expected: uint32(9),
		},
		{
			Value:     uint8(10),
			Expected: uint32(10),
		},

		{
			Value:     uint8(249),
			Expected: uint32(249),
		},
		{
			Value:     uint8(250),
			Expected: uint32(250),
		},
		{
			Value:     uint8(251),
			Expected: uint32(251),
		},
		{
			Value:     uint8(252),
			Expected: uint32(252),
		},
		{
			Value:     uint8(253),
			Expected: uint32(253),
		},
		{
			Value:     uint8(254),
			Expected: uint32(254),
		},
		{
			Value:     uint8(255),
			Expected: uint32(255),
		},









		{
			Value:    uint16(0),
			Expected: uint32(0),
		},
		{
			Value:    uint16(1),
			Expected: uint32(1),
		},
		{
			Value:    uint16(2),
			Expected: uint32(2),
		},
		{
			Value:    uint16(3),
			Expected: uint32(3),
		},
		{
			Value:    uint16(4),
			Expected: uint32(4),
		},
		{
			Value:    uint16(6),
			Expected: uint32(6),
		},
		{
			Value:    uint16(7),
			Expected: uint32(7),
		},
		{
			Value:    uint16(8),
			Expected: uint32(8),
		},
		{
			Value:    uint16(9),
			Expected: uint32(9),
		},
		{
			Value:    uint16(10),
			Expected: uint32(10),
		},

		{
			Value:    uint16(249),
			Expected: uint32(249),
		},
		{
			Value:    uint16(250),
			Expected: uint32(250),
		},
		{
			Value:    uint16(251),
			Expected: uint32(251),
		},
		{
			Value:    uint16(252),
			Expected: uint32(252),
		},
		{
			Value:    uint16(253),
			Expected: uint32(253),
		},
		{
			Value:    uint16(254),
			Expected: uint32(254),
		},
		{
			Value:    uint16(255),
			Expected: uint32(255),
		},
		{
			Value:    uint16(256),
			Expected: uint32(256),
		},
		{
			Value:    uint16(257),
			Expected: uint32(257),
		},
		{
			Value:    uint16(258),
			Expected: uint32(258),
		},
		{
			Value:    uint16(259),
			Expected: uint32(259),
		},
		{
			Value:    uint16(260),
			Expected: uint32(260),
		},
		{
			Value:    uint16(261),
			Expected: uint32(261),
		},

		{
			Value:    uint16(65529),
			Expected: uint32(65529),
		},
		{
			Value:    uint16(65530),
			Expected: uint32(65530),
		},
		{
			Value:    uint16(65531),
			Expected: uint32(65531),
		},
		{
			Value:    uint16(65532),
			Expected: uint32(65532),
		},
		{
			Value:    uint16(65533),
			Expected: uint32(65533),
		},
		{
			Value:    uint16(65534),
			Expected: uint32(65534),
		},
		{
			Value:    uint16(65535),
			Expected: uint32(65535),
		},









		{
			Value:    uint32(0),
			Expected: uint32(0),
		},
		{
			Value:    uint32(1),
			Expected: uint32(1),
		},
		{
			Value:    uint32(2),
			Expected: uint32(2),
		},
		{
			Value:    uint32(3),
			Expected: uint32(3),
		},
		{
			Value:    uint32(4),
			Expected: uint32(4),
		},
		{
			Value:    uint32(6),
			Expected: uint32(6),
		},
		{
			Value:    uint32(7),
			Expected: uint32(7),
		},
		{
			Value:    uint32(8),
			Expected: uint32(8),
		},
		{
			Value:    uint32(9),
			Expected: uint32(9),
		},
		{
			Value:    uint32(10),
			Expected: uint32(10),
		},

		{
			Value:    uint32(249),
			Expected: uint32(249),
		},
		{
			Value:    uint32(250),
			Expected: uint32(250),
		},
		{
			Value:    uint32(251),
			Expected: uint32(251),
		},
		{
			Value:    uint32(252),
			Expected: uint32(252),
		},
		{
			Value:    uint32(253),
			Expected: uint32(253),
		},
		{
			Value:    uint32(254),
			Expected: uint32(254),
		},
		{
			Value:    uint32(255),
			Expected: uint32(255),
		},
		{
			Value:    uint32(256),
			Expected: uint32(256),
		},
		{
			Value:    uint32(257),
			Expected: uint32(257),
		},
		{
			Value:    uint32(258),
			Expected: uint32(258),
		},
		{
			Value:    uint32(259),
			Expected: uint32(259),
		},
		{
			Value:    uint32(260),
			Expected: uint32(260),
		},
		{
			Value:    uint32(261),
			Expected: uint32(261),
		},

		{
			Value:    uint32(65529),
			Expected: uint32(65529),
		},
		{
			Value:    uint32(65530),
			Expected: uint32(65530),
		},
		{
			Value:    uint32(65531),
			Expected: uint32(65531),
		},
		{
			Value:    uint32(65532),
			Expected: uint32(65532),
		},
		{
			Value:    uint32(65533),
			Expected: uint32(65533),
		},
		{
			Value:    uint32(65534),
			Expected: uint32(65534),
		},
		{
			Value:    uint32(65535),
			Expected: uint32(65535),
		},
		{
			Value:    uint32(65536),
			Expected: uint32(65536),
		},
		{
			Value:    uint32(65537),
			Expected: uint32(65537),
		},
		{
			Value:    uint32(65538),
			Expected: uint32(65538),
		},
		{
			Value:    uint32(65539),
			Expected: uint32(65539),
		},
		{
			Value:    uint32(65540),
			Expected: uint32(65540),
		},
		{
			Value:    uint32(65541),
			Expected: uint32(65541),
		},

		{
			Value:    uint32(4294967289),
			Expected: uint32(4294967289),
		},
		{
			Value:    uint32(4294967290),
			Expected: uint32(4294967290),
		},
		{
			Value:    uint32(4294967291),
			Expected: uint32(4294967291),
		},
		{
			Value:    uint32(4294967292),
			Expected: uint32(4294967292),
		},
		{
			Value:    uint32(4294967293),
			Expected: uint32(4294967293),
		},
		{
			Value:    uint32(4294967294),
			Expected: uint32(4294967294),
		},
		{
			Value:    uint32(4294967295),
			Expected: uint32(4294967295),
		},









		{
			Value:    uint64(0),
			Expected: uint32(0),
		},
		{
			Value:    uint64(1),
			Expected: uint32(1),
		},
		{
			Value:    uint64(2),
			Expected: uint32(2),
		},
		{
			Value:    uint64(3),
			Expected: uint32(3),
		},
		{
			Value:    uint64(4),
			Expected: uint32(4),
		},
		{
			Value:    uint64(6),
			Expected: uint32(6),
		},
		{
			Value:    uint64(7),
			Expected: uint32(7),
		},
		{
			Value:    uint64(8),
			Expected: uint32(8),
		},
		{
			Value:    uint64(9),
			Expected: uint32(9),
		},
		{
			Value:    uint64(10),
			Expected: uint32(10),
		},

		{
			Value:    uint64(249),
			Expected: uint32(249),
		},
		{
			Value:    uint64(250),
			Expected: uint32(250),
		},
		{
			Value:    uint64(251),
			Expected: uint32(251),
		},
		{
			Value:    uint64(252),
			Expected: uint32(252),
		},
		{
			Value:    uint64(253),
			Expected: uint32(253),
		},
		{
			Value:    uint64(254),
			Expected: uint32(254),
		},
		{
			Value:    uint64(255),
			Expected: uint32(255),
		},
		{
			Value:    uint64(256),
			Expected: uint32(256),
		},
		{
			Value:    uint64(257),
			Expected: uint32(257),
		},
		{
			Value:    uint64(258),
			Expected: uint32(258),
		},
		{
			Value:    uint64(259),
			Expected: uint32(259),
		},
		{
			Value:    uint64(260),
			Expected: uint32(260),
		},
		{
			Value:    uint64(261),
			Expected: uint32(261),
		},

		{
			Value:    uint64(65529),
			Expected: uint32(65529),
		},
		{
			Value:    uint64(65530),
			Expected: uint32(65530),
		},
		{
			Value:    uint64(65531),
			Expected: uint32(65531),
		},
		{
			Value:    uint64(65532),
			Expected: uint32(65532),
		},
		{
			Value:    uint64(65533),
			Expected: uint32(65533),
		},
		{
			Value:    uint64(65534),
			Expected: uint32(65534),
		},
		{
			Value:    uint64(65535),
			Expected: uint32(65535),
		},
		{
			Value:    uint64(65536),
			Expected: uint32(65536),
		},
		{
			Value:    uint64(65537),
			Expected: uint32(65537),
		},
		{
			Value:    uint64(65538),
			Expected: uint32(65538),
		},
		{
			Value:    uint64(65539),
			Expected: uint32(65539),
		},
		{
			Value:    uint64(65540),
			Expected: uint32(65540),
		},
		{
			Value:    uint64(65541),
			Expected: uint32(65541),
		},

		{
			Value:    uint64(4294967289),
			Expected: uint32(4294967289),
		},
		{
			Value:    uint64(4294967290),
			Expected: uint32(4294967290),
		},
		{
			Value:    uint64(4294967291),
			Expected: uint32(4294967291),
		},
		{
			Value:    uint64(4294967292),
			Expected: uint32(4294967292),
		},
		{
			Value:    uint64(4294967293),
			Expected: uint32(4294967293),
		},
		{
			Value:    uint64(4294967294),
			Expected: uint32(4294967294),
		},
		{
			Value:    uint64(4294967295),
			Expected: uint32(4294967295),
		},
	}


	for testNumber, test := range tests {

		var datum Type

		if err := datum.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, datum.Code(); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}
	}
}

func TestTypeScanError(t *testing.T) {

	tests := []struct{
		Value interface{}
	}{
		{
			Value: []byte{},
		},
		{
			Value: []byte(nil),
		},
		{
			Value: []byte(""),
		},
		{
			Value: []byte("FB_VISUAL_APPLE_BANANA_CHERRY"),
		},



		{
			Value: "",
		},
		{
			Value: "FB_VISUAL_APPLE_BANANA_CHERRY",
		},



		{
			Value: int8(-128),
		},
		{
			Value: int8(-1),
		},



		{
			Value: int16(-32768),
		},
		{
			Value: int16(-1),
		},



		{
			Value: int32(-2147483648),
		},
		{
			Value: int32(-1),
		},



		{
			Value: int64(-9223372036854775808),
		},
		{
			Value: int64(-1),
		},
		{
			Value: int64(4294967296),
		},
		{
			Value: int64(9223372036854775807),
		},



		{
			Value: uint64(4294967296),
		},
		{
			Value: uint64(18446744073709551615),
		},



		{
			Value: false,
		},
		{
			Value: true,
		},



		{
			Value: -1234.5678,
		},
	}


	for testNumber, test := range tests {

		var datum Type

		if err := datum.Scan(test.Value); nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one, for: (%T) %#v", testNumber, test.Value, test.Value)
			continue
		}
	}
}
