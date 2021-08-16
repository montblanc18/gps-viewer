package omega2plus

type Uplink struct {
	Payloads Payloads `json:"payloads" validate:"required"`
}

type Payloads struct {
	// あとで埋める
}
