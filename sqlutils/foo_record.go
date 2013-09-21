package sqlutils

type fooRecord struct {
	Id   int64  `json:"id" field:"id,primary"`
	Name string `json:"name" field:"name"`
	// Phone    string `json:"phone" field:"phone"`
	Type     string `json:"type"`
	Internal int    `json:-`
}
