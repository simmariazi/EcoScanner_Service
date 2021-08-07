package database

import (
	"database/sql"
	"fmt"
	"log"

	model "main/model"

	_ "github.com/go-sql-driver/mysql"
)

var connectionString string = "root:rladndwo3@tcp(121.166.4.186:3152)/eco_bot?charset=utf8"

func CallMemberSelection() []model.Member {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} //에러가 있으면 프로그램을 종료해라

	fmt.Println("connect success", db)

	defer db.Close()

	var result model.Member
	var results []model.Member

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
