package model

type ProductDetail struct {
	ProductId     int
	ProductCode   string
	ProductName   string
	ProductUrl    string
	Description   string
	Detail        Detail
	DeliveryInfo  DeliveryInfo
	ProductPrice  int
	ProductOption []string
	Seller        SellerInfo
	EcoCert       EcoCertification
}
