package usecase

import (
	"encoding/json"
	"log"
	"time"

	"github.com/bllooop/monitoringapi/backend/internal/domain"
	"github.com/bllooop/monitoringapi/backend/internal/repository"
	logger "github.com/bllooop/monitoringapi/pkg/logging"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		logger.Log.Panic().Err(err).Msg(msg)
		return
	}
}

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

func (s *MonitoringUsecase) PingConsumer(amqpURL string) {
	conn, err := connectRabbitMQ(amqpURL)
	failOnError(err, "Не удалось подключиться к RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Не удалось открыть канал")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"PingQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Не удалось объявить очередь")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Не удалось зарегистрировать потребителя")
	var forever chan struct{}

	go func() {
		for d := range msgs {
			var result domain.PingResult
			err := json.Unmarshal(d.Body, &result)
			failOnError(err, "Ошибка обработки JSON")
			_, err = s.CreateData(result)
			failOnError(err, " Не удалось сохранить результат пинга")
			logger.Log.Debug().Msgf("Получен результат пинга из RabbitMQ: %v", result)
		}
	}()

	log.Printf(" [*] Ожидание сообщений. Для выхода нажмите CTRL+C ")
	<-forever
}

func connectRabbitMQ(url string) (*amqp.Connection, error) {
	var conn *amqp.Connection
	var err error
	maxRetries := 20
	retryInterval := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		conn, err = amqp.Dial(url)
		if err == nil {
			logger.Log.Debug().Msg("Успешно подключились к RabbitMQ")
			return conn, nil
		}
		logger.Log.Debug().Msgf("RabbitMQ не готов (попытка %d/%d), повтор подключения %v... Ошибка: %v", i+1, maxRetries, retryInterval, err)
		time.Sleep(retryInterval)
	}

	return nil, err
}
