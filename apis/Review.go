package apis

import (
	"encoding/json"
	"main/function"
	"main/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetReviews godoc
// @tags Review
// @summary 후기 목록 조회
// @Accept  json
// @Produce json
// @Router	/review [get]
// @Success 200 "Success"
func GetReviews(c *gin.Context) {
	c.JSON(200, function.GetReviewList())
}

// WriteReview godoc
// @tags Review
// @summary 후기 작성
// @Param memberNo body int true "회원 번호"
// @Param productId body int true "SKU"
// @Param contents body string true "후기 내용"
// @Param reviewRating body int true "후기 평점"
// @Accept  json
// @Produce json
// @Router	/review/post [put]
// @Success 200 "Success"
func WriteReview(c *gin.Context) {
	var review model.ReviewData
	json.NewDecoder(c.Request.Body).Decode(&review)

	memberNo := review.MemberNo
	productId := review.ProductId
	contents := review.Contents
	reviewRating := review.ReviewRating

	function.WriteReviewPost(memberNo, productId, contents, reviewRating)
	c.JSON(200, "success")
}

// ModifyReview godoc
// @tags Review
// @summary 후기 수정
// @Param id body int true "후기 게시글 번호"
// @Param memberNo body int true "회원 번호"
// @Param contents body string true "후기 내용"
// @Param reviewRating body int true "후기 평점"
// @Accept  json
// @Produce json
// @Router	/review/post [post]
// @Success 200 "Success"
func ModifyReview(c *gin.Context) {
	var review model.ReviewUpdate
	json.NewDecoder(c.Request.Body).Decode(&review)

	id := review.Id
	memberNo := review.MemberNo
	contents := review.Contents
	reviewRating := review.ReviewRating

	function.ModifyReviewPost(id, memberNo, contents, reviewRating)
	c.JSON(200, "success")
}

// DeleteReview godoc
// @tags Review
// @summary 후기 삭제
// @Param reviewId query int true "후기 게시글 번호"
// @Param memberNo query int true "회원 번호"
// @Accept  json
// @Produce json
// @Router	/review [delete]
// @Success 200 "Success"
func DeleteReview(c *gin.Context) {
	reviewId, _ := strconv.Atoi(c.Query("reviewId"))
	memberNo, _ := strconv.Atoi(c.Query("memberNo"))
	function.DeleteReviewPost(reviewId, memberNo)
	c.JSON(200, "삭제 완료")
}
