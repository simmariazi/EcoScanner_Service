package model

type Product struct {
	ProductId        int
	ProductName      string
	ProductThumbnail string
	ProductUrl       string
	ProductPrice     int
	Seller           SellerInfo
	IsUsed           int
	CategoryId       int
	CategoryName     string
}
