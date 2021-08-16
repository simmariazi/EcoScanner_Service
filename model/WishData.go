package model

type WishListData struct {
	Itemtype  string `db:"type"`
	Name      string `db:"name"`
	Thumbnail string `db:"thumnail"`
	Price     int    `db:"price"`
	EcoInfo   string `db:"ecoInfo"`
	Url       string `db:"url"`
}
