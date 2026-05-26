package announcements

type TextMarshaler struct {
	Value *Announcement
}

func (a *TextMarshaler) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *TextMarshaler) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }

// DecodeFromBase64 decodes announcement from base64 string, validates it and returns.
func DecodeFromBase64(payload string) (*Announcement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// auto-compute announcement key

// Encode encodes announcement message into a msgpack base64 string.
func Encode(msg *Announcement) (string, error) { _ = "STUB: not implemented"; return "", nil }
