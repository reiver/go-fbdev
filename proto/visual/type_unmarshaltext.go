package fb_visual

func (receiver *Type) UnmarshalText(text []byte) error {
	return receiver.Scan(text)
}
