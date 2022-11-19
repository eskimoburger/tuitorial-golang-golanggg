package domain

type TodoReq struct {
	Task string `json:"task"`
}

type TodoRes struct {
	Id         uint   `json:"id"`
	Task       string `json:"task"`
	CreateDate string `json:"create_date"`
	UpdateDate string `json:"update_date"`
}
