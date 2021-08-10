package model

type EntSeller struct {
	Id          int
	Seller_name string
	Seller_url  string
	Create_date string
}

type EntMember struct {
	Member_no   int
	Member_id   string
	Nickname    string
	IsUsed      bool
	Update_date string
	Create_date string
}

type EntProductList struct {
	Id         int
	Thumnail   string
	ProductUrl string
	Seller_id  int
	Is_used    int
}
