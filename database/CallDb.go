package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	model "main/model"

	_ "github.com/go-sql-driver/mysql"
)

var connectionString string = ""

func InitConnectionString(connection string) {
	connectionString = connection
}

func CallMemberSelection() []model.EntMember {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	var result model.EntMember
	var results []model.EntMember

	rows, err := db.Query("SELECT * FROM member")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.Member_no, &result.Member_id, &result.Nickname, &result.IsUsed, &result.Update_date, &result.Create_date)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, result)
		// fmt.Println("rows", seller.Id, seller.Seller_name, seller.Seller_url, seller.Create_date)
	}

	return results
}

func CallSellerSelection() []model.EntSeller {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	var result model.EntSeller
	var results []model.EntSeller
	query := "SELECT id, seller_name, seller_url, eco_certification, about_seller FROM seller"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Seller_name, &result.Seller_url, &result.Eco_certification, &result.About_seller)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, result)

	}

	return results
}

func CallProductSimpleSelection(page int) []model.EntProductList {

	firstnumber := (page-1)*12 + 1

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	var result model.EntProductList
	var results []model.EntProductList

	var rows *sql.Rows

	if page == 0 {

		rows, err = db.Query("SELECT pl.id, p.name, pl.thumbnail, p.price, pl.productUrl, pl.seller_id, s.seller_name, pl.category_id, c.name, pl.is_used FROM product_list pl ,product p ,category c ,seller s where p.id = pl.id AND s.id = pl.seller_id AND pl.category_id = c.id")

	} else {

		rows, err = db.Query("SELECT pl.id, p.name, pl.thumbnail, p.price, pl.productUrl, pl.seller_id, s.seller_name, pl.category_id, c.name, pl.is_used FROM product_list pl ,product p ,category c ,seller s where p.id = pl.id AND s.id = pl.seller_id AND pl.category_id = c.id LIMIT " + strconv.Itoa(firstnumber-1) + ", 12")

	}

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.ProductName, &result.Thumbnail, &result.ProductPrice, &result.ProductUrl, &result.Seller_id, &result.Seller_name, &result.Category_id, &result.Category_name, &result.Is_used)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, result)

	}

	return results
}

func FindSellerNameById(sellerId int) string {
	db, err := sql.Open("mysql", connectionString)
	var sellerName string = ""
	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	rows, err := db.Query("SELECT seller_name FROM seller WHERE id = " + strconv.Itoa(sellerId))

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&sellerName)
		if err != nil {
			log.Fatal(err)
		}
	}

	return sellerName
}

func FindProductNameById(productId int) string {
	db, err := sql.Open("mysql", connectionString)
	var productName string = ""
	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	rows, err := db.Query("SELECT name FROM product WHERE id = " + strconv.Itoa(productId))

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&productName)
		if err != nil {
			log.Fatal(err)
		}
	}

	return productName
}

func FindCategoryIdByName(categoryId int) string {
	db, err := sql.Open("mysql", connectionString)
	var categoryName string = ""
	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	rows, err := db.Query("SELECT name FROM category WHERE id = " + strconv.Itoa(categoryId))

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&categoryName)
		if err != nil {
			log.Fatal(err)
		}
	}

	return categoryName
}

func CallProductDetailSelection(productId int) model.EntProductDetail {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	var result model.EntProductDetail

	query := "SELECT * FROM product WHERE id =" + strconv.Itoa(productId)
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Name, &result.Productcode, &result.Mainimage, &result.Description, &result.Detail,
			&result.DeliveryTime, &result.ShippingFee, &result.Price, &result.Option, &result.Seller_id, &result.Eco_certification,
			&result.Status, &result.Product_url, &result.Update_date, &result.Create_date)
		if err != nil {
			log.Fatal(err)
		}

	}

	return result
}

