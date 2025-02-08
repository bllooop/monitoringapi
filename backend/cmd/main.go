package main

import (
	running "github.com/bllooop/monitoringapi/backend/internal/server"

	_ "github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	running.Run()
}
