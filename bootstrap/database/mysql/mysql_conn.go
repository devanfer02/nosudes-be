package mysql

import (
	"fmt"
	"os"
	"strings"

	"github.com/devanfer02/nosudes-be/bootstrap/env"
	"github.com/devanfer02/nosudes-be/utils/layers"
	"github.com/devanfer02/nosudes-be/utils/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	migrationDownFile = "./bootstrap/database/mysql/migrations/down/drop_tables.sql"
	migrationUpDir    = "./bootstrap/database/mysql/migrations/up"
)

func NewMysqlConn() *sqlx.DB {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		env.ProcEnv.DBUser,
		env.ProcEnv.DBPassword,
		env.ProcEnv.DBHost,
		env.ProcEnv.DBPort,
		env.ProcEnv.DBName,
	)

	db, err := sqlx.Open("mysql", dsn)

	if err != nil {
		logger.FatalLog(layers.Mysql, "failed to open database", err)
	}

	if err = db.Ping(); err != nil {
		logger.FatalLog(layers.Mysql, "could not ping to database", err)
	}

	migrateUp(db)

	return db
}

func DropAllTables(db *sqlx.DB) {
	content, err := os.ReadFile(migrationDownFile)

	if err != nil {
		logger.FatalLog(layers.Mysql, fmt.Sprintf("could not read file {%s}", migrationDownFile), err)
	}

	queries := strings.Split(string(content), ";")
	fmt.Print(queries)

	for _, query := range queries {
		_, err = db.Exec(query)

		if err != nil {
			logger.FatalLog(layers.Mysql, fmt.Sprintf("could not execute query: %s", query), err)
		}
	}
}

func migrateUp(db *sqlx.DB) {
	migrationsFile := getMigrationsUpFile()

	for _, filename := range migrationsFile {
		executeQueryInFile(db, filename)
	}
}

func getMigrationsUpFile() []string {

	files, err := os.ReadDir(migrationUpDir)

	if err != nil {
		logger.FatalLog(layers.Mysql, fmt.Sprintf("cant read directory %s", migrationUpDir), err)
	}

	migrationsFile := make([]string, 0)

	for _, file := range files {
		migrationsFile = append(
			migrationsFile,
			fmt.Sprintf("%s/%s", migrationUpDir, file.Name()),
		)
	}

	return migrationsFile
}

func executeQueryInFile(db *sqlx.DB, filename string) {

	content, err := os.ReadFile(filename)

	if err != nil {
		logger.FatalLog(layers.Mysql, fmt.Sprintf("could not read file {%s}", filename), err)
	}

	_, err = db.Exec(string(content))

	if err != nil {
		logger.FatalLog(layers.Mysql, fmt.Sprintf("could not execute sql in file {%s}", filename), err)
	}

	logger.Logger(layers.Mysql, fmt.Sprintf("migration file {%s} success", filename))
}
