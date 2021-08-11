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

	rows, err := db.Query("SELECT * FROM seller WHERE id = " + strconv.Itoa(sellerId))

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
