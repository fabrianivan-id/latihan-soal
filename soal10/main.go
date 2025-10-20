package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Row struct {
	IBAN             string
	Amount           float64
	TransactionCount int
}

func main() {
	dsn := "user:pass@tcp(localhost:3306)/bank?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	const q = `
		SELECT c.iban,
			b.amount,
			COALESCE(t.cnt, 0) AS transaction_count
		FROM balances b
		JOIN customers c ON c.id = b.customer_id
		LEFT JOIN (
		SELECT customer_id, COUNT(*) AS cnt
		FROM transactions
		GROUP BY customer_id
		) t ON t.customer_id = b.customer_id
		WHERE b.amount < 0
		ORDER BY b.amount ASC;
	`

	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var results []Row
	for rows.Next() {
		var r Row
		if err := rows.Scan(&r.IBAN, &r.Amount, &r.TransactionCount); err != nil {
			log.Fatal(err)
		}
		results = append(results, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Print report
	fmt.Println("iban | amount | transaction_count")
	for _, r := range results {
		fmt.Printf("%s | %.2f | %d\n", r.IBAN, r.Amount, r.TransactionCount)
	}
}
