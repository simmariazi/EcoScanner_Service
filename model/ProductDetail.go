package model

type ProductDetail struct {
	ProductId     int
	ProductCode   string
	ProductName   string
	ProductUrl    string
	Description   string
	IsUsed        int
	Detail        DetailInfo
	Delivery      DeliveryInfo
	ProductPrice  int
	ProductOption string
	Seller        SellerInfo
	EcoCert       EcoCertification
	Thumbnail     string
}
