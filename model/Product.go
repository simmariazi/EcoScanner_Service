package model

type Product struct {
	ProductId        int
	ProductName      string
	ProductThumbnail string
	ProductUrl       string
	ProductPrice     int
	Seller           SellerInfo
	EcoCert          EcoCertification
	IsUsed           int
}
