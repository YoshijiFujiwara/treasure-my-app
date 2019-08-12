package main

import (
	"database/sql"
	"flag"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"log"
)

package main

import (
"database/sql"
"flag"
"fmt"
_ "github.com/go-sql-driver/mysql"
"log"
)

type Report struct {
	Id        string `json:"id"`
	Done      string `json:"done"`
	Todo      string `json:"todo"`
	CreatedAt string `json:"created_at"`
	Tags      []Tag  `json:"tags"`
}

type Tag struct {
	Id       string `json:"id"`
	ReportId string `json:"report_id"`
	Name     string `json:"name"`
}

func main() {
	// データベースの導通確認
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/treasure_app")
	if err != nil {
		panic(err)
	}
	fmt.Printf("db = %+v\n", db.Stats())

	var (
		query = flag.String("query", "",
			"select, insert, update, deleteくらいか")
		//done      = flag.String("done", "", "やったこと")
		//todo      = flag.String("todo", "", "新たな課題")
		////report_id = flag.String("report_id", "", "レポートのID")
		//tags = flag.String("tags", "", "タグ。カンマ区切りで")
	)
	flag.Parse()

	//fmt.Println(*done)
	//fmt.Println(*todo)

	switch *query {
	case "SELECT":
		// レポートの一覧
		reportRows, err := db.Query("SELECT * FROM reports")
		if err != nil {
			log.Fatal(err)
		}
		// タグの一覧
		tagRows, err := db.Query("SELECT * FROM tags")
		if err != nil {
			log.Fatal(err)
		}

		var reportResult []Report
		for reportRows.Next() {
			report := Report{}
			if err := reportRows.Scan(&report.Id, &report.Done, &report.Todo, &report.CreatedAt); err != nil {
				log.Fatal(err)
			}

			var tagResult []Tag
			for tagRows.Next() {
				tag := Tag{}
				if err := tagRows.Scan(&tag.Id, &tag.ReportId, &tag.Name); err != nil {
					log.Fatal(err)
				}
				if (tag.ReportId == report.Id) {
					tagResult = append(tagResult, tag)
				}
			}
			report.Tags = tagResult
			reportResult = append(reportResult, report)
		}
		for _, r := range reportResult {
			fmt.Println("id: ", r.Id, ", todo: ", r.Todo, ", done: ", r.Done, ", created_at: ", r.CreatedAt, ", tags: ", r.Tags)
		}

	case "INSERT":
		//ins, err := db.Prepare("INSERT INTO reports(todo, done) VALUES(?,?)")
		//if err != nil {
		//	log.Fatal(err)
		//}
		//result, err := ins.Exec(*todo, *done)
		//
		//tagList := strings.Split(*tags, ",")
		//for _, tag := range tagList {
		//	tagIns, err := db.Prepare("INSERT INTO tags(report_id, name) VALUES(?,?)")
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	tagIns.Exec(result.LastInsertId(), tag)
		//}

	case "UPDATE":

	case "DELETE":

	}
}

