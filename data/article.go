package data

type CommandCenterArticleDto struct {
	ID     int
	Author string
}

type GetCommentedArticlesReq struct {
	Cursor int `json:"cursor"`
	Size   int `json:"size"`
}

type GetCommentedArticlesResp struct {
	Data    []int  `json:"data"`
	Message string `json:"message"`
}
