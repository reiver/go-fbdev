/*
Package fb_visual provides Go equivalents, for the following in the C based API to the Frame Buffer Device:

• FB_VISUAL_MONO01

• FB_VISUAL_MONO10

• FB_VISUAL_TRUECOLOR

• FB_VISUAL_PSEUDOCOLOR

• FB_VISUAL_DIRECTCOLOR

• FB_VISUAL_STATIC_PSEUDOCOLOR

• FB_VISUAL_FOURCC

It also provides a type to help working with this type of information easier.


Example Usage

	var visual fb_visual.Type
	
	// ...
	
	switch visual {
	case fb_visual.CodeMono01:
		//@TODO
		
	case fb_visual.CodeMono10:
		//@TODO
		
	case fb_visual.CodeTrueColor:
		//@TODO
		
	case fb_visual.CodePseudoColor:
		//@TODO
		
	case fb_visual.CodeDirectColor:
		//@TODO
		
	case fb_visual.CodeStaticPseudoColor:
		//@TODO
		
	case fb_visual.CodeFourCC:
		//@TODO
		
	default:
		//@TODO
		
	}


Compatibility

Note that this will work seamlessly with "database/sql", and anything that is compatibile with encoding.TextMarshaler & encoding.TextUnmarshaler, such as "encoding/json".

*/
package fb_visual
