package khumu

type IsAuthorResp struct {
	Data struct {
		IsAuthor bool `json:"is_author"`
	} `json:"data"`
}

type IsAuthorReq struct {
	Author string `json:"author"`
}
