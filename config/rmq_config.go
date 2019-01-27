package config

type RmqConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Address  string `json:"address"`
	Vhost    string `json:"vhost"`
}

type ConsumerConfig struct {
	Name              string `json:"name"`
	QName             string `json:"q_name"`
	EnableRetry       bool   `json:"enable_retry"`
	HandleDeadMessage bool   `json:"handle_dead_message"`
	RetryCount        int    `json:"retry_count"`
	AutoAck           bool   `json:"auto_ack"`
	NoAck             bool   `json:"no_ack"`
	Exclusive         bool   `json:"exclusive"`
	NoWait            bool   `json:"no_wait"`
	NoLocal           bool   `json:"no_local"`
}
