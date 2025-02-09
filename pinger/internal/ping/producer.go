package ping

import (
	"context"
	"encoding/json"
	"time"

	config "github.com/bllooop/monitoringapi/pinger/config"
	domain "github.com/bllooop/monitoringapi/pinger/internal/domain"
	logger "github.com/bllooop/monitoringapi/pkg/logging"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		logger.Log.Fatal().Err(err).Msg(msg)
		return
	}
}

func PingProducer(result domain.PingResult) {
	conn, err := connectRabbitMQ(config.AmqpURL)
	/*if err != nil {
		logger.Log.Error().Err(err).Msg("")
		logger.Log.Fatal().Msg("Failed to connect to RabbitMQ")
		return
	}*/
	failOnError(err, "Failed to connect to RabbitMQ / Не удалось подключиться к RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel / Не удалось открыть канал")

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body, err := json.Marshal(result)
	if err != nil {
		logger.Log.Error().Err(err).Msg("")
		logger.Log.Fatal().Msg("Ошибка обработки JSON")
		return
	}
	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	failOnError(err, "Не удалось опубликовать сообщение")
	logger.Log.Debug().Msgf("Отправлен результат пинга в RabbitMQ: %v", result)
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
