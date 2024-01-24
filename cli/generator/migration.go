package generator

import (
	"fmt"
	"os"
	"time"

	"github.com/daqing/waterway/cli/helper"
)

func GenMigration(xargs []string) {
	if len(xargs) == 0 {
		helper.Help("cli g migration [name]")
	}

	ts := time.Now().Format("20060102150405")

	fmt.Println(GenerateMigration("pg", ts, xargs[0]))
	fmt.Println(GenerateMigration("sqlite3", ts, xargs[0]))
}

func GenerateMigration(dir, ts, name string) string {
	targetPath := fmt.Sprintf("./db/%s/%s_%s.sql", dir, ts, name)
	if _, err := os.Stat(targetPath); err == nil {
		// target file exists
		fmt.Println("target file exists")
		return ""
	}

	if err := os.WriteFile(targetPath, []byte("\n"), 0644); err != nil {
		panic(err)
	}

	return targetPath
}