func FindWishListById(memberno int) []model.WishListData {
	db, err := sql.Open("mysql", connectionString)

	var wishlistData model.WishListData
	var wishlistsData []model.WishListData

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	rows, err := db.Query(`SELECT 'product'           as 'type'
	, p.name              as 'name'
	, p.mainImage         as 'thumbnail'
	, p.price             as 'price'
	, p.eco_certification as 'ecoInfo'
	, p.product_url       as 'url'
 FROM wishlist_product wp
	, product p
WHERE wp.member_no =` + strconv.Itoa(memberno) +
		` AND wp.product_id = p.id
UNION ALL
SELECT 'seller'            as 'type'
	, s.seller_name       as 'name'
	, ''                  as 'thumbnail'
	, 0                  as 'price'
	, ''                  as 'ecoInfo'
	, s.seller_url        as 'url'
 FROM wishlist_seller ws
	, seller s
WHERE ws.member_no =` + strconv.Itoa(memberno) +
		` AND ws.seller_id = s.id;`)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&wishlistData.Itemtype, &wishlistData.Name, &wishlistData.Thumbnail,
			&wishlistData.Price, &wishlistData.EcoInfo, &wishlistData.Url)
		if err != nil {
			log.Fatal(err)
		}

		wishlistsData = append(wishlistsData, wishlistData)

	}

	return wishlistsData
}

func AddWishListProduct(memberNo int, productId int) int {
	db, err := sql.Open("mysql", connectionString)
	var isAdd int
	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	db.QueryRow("SELECT IF(COUNT(*)>0,FALSE ,TRUE) as 'isAdd' FROM wishlist_product wp WHERE member_no =" + strconv.Itoa(memberNo) + " AND product_id =" + strconv.Itoa(productId)).Scan(&isAdd)

	if isAdd == 1 {
		db.Exec("INSERT into wishlist_product (member_no, product_id) values (" + strconv.Itoa(memberNo) + "," + strconv.Itoa(productId) + ")")
	} else {
		return 0
	}

	return 1

}

func AddwishListSeller(memberNo int, sellerId int) int {

	db, err := sql.Open("mysql", connectionString)
	var isAdd int
	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	db.QueryRow("SELECT IF(COUNT(*)>0,FALSE ,TRUE) as 'isAdd' FROM wishlist_seller ws WHERE member_no =" + strconv.Itoa(memberNo) + " AND seller_id =" + strconv.Itoa(sellerId)).Scan(&isAdd)

	if isAdd == 1 {
		db.Exec("INSERT into wishlist_seller (member_no, seller_id) values (" + strconv.Itoa(memberNo) + ", " + strconv.Itoa(sellerId) + ")")
	} else {
		return 0
	}

	return 1

}

func DeleteAllWishListProduct(memberNo int) {

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	db.Exec("DELETE FROM wishlist_product wp WHERE member_no =" + strconv.Itoa(memberNo))

}

func DeleteWishListProduct(memberNo int, productId int) {

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	_, err = db.Exec("DELETE FROM wishlist_product wp WHERE member_no =" + strconv.Itoa(memberNo) + " AND product_id = " + strconv.Itoa(productId))

	if err != nil {
		fmt.Println(err.Error())
	}
}

func DeleteAllWishListSeller(memberNo int) {

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	db.Exec("DELETE FROM wishlist_seller WHERE member_no =" + strconv.Itoa(memberNo))

}

func DeleteWishListSeller(memberNo int, sellerId int) {

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	db.Exec("DELETE FROM wishlist_seller WHERE member_no =" + strconv.Itoa(memberNo) + " AND seller_id = " + strconv.Itoa(sellerId))

}

func CompareProductDetail(productId string) []model.ProductDetail {

	var compareproduct model.ProductDetail
	var compareproducts []model.ProductDetail

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	rows, err := db.Query("SELECT id, name, mainImage, price, product_url FROM product WHERE id IN(" + productId + ")")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&compareproduct.ProductId, &compareproduct.ProductName,
			&compareproduct.Thumbnail, &compareproduct.ProductPrice, &compareproduct.ProductUrl)
		if err != nil {
			log.Fatal(err)
		}

		compareproducts = append(compareproducts, compareproduct)

	}

	return compareproducts
}

