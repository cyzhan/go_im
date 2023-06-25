package bo

type JwtPayload struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	VdID int64  `json:"vdID"`
	TS   int64  `json:"ts"`
}
