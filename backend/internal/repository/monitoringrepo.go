package repository

import (
	"context"
	"fmt"

	"github.com/bllooop/monitoringapi/backend/internal/domain"
	logger "github.com/bllooop/monitoringapi/pkg/logging"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MonitoringPostgres struct {
	pg *pgxpool.Pool
}

func NewMonitoringPostgres(pg *pgxpool.Pool) *MonitoringPostgres {
	return &MonitoringPostgres{
		pg: pg,
	}
}

func (r *MonitoringPostgres) CreateData(data domain.PingResult) (int, error) {
	tr, err := r.pg.Begin(context.Background())
	if err != nil {
		return 0, err
	}
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (ip, timestamp, last_success) VALUES ($1,$2,$3) RETURNING id", pingTable)
	logger.Log.Debug().Str("query", createListQuery).Msg("Executing createdata query / Выполнение запроса createdata")
	row := tr.QueryRow(context.Background(), createListQuery, data.IP, data.Timestamp, data.LastSuccess)
	if err := row.Scan(&id); err != nil {
		tr.Rollback(context.Background())
		return 0, err
	}
	err = tr.Commit(context.Background())
	if err != nil {
		return 0, err
	}
	logger.Log.Debug().Int("ping_id", id).Msg("Successfully saved data / Успешно сохранены данные")
	return id, nil
}
func (r *MonitoringPostgres) GetData(name string) ([]domain.PingResult, error) {

	query := fmt.Sprintf(`SELECT ip, timestamp, last_success FROM %s`, pingTable)

	var data []domain.PingResult
	row, err := r.pg.Query(context.Background(), query)
	if err != nil {
		return data, err
	}
	defer row.Close()
	for row.Next() {
		k := domain.PingResult{}
		err := row.Scan(&k.IP, &k.Timestamp, &k.LastSuccess)
		if err != nil {
			return nil, err
		}
		data = append(data, k)
	}
	if err = row.Err(); err != nil {
		return nil, err
	}
	logger.Log.Debug().Int("data_count", len(data)).Msg("Successfully fetched pings / Успешно получены данные о пингах")
	return data, nil
}
