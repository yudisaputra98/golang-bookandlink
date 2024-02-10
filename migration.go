package main

import (
	"fmt"
	"github.com/yudisaputra/assignment-bookandlink/app/job/entity"
	"github.com/yudisaputra/assignment-bookandlink/core"
	"github.com/yudisaputra/assignment-bookandlink/database"
	"os"
)

func main() {
	//println(len(os.Args)) argument terminal
	if len(os.Args) > 1 {
		core.LoadEnv()
		database.Mysql()
		for i, v := range os.Args {
			if i != 0 {
				switch v {
				case "up":
					database.Instance.AutoMigrate(&entity.Job{})
					fmt.Println("Run migration successful")
				case "down":
					database.Instance.Migrator().DropTable("jobs")
					fmt.Println("Successful drop all table")
				}
			}
		}
	}
}
