package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type User struct {
	ID   int64  `xorm:"id pk autoincr"`
	Name string `xorm:"name"`
	Age  int    `xorm:"age"`
}

// createTable テーブルを作成する
func createTable(engine xorm.Engine) {
	err := engine.CreateTables(User{})
	if err != nil {
		log.Fatalf("テーブルの生成に失敗しました。: %v", err)
	}
	fmt.Println("テーブル作成が成功しました。")
}

// insert テーブルにレコードを追加する
func insert(engine xorm.Engine) {
	user := User{
		Name: "tanaka",
		Age:  20,
	}
	_, err := engine.Table("user").Insert(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("レコードの追加が完了しました。")
}

// get 単体取得(1レコードを取得)
func get(engine xorm.Engine) {
	user := User{}
	// idフィールドは Autoincrement なので Insert済みなのであれば id = 1 のユーザが取得できるはず
	result, err := engine.ID(1).Get(&user)
	if err != nil {
		log.Fatal(err)
	}
	if !result {
		log.Fatal("ユーザーが見つかりませんでした。")
	}
	fmt.Printf("取得したレコード :%+v\n", user)
}

func main() {
	//engineを作成します。
	engine, err := xorm.NewEngine("mysql", "root:password@tcp(db:3306)/test_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	//テーブルを作成する
	createTable(*engine)
	//テーブルに追加する
	insert(*engine)
	//テーブルから取得する
	get(*engine)

	fmt.Println("うまく動きました。")
}
