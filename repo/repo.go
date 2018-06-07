package repo

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/goforbroke1006/watermarksvc/util/fs"
	"time"
)

const DataBaseFile = "./watermark.db"
//const DataBaseFile = "/home/goforbroke/go_projects/src/github.com/goforbroke1006/watermarksvc/watermark.db"

func CreateSchema() error {
	if !fs.IsFileExists(DataBaseFile) {
		db, err := sql.Open("sqlite3", DataBaseFile)
		if nil != err {
			return err
		}
		defer db.Close()

		_, err = db.Exec(`
CREATE TABLE files (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	created_at DATE NULL,
	filename VARCHAR(2048) NULL
);`,
		)
		return err
	}
	return nil
}

func HasFilename(db *sql.DB, filename string) bool {
	row := db.QueryRow("SELECT * FROM files WHERE filename = ?", filename)
	return nil != row
}

func AddFilename(db *sql.DB, filename string) error {
	stmt, err := db.Prepare("INSERT INTO filename(created_at, filename) values(?, ?)")
	if nil != err {
		return err
	}

	_, err = stmt.Exec(time.Now(), filename)

	return err
}
