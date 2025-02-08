package repository

import (
	"github.com/bllooop/monitoringapi/backend/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	MonitoringService
}

type MonitoringService interface {
	GetData(name string) ([]domain.PingResult, error)
	CreateData(data domain.PingResult) (int, error)
}

func NewRepository(pg *pgxpool.Pool) *Repository {
	return &Repository{
		MonitoringService: NewMonitoringPostgres(pg),
	}
}
