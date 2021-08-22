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

		recommend.BoardId = boarddata[i].Id
		recommend.Nickname = db.FindMemberNameByMemnerNo(boarddata[i].Member_no)
		recommend.Title = boarddata[i].Title
		recommend.CreateDate = boarddata[i].Create_date

		recommends = append(recommends, recommend)

	}

	return recommends

}
