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

	rows, err := db.Query("SELECT * FROM seller")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Seller_name, &result.Seller_url, &result.Create_date)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, result)

	}

	return results
}

func CallProductSimpleSelection() []model.EntProductList {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	var result model.EntProductList
	var results []model.EntProductList

	rows, err := db.Query("SELECT * FROM product_list")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Thumnail, &result.ProductUrl, &result.Seller_id, &result.Is_used)
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

	rows, err := db.Query("SELECT * FROM product WHERE id = " + strconv.Itoa(productId))

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

func CallProductDetailSelection() []model.EntProductDetail {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	var result model.EntProductDetail
	var results []model.EntProductDetail

	rows, err := db.Query("SELECT * FROM product")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Name, &result.Productcode, &result.Mainimage, &result.Description, &result.Detail_id,
			&result.Delivery_id, &result.Price, &result.Option, &result.Seller_id, &result.Eco_certification, &result.Create_date,
			&result.Update_date, &result.Status, &result.Product_url)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, result)

	}

	return results
}

func FindDeliveryInfoByDeliveryId(deliveryId int) model.DeliveryInfo {
	db, err := sql.Open("mysql", connectionString)

	var deliveryInfo model.DeliveryInfo

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	rows, err := db.Query("SELECT * FROM deliveryinfo WHERE id = " + strconv.Itoa(deliveryId))

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&deliveryInfo)
		if err != nil {
			log.Fatal(err)
		}
	}

	return deliveryInfo
}

func FindDetailInfoByDetailId(detailId int) model.DetailInfo {
	db, err := sql.Open("mysql", connectionString)

	var detailInfo model.DetailInfo

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	rows, err := db.Query("SELECT * FROM detail WHERE id = " + strconv.Itoa(detailId))

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&detailInfo)
		if err != nil {
			log.Fatal(err)
		}
	}

	return detailInfo
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

	db.QueryRow("SELECT IF(COUNT(*)>0,FALSE ,TRUE) as 'Add' FROM wishlist_product wp WHERE member_no = 1004 AND product_id = 263;").Scan(&isAdd)

	if isAdd == 1 {
		db.Exec("INSERT into wishlist_product (member_no, product_id) values (memberNo, productId)")
	} else {
		return 0
	}

	return 1

}
