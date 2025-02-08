package usecase

import (
	"github.com/bllooop/monitoringapi/backend/internal/domain"
	"github.com/bllooop/monitoringapi/backend/internal/repository"
)

type MonitoringUsecase struct {
	repo repository.MonitoringService
}

func NewMonitoringUsecase(repo repository.MonitoringService) *MonitoringUsecase {
	return &MonitoringUsecase{
		repo: repo,
	}
}

func (s *MonitoringUsecase) GetData(name string) ([]domain.PingResult, error) {

	return s.repo.GetData(name)
}

func (s *MonitoringUsecase) CreateData(data domain.PingResult) (int, error) {
	return s.repo.CreateData(data)
}
