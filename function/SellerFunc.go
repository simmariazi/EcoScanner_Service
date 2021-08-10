package function

import (
	db "main/database"
	"main/model"
)

func GetSellerList() []model.SellerInfo {
	var sellers []model.EntSeller
	var sellerinfo model.SellerInfo
	var sellerinfos []model.SellerInfo

	sellers = db.CallSellerSelection()

	for i := 0; i < len(sellers); i++ {
		sellerinfo.SellerId = sellers[i].Id
		sellerinfo.SellerName = sellers[i].Seller_name
		sellerinfo.SellerIntroduction = sellers[i].Seller_url

		sellerinfos = append(sellerinfos, sellerinfo)
	}

	return sellerinfos
}

func GetSellerName(sellerId int) string {
	return db.FindSellerNameById(sellerId)
}
