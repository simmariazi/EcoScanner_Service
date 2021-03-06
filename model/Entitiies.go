package model

type EntSeller struct {
	Id                int
	Seller_name       string
	Seller_url        string
	Eco_certification int
	About_seller      string
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
	Id            int
	ProductName   string
	Thumbnail     string
	ProductUrl    string
	ProductPrice  int
	Seller_id     int
	Seller_name   string
	Category_id   int
	Category_name string
	Is_used       int
}

type EntProductDetail struct {
	Id                int
	Name              string
	Productcode       string
	Mainimage         string
	Description       string
	Detail            string
	DeliveryTime      string
	ShippingFee       string
	Price             int
	Option            string
	Seller_id         int
	Eco_certification string
	Status            int
	Product_url       string
	Update_date       string
	Create_date       string
}

type EntWishListProduct struct {
	MemberNo    int
	Product_id  int
	Create_date string
}

type EntWishListSeller struct {
	MemberNo    int
	Seller_id   int
	Create_date string
}

type EntBoardRecommend struct {
	Id          int
	Member_no   int
	Title       string
	Contents    string
	Create_date string
	Update_date string
}

type EntProductReview struct {
	Id            int
	Member_no     int
	Product_id    int
	Contents      string
	Review_rating int
	Create_date   string
	Update_date   string
}
