package model

type Product struct {
	ProductId        int
	ProductName      string
	ProductUrl       string
	ProductPrice     int
	SellerInfo       SellerInfo
	EcoCertification EcoCertification
}
