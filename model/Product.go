package model

type Product struct {
	ProductId        int
	ProductName      string
	ProductThumbnail string
	ProductUrl       string
	ProductPrice     int
	SellerInfo       SellerInfo
	EcoCertification EcoCertification
}
