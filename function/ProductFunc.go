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

		productSimple.IsUsed = 1

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

func GetProductDetailList() []model.ProductDetail {
	var products []model.EntProductDetail
	var productDetail model.ProductDetail
	var productsDetail []model.ProductDetail
	var detailInfo model.DetailInfo
	var deliveryInfo model.DeliveryInfo
	var sellerInfo model.SellerInfo
	var ecoCertification model.EcoCertification

	products = db.CallProductDetailSelection()

	for i := 0; i < len(products); i++ {

		// isUsed 0일때 미진열
		if products[i].Status == 0 {
			continue
		}

		productDetail.IsUsed = 1

		productDetail.ProductId = products[i].Id
		productDetail.ProductName = products[i].Name
		productDetail.ProductCode = products[i].Productcode
		productDetail.Thumnail = products[i].Mainimage
		productDetail.Description = products[i].Description
		productDetail.ProductPrice = products[i].Price
		productDetail.ProductOption = products[i].Option

		//DeliveryInfo
		deliveryInfo.DeliveryId = products[i].Delivery_id
		deliveryInfo.DeliveryTime = (db.FindDeliveryInfoByDeliveryId(products[i].Delivery_id).DeliveryTime)
		deliveryInfo.ShippingFee = (db.FindDeliveryInfoByDeliveryId(products[i].Delivery_id).ShippingFee)

		// DetailInfo -> 컬럼 3개 아니어서 물어봐야됨

		// Ecocertification info

		// SellerInfo
		sellerInfo.SellerId = products[i].Seller_id
		sellerInfo.SellerName = GetSellerName(products[i].Seller_id)
		// 상품 조회시에는 빈 값
		sellerInfo.SellerIntroduction = ""
		productDetail.Seller = sellerInfo

		productsDetail = append(productsDetail, productDetail)
	}

	return productsDetail
}

func GetProductName(productId int) string {
	return db.FindProductNameById(productId)
}
