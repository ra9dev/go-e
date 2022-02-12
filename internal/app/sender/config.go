package sender

// TimeSenderConfig configures time sender
type TimeSenderConfig struct {
	ListenAddr string `short:"l" long:"listen" env:"LISTEN" description:"Listen Address (format: :8080|127.0.0.1:8080)" required:"false" default:":8080"`
}
