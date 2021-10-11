package model

type ProductDetail struct {
	ProductId        int
	ProductCode      string `json:"ProductCode,omitempty"`
	ProductName      string
	ProductUrl       string
	Description      string `json:"Description,omitempty"`
	IsUsed           int    `json:"IsUsed,omitempty"`
	Detail           string `json:"Detail,omitempty"`
	DeliveryTime     string `json:"DeliveryTime,omitempty"`
	ShippingFee      string `json:"ShippingFee,omitempty"`
	ProductPrice     int
	ProductOption    string     `json:"ProductOption,omitempty"`
	Seller           SellerInfo `json:"Seller,omitempty"`
	Thumbnail        string
	Ecocertification string `json:"Ecocertification,omitempty"`
}
