package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type (
	RedisConfig struct {
		RedisMaster          string
		RedisSentinelServers []string
		RedisPoolSize        int
		RedisDB              int
		RedisPassword        string
	}
	AppConfig struct {
		RedisConfig
		DatabaseName            string
		DatabaseHost            string
		DatabasePort            int
		MgoTimeOut              int
		MgoPoolSize             int
		LogFilePath             string
		LogLevel                string
		LogRotation             bool
		LogFormatJSON           bool
		AmqpServerUrl           string
		EnableConsumer          bool
		NameSpace               string
		ActionKind              string
		ConsumerEndPoint        string
		ReportEvent             string
		SchedulerEndPoint       string
		SchedulerNextBatchEvent string
		NewRelic                map[string]interface{}
		RmqInitConfig           map[string]interface{}
		RmqConsumerConfig       map[string]interface{}
	}
)

func LoadConfig(filePath string) (a *AppConfig) {
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	var config = new(AppConfig)

	config.DatabaseName = viper.GetString("database_name")
	config.DatabaseHost = viper.GetString("database_host")
	config.DatabasePort = viper.GetInt("database_port")
	config.MgoTimeOut = viper.GetInt("mgo_timeout")
	config.MgoPoolSize = viper.GetInt("mgo_pool_size")

	config.LogFilePath = viper.GetString("log_file")
	config.LogLevel = viper.GetString("log_level")
	config.LogRotation = viper.GetBool("log_rotation")
	config.LogFormatJSON = viper.GetBool("log_format_json")

	config.RedisMaster = viper.GetString("redis_master")
	config.RedisSentinelServers = strings.Split(viper.GetString("redis_sentinels"), ",")
	config.RedisDB = viper.GetInt("redis_db")
	config.RedisPoolSize = viper.GetInt("redis_pool_size")
	config.RedisPassword = viper.GetString("redis_password")

	config.NameSpace = viper.GetString("namespace")
	config.ActionKind = viper.GetString("action_kind")

	config.AmqpServerUrl = viper.GetString("amqp_server_url")
	config.ConsumerEndPoint = viper.GetString("consumer_url")
	config.ReportEvent = viper.GetString("report_event")
	config.SchedulerEndPoint = viper.GetString("scheduler_url")
	config.SchedulerNextBatchEvent = viper.GetString("scheduler_next_batch_event")
	config.EnableConsumer = viper.GetBool("enable_consumer")
	config.NewRelic = viper.GetStringMap("new_relic")
	config.RmqInitConfig = viper.GetStringMap("rmq_config")
	config.RmqConsumerConfig = viper.GetStringMap("scheduler_consumer_config")

	return config
}
