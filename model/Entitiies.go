package model

type Seller struct {
	Id          int
	Seller_name string
	Seller_url  string
	Create_date string
}

type Member struct {
	Member_no   int
	Member_id   string
	Nickname    string
	IsUsed      bool
	Update_date string
	Create_date string
}
