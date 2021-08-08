package function

import (
	db "main/database"
	"main/model"
)

func ModifyMemberInfo(isAuth bool, memberNo int, nickname string) bool {
	var result bool = false

	// sellers := GetSellers()
	// 멤버 정보 수정 로직 작성

	return result
}

func DeleteMemberInfo(isAuth bool, memberNo int) bool {
	var result bool = false

	// 멤버 정보 삭제 로직 작성(soft delete)

	return result
}

func GetMemberList() []model.MemberInfo {
	var members []model.EntMember
	var memberInfo model.MemberInfo
	var memberInfos []model.MemberInfo
	members = db.CallMemberSelection()

	for i := 0; i < len(members); i++ {
		memberInfo.MemberId = members[i].Member_id
		memberInfo.MemberNo = members[i].Member_no
		memberInfo.Nickname = members[i].Nickname

		memberInfos = append(memberInfos, memberInfo)
	}

	return memberInfos

}

// func LoopObjectField(object interface{}) {
// 	e := reflect.ValueOf(object).Elem()
// 	fieldNum := e.NumField()
// 	for i := 0; i < fieldNum; i ++{
// 		v := e.Field(i)
// 		t := e.Type().Field(i)
// 		fmt.Printf("Name: %s / Type: %s / Value: %v / Tag: %s \n",
// 			t.Name, t.Type, v.Interface(), t.Tag.Get("custom"))
// 	}
// }

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

func GetProductSimpleList() []model.Product {
	var products []model.EntProduct
	var productSimple model.Product
	var productsSimple []model.Product

	products = db.CallProductSelection()

	for i := 0; i < len(products); i++ {
		productSimple.ProductId = products[i].Id
		productSimple.ProductUrl = products[i].ProductUrl

		productsSimple = append(productsSimple, productSimple)
	}

	return productsSimple

}
