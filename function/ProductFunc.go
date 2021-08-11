package function

import (
	db "main/database"
	"main/model"
)

func GetProductSimpleList() []model.Product {
	var products []model.EntProductList
	var productSimple model.Product
	var productsSimple []model.Product
	var sellerInfo model.SellerInfo

	products = db.CallProductSimpleSelection()

	for i := 0; i < len(products); i++ {

		// isUsed 0일때 미진열
		if products[i].Is_used == 0 {
			continue
		}

		productSimple.ProductId = products[i].Id
		productSimple.ProductThumbnail = products[i].Thumnail
		productSimple.ProductUrl = products[i].ProductUrl
		productSimple.ProductName = db.FindProductNameById(products[i].Id)

		// SellerInfo
		sellerInfo.SellerId = products[i].Seller_id
		sellerInfo.SellerName = GetSellerName(products[i].Seller_id)
		// 상품 조회시에는 빈 값
		sellerInfo.SellerIntroduction = ""
		productSimple.Seller = sellerInfo

		productsSimple = append(productsSimple, productSimple)
	}

	return productsSimple
}

func GetProductName(productId int) string {
	return db.FindProductNameById(productId)
}
