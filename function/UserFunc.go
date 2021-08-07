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

func SelectMemeber() []model.MemberInfo {
	var members []model.Member
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

	//return memberInfo
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
