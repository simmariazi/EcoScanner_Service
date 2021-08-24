package model

type Board struct {
	MemberNo int    `json:"memberNo"`
	Title    string `json:"title"`
	Contents string `json:"contents"`
}
