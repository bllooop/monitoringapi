package usecase

import (
	"github.com/bllooop/monitoringapi/backend/internal/domain"
	"github.com/bllooop/monitoringapi/backend/internal/repository"
)

type Usecase struct {
	MonitoringService
}

type MonitoringService interface {
	GetData(name string) ([]domain.PingResult, error)
	CreateData(data domain.PingResult) (int, error)
	PingConsumer(amqpURL string)
}

func NewUsecase(repository *repository.Repository) *Usecase {
	return &Usecase{
		MonitoringService: NewMonitoringUsecase(repository),
	}
}
