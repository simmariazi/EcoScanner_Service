package function

import (
	db "main/database"
	"main/model"
)

func GetReviewList() []model.ReviewData {

	var dbReview []model.EntProductReview
	var review model.ReviewData
	var reviews []model.ReviewData

	dbReview = db.CallReviewList()

	for i := 0; i < len(dbReview); i++ {

		review.ReviewId = dbReview[i].Id
		review.MemberNo = dbReview[i].Member_no
		review.ProductId = dbReview[i].Product_id
		review.Contents = dbReview[i].Contents
		review.ReviewRating = dbReview[i].Review_rating
		review.CreateDate = dbReview[i].Create_date

		reviews = append(reviews, review)

	}

	return reviews
}

func WriteReviewPost(memberNo int, productId int, contents string, reviewRating int) {
	db.AddReviewPost(memberNo, productId, contents, reviewRating)
}

func ModifyReviewPost(reviewId int, memberNo int, contents string, reviewRating int) {
	db.ModifyReviewPost(reviewId, memberNo, contents, reviewRating)
}

func DeleteReviewPost(reviewId int, memberNo int) {
	db.DeleteReviewPost(reviewId, memberNo)
}
