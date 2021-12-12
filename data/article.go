package data

type CommandCenterArticleDto struct {
	ID     int
	Author string
}

type GetCommentedArticlesReq struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type GetCommentedArticlesResp struct {
	Data    []int  `json:"data"`
	Message string `json:"message"`
}
