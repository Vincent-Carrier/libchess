package chess

func (sq Sq) MarshalText() ([]byte, error) {
	return []byte(sq.String()), nil
}

func (sq *Sq) UnmarshalText(bytes []byte) error {
	panic("implement me")
}
