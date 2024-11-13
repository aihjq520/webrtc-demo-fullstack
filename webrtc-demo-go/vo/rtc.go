package vo

type SDPRes struct {
	Sdp  string `json:"sdp"`
	Type string `json:"type"`
}

type RTCRes struct {
	Sdp       SDPRes               `json:"sdp"`
	Candidate struct{ key string } `json:"candidate"`
	Key       string               `json:"key"`
}
