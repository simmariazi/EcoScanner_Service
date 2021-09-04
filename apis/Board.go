package apis

import (
	"encoding/json"
	"main/function"
	"main/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBoards(c *gin.Context) {
	c.JSON(200, function.GetBoardList())
}

func WriteBoard(c *gin.Context) {
	var board model.Board
	json.NewDecoder(c.Request.Body).Decode(&board)

	memberNo := board.MemberNo
	title := board.Title
	contents := board.Contents

	function.WriteBoardPost(memberNo, title, contents)
	c.JSON(200, "success")
}

func GetBoard(c *gin.Context) {
	boardId, _ := strconv.Atoi(c.Query("boardId"))
	c.JSON(200, function.GetBoardPost(boardId))
}

func ModifyBoard(c *gin.Context) {
	var board model.BoardUpdate
	json.NewDecoder(c.Request.Body).Decode(&board)

	id := board.Id
	memberNo := board.MemberNo
	title := board.Title
	contents := board.Contents

	function.ModifyBoardPost(id, memberNo, title, contents)
	c.JSON(200, "success")
}

func DeleteBoard(c *gin.Context) {
	boardId, _ := strconv.Atoi(c.Query("boardId"))
	memberNo, _ := strconv.Atoi(c.Query("memberNo"))
	function.DeleteBoardPost(boardId, memberNo)
	c.JSON(200, "삭제 완료")
}
