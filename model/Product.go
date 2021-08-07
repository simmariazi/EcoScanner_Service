package model

type Product struct {
	productId        int
	productName      string
	productUrl       string
	productPrice     int
	sellerInfo       SellerInfo
	ecoCertification EcoCertification
}
