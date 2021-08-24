package model

type BoardUpdate struct {
	Id       int    `json:"id"`
	MemberNo int    `json:"memberNo"`
	Title    string `json:"title"`
	Contents string `json:"contents"`
}