func CallBoardList() []model.EntBoardRecommend {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	var result model.EntBoardRecommend
	var results []model.EntBoardRecommend

	rows, err := db.Query("SELECT * FROM board_recommend")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Member_no, &result.Title, &result.Contents, &result.Create_date, &result.Update_date)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, result)

	}

	return results

}

func FindMemberNameByMemberNo(memberNo int) string {
	db, err := sql.Open("mysql", connectionString)
	var memberName string = ""
	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	rows, err := db.Query("SELECT nickname FROM member WHERE member_no = " + strconv.Itoa(memberNo))

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&memberName)
		if err != nil {
			log.Fatal(err)
		}
	}

	return memberName
}

func AddRecommendPost(memberNo int, title string, contents string) {

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	db.Exec("INSERT into board_recommend (member_no, title, contents) values (" + strconv.Itoa(memberNo) + ", '" + title + "' , '" + contents + "' )")

}

func FindPostById(boardId int) model.Recommend {

	var result model.Recommend

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	rows, err := db.Query(`SELECT br.id
								, m.nickname 
								, br.title
								, br.contents 
								, br.create_date 
								, br.update_date 
 							FROM board_recommend br
								, member m 
							WHERE br.member_no = m.member_no
							  AND br.id = ` + strconv.Itoa(boardId))

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Nickname, &result.Title, &result.Contents, &result.CreateDate, &result.UpdateDate)
		if err != nil {
			log.Fatal(err)
		}
	}

	return result

}

func ModifyRecommendPost(boardId int, memberNo int, title string, contents string) {

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	db.Exec("UPDATE board_recommend br SET title = '" + title + "', contents = '" + contents + "', update_date = NOW() WHERE br.id = " + strconv.Itoa(boardId) + " AND br.member_no = " + strconv.Itoa(memberNo))

}

func DeleteRecommendPost(boardId int, memberNo int) {

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	db.Exec("DELETE FROM board_recommend br WHERE id =" + strconv.Itoa(boardId) + " AND member_no =" + strconv.Itoa(memberNo))

}

func CallReviewList() []model.EntProductReview {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	var result model.EntProductReview
	var results []model.EntProductReview

	rows, err := db.Query("SELECT * FROM product_review pr")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Member_no, &result.Product_id, &result.Contents, &result.Review_rating, &result.Create_date, &result.Update_date)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, result)

	}

	return results
}

func AddReviewPost(memberNo int, productId int, contents string, reviewRating int) {

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	query := "INSERT into product_review (member_no, product_id, contents, review_rating) values (" + strconv.Itoa(memberNo) + " , " + strconv.Itoa(productId) + " , '" + contents + "'," + strconv.Itoa(reviewRating) + ")"

	db.Exec(query)

}

func ModifyReviewPost(reviewId int, memberNo int, contents string, reviewRating int) {

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	query := "UPDATE product_review pr SET contents = '" + contents + "', review_rating = " + strconv.Itoa(reviewRating) + ", update_date = NOW() WHERE pr.id = " + strconv.Itoa(reviewId) + " AND pr.member_no = " + strconv.Itoa(memberNo)

	db.Exec(query)

}

func DeleteReviewPost(reviewId int, memberNo int) {

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	query := "DELETE FROM product_review pr WHERE id =" + strconv.Itoa(reviewId) + " AND member_no =" + strconv.Itoa(memberNo)

	db.Exec(query)

}

func SearchProductsByProductName(productName string) []model.Product {

	var result model.Product
	var results []model.Product

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	query := "SELECT pl.id, p.name, pl.thumbnail, pl.productUrl, pl.seller_id, s.seller_name, pl.category_id, c.name, pl.is_used FROM product_list pl, product p, category c, seller s where p.id = pl.id AND s.id = pl.seller_id AND pl.category_id = c.id AND p.name like '%" + productName + "%'"

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.ProductId, &result.ProductName, &result.ProductThumbnail, &result.ProductUrl, &result.Seller.SellerId, &result.Seller.SellerName, &result.CategoryId, &result.CategoryName, &result.IsUsed)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, result)
	}

	return results

}
