package main

import (
    "database/sql"
	_ "github.com/lib/pq"
)

func main() {
	// We only use "database/sql" directly, but the "pq" driver is loaded in the background.
}
