package function

import (
	db "main/database"
	"main/model"
)

func GetBoardList() []model.Recommend {
	var boarddata []model.EntBoardRecommend
	var recommend model.Recommend
	var recommends []model.Recommend

	for i := 0; i < len(boarddata); i++ {

		recommend.Id = boarddata[i].Id
		recommend.Nickname = db.FindMemberNameByMemberNo(boarddata[i].Member_no)
		recommend.Title = boarddata[i].Title
		recommend.CreateDate = boarddata[i].Create_date

		recommends = append(recommends, recommend)

	}

	return recommends

}

func WriteBoardPost(memberNo int, title string, contents string) {
	db.AddRecommendPost(memberNo, title, contents)
}

func GetBoardPost(boardId int) model.Recommend {
	return db.FindPostById(boardId)
}

func ModifyBoardPost(boardId int, memberNo int, title string, contents string) {
	db.ModifyRecommendPost(boardId, memberNo, title, contents)
}
