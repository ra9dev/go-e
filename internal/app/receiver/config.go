package receiver

import "time"

// TimeConsumerConfig configures client time consuming
type TimeConsumerConfig struct {
	APIAddr           string        `long:"api" env:"API_ADDR" description:"API Address (format: :8080|127.0.0.1:8080)" required:"false" default:":8080"`
	RequestsFrequency time.Duration `long:"requests-frequency" env:"REQUESTS_FREQUENCY" description:"time consumer requests frequency" required:"false" default:"1s"`
}
