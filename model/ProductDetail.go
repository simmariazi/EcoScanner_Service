package model

type ProductDetail struct {
	productId        int
	productCode      string
	productName      string
	productUrl       string
	description      string
	detail           Detail
	deliveryInfo     DeliveryInfo
	productPrice     int
	productOption    []string
	sellerInfo       SellerInfo
	ecoCertification EcoCertification
}
