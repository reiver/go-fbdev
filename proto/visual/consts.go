package fb_visual

const (
	CodeMono01            = uint32(0) // Monochr. 1=Black 0=White.           In the C based API to the Frame Buffer Device, this is equivalent to FB_VISUAL_MONO01.
	CodeMono10            = uint32(1) // Monochr. 1=White 0=Black.           In the C based API to the Frame Buffer Device, this is equivalent to FB_VISUAL_MONO10.
	CodeTrueColor         = uint32(2) // True color.                         In the C based API to the Frame Buffer Device, this is equivalent to FB_VISUAL_TRUECOLOR.
	CodePseudoColor       = uint32(3) // Pseudo color (like atari).          In the C based API to the Frame Buffer Device, this is equivalent to FB_VISUAL_PSEUDOCOLOR.
	CodeDirectColor       = uint32(4) // Direct color.                       In the C based API to the Frame Buffer Device, this is equivalent to FB_VISUAL_DIRECTCOLOR.
	CodeStaticPseudoColor = uint32(5) // Pseudo color readonly.              In the C based API to the Frame Buffer Device, this is equivalent to FB_VISUAL_STATIC_PSEUDOCOLOR.
	CodeFourCC            = uint32(6) // Visual identified by a V4L2 FOURCC. In the C based API to the Frame Buffer Device, this is equivalent to FB_VISUAL_FOURCC.
)
