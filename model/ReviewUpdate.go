package model

type ReviewUpdate struct {
	Id           int    `json:"id"`
	MemberNo     int    `json:"memberNo"`
	Contents     string `json:"contents"`
	ReviewRating int    `json:"reviewRating"`
}
