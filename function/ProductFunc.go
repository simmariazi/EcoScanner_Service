package function

import (
	db "main/database"
	"main/model"
)

func GetProductSimpleList(page int, count int) []model.Product {
	var products []model.EntProductList
	var productSimple model.Product
	var productsSimple []model.Product
	var sellerInfo model.SellerInfo

	products = db.CallProductSimpleSelection(page, count)

	for i := 0; i < len(products); i++ {

		// isUsed 0일때 미진열
		if products[i].Is_used == 0 {
			continue
		}

		productSimple.IsUsed = 1

		productSimple.ProductId = products[i].Id
		productSimple.ProductThumbnail = products[i].Thumbnail
		productSimple.ProductUrl = products[i].ProductUrl
		productSimple.ProductPrice = products[i].ProductPrice
		productSimple.ProductName = products[i].ProductName
		productSimple.CategoryId = products[i].Category_id
		productSimple.CategoryName = products[i].Category_name

		// SellerInfo
		sellerInfo.SellerId = products[i].Seller_id
		sellerInfo.SellerName = products[i].Seller_name
		// 상품 조회시에는 빈 값
		sellerInfo.SellerIntroduction = ""
		productSimple.Seller = sellerInfo

		productsSimple = append(productsSimple, productSimple)
	}

	return productsSimple
}

func GetProductDetail(productId int) model.ProductDetail {
	var products model.EntProductDetail
	var productDetail model.ProductDetail
	var sellerInfo model.SellerInfo

	products = db.CallProductDetailSelection(productId)

	// isUsed 0일때 미진열
	if products.Status == 0 {
		return productDetail
	}

	productDetail.IsUsed = 1

	productDetail.ProductId = products.Id
	productDetail.ProductName = products.Name
	productDetail.ProductCode = products.Productcode
	productDetail.Thumbnail = products.Mainimage
	productDetail.Description = products.Description
	productDetail.Detail = products.Detail
	productDetail.DeliveryTime = products.DeliveryTime
	productDetail.ShippingFee = products.ShippingFee
	productDetail.ProductPrice = products.Price
	productDetail.ProductOption = products.Option
	productDetail.Ecocertification = products.Eco_certification
	productDetail.ProductUrl = products.Product_url

	// SellerInfo
	sellerInfo.SellerId = products.Seller_id
	sellerInfo.SellerName = GetSellerName(products.Seller_id)

	// 상품 조회시에는 빈 값
	sellerInfo.SellerIntroduction = ""
	productDetail.Seller = sellerInfo

	return productDetail
}

func GetProductName(productId int) string {
	return db.FindProductNameById(productId)
}

func GetCompareProduct(productId string) []model.ProductDetail {
	return db.CompareProductDetail(productId)
}

func SearchProductsByProductName(productName string) []model.Product {
	return db.SearchProductsByProductName(productName)
}
