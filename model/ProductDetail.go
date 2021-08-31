package model

type ProductDetail struct {
	ProductId        int
	ProductCode      string
	ProductName      string
	ProductUrl       string
	Description      string
	IsUsed           int
	Detail           string
	DeliveryTime     string
	ShippingFee      string
	ProductPrice     int
	ProductOption    string
	Seller           SellerInfo
	Thumbnail        string
	Ecocertification int
}
