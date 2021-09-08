package apis

import (
	"encoding/json"
	"main/function"
	"main/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBoards godoc
// @tags Board
// @summary 게시판 목록 조회
// @Accept  json
// @Produce json
// @Router	/board
// @Success 200 "Success"
func GetBoards(c *gin.Context) {
	c.JSON(200, function.GetBoardList())
}

// WriteBoard godoc
// @tags Board
// @summary 게시글 작성
// @Param memberNo body string true "회원 번호"
// @Param title body string true "게시글 제목"
// @Param contents body string true "게시글 내용"
// @Accept  json
// @Produce json
// @Router	/board/post
// @Success 200 "Success"
func WriteBoard(c *gin.Context) {
	var board model.Board
	json.NewDecoder(c.Request.Body).Decode(&board)

	memberNo := board.MemberNo
	title := board.Title
	contents := board.Contents

	function.WriteBoardPost(memberNo, title, contents)
	c.JSON(200, "success")
}

// GetBoard godoc
// @tags Board
// @summary 게시글 조회
// @Param boardId query string true "게시글 번호"
// @Accept  json
// @Produce json
// @Router	/board/post
// @Success 200 "Success"
func GetBoard(c *gin.Context) {
	boardId, _ := strconv.Atoi(c.Query("boardId"))
	c.JSON(200, function.GetBoardPost(boardId))
}

// ModifyBoard godoc
// @tags Board
// @summary 게시글 수정
// @Param id body string true "게시글 번호"
// @Param memberNo body string true "회원 번호"
// @Param title body string true "게시글 제목"
// @Param contents body string true "게시글 내용"
// @Accept  json
// @Produce json
// @Router	/board/post
// @Success 200 "Success"
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

// DeleteBoard godoc
// @tags Board
// @summary 게시글 삭제
// @Param boardId query string true "게시글 번호"
// @Param memberNo query string true "회원 번호"
// @Accept  json
// @Produce json
// @Router	/board
// @Success 200 "Success"
func DeleteBoard(c *gin.Context) {
	boardId, _ := strconv.Atoi(c.Query("boardId"))
	memberNo, _ := strconv.Atoi(c.Query("memberNo"))
	function.DeleteBoardPost(boardId, memberNo)
	c.JSON(200, "삭제 완료")
}
