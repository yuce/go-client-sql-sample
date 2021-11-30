package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/hazelcast/hazelcast-go-client/sql/driver"
)

func execSQL(db *sql.DB, text string) error {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil
	}
	lt := strings.ToLower(text)
	if strings.HasPrefix(lt, "select") || strings.HasPrefix(lt, "show") {
		return query(db, text)
	}
	return exec(db, text)
}

func query(db *sql.DB, text string) error {
	rows, err := db.Query(text)
	if err != nil {
		return fmt.Errorf("querying: %w", err)
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("retrieving columns: %w", err)
	}
	fmt.Println(strings.Join(cols, "\t"))
	fmt.Println("---")
	row := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		row[i] = new(interface{})
	}
	rowStr := make([]string, len(cols))
	for rows.Next() {
		if err := rows.Scan(row...); err != nil {
			return fmt.Errorf("scanning row: %w", err)
		}
		for i, v := range row {
			val := *(v.(*interface{}))
			rowStr[i] = fmt.Sprintf("%v", val)
		}
		fmt.Println(strings.Join(rowStr, "\t"))
	}
	return nil
}

func exec(db *sql.DB, text string) error {
	r, err := db.Exec(text)
	if err != nil {
		return fmt.Errorf("executing: %w", err)
	}
	ra, err := r.RowsAffected()
	if err != nil {
		return nil
	}
	fmt.Printf("---\nAffected rows: %d\n\n", ra)
	return nil
}

func fatal(format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	fmt.Fprintln(os.Stderr, text)
	os.Exit(1)
}

func main() {
	connStr := flag.String("c", "", "connection string")
	path := flag.String("f", "", "path to the SQL file")
	cmd := flag.String("e", "", "execute the given SQL string")
	flag.Parse()
	if *path == "" && *cmd == "" {
		flag.Usage()
		os.Exit(1)
	}
	db, err := sql.Open("hazelcast", *connStr)
	if err != nil {
		fatal("opening connection: %s", err.Error())
	}
	defer db.Close()
	if *path != "" && *cmd != "" {
		fatal("-f and -e are mutually exclusive")
	}
	var lines []string
	if *path != "" {
		b, err := ioutil.ReadFile(*path)
		if err != nil {
			fatal("reading SQL file: %s", err.Error())
		}
		text := string(b)
		lines = strings.Split(text, ";\n")
	}
	if *cmd != "" {
		lines = strings.Split(*cmd, ";")
	}
	for _, line := range lines {
		fmt.Println(">>>", line)
		if err := execSQL(db, line); err != nil {
			fatal("%s", err.Error())
		}
	}
}
