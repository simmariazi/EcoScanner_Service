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
		productSimple.CategoryId = products[i].Category_id
		productSimple.CategoryName = db.FindCategoryIdByName(products[i].Category_id)

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

	var sellerInfo model.SellerInfo

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
		productDetail.Thumbnail = products[i].Mainimage
		productDetail.Description = products[i].Description
		productDetail.Detail = products[i].Detail
		productDetail.DeliveryTime = products[i].DeliveryTime
		productDetail.ShippingFee = products[i].ShippingFee
		productDetail.ProductPrice = products[i].Price
		productDetail.ProductOption = products[i].Option
		productDetail.Ecocertification = products[i].Eco_certification
		productDetail.ProductUrl = products[i].Product_url

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

func GetCompareProduct(productId string) []model.ProductDetail {
	return db.CompareProductDetail(productId)
}
