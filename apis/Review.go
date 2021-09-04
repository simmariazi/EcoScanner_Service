package apis

import (
	"encoding/json"
	"main/function"
	"main/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetReviews(c *gin.Context) {
	c.JSON(200, function.GetReviewList())
}

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

func DeleteReview(c *gin.Context) {
	reviewId, _ := strconv.Atoi(c.Query("reviewId"))
	memberNo, _ := strconv.Atoi(c.Query("memberNo"))
	function.DeleteReviewPost(reviewId, memberNo)
	c.JSON(200, "삭제 완료")
}
