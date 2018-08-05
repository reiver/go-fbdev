package fb_visual

func (receiver Type) MarshalText() (text []byte, err error) {
	return []byte(receiver.String()), nil
}
