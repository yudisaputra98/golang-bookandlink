package main

import (
	"fmt"
	"github.com/yudisaputra/assignment-bookandlink/app/job/entity"
	entity2 "github.com/yudisaputra/assignment-bookandlink/app/process/entity"
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
					database.Instance.AutoMigrate(&entity2.Process{})
					fmt.Println("Run migration successful")
				case "down":
					database.Instance.Migrator().DropTable("jobs")
					database.Instance.Migrator().DropTable("processes")
					fmt.Println("Successful drop all table")
				}
			}
		}
	}
}
