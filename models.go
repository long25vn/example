package main

import (
	"encoding/json"
	gocache "github.com/patrickmn/go-cache"
	"os"
)

// Config struct
type Config struct {
	Cache *gocache.Cache

	AuthURL string `json:"auth_url"`
	BlogURL string `json:"blog_url"`

	JWTKey string `json:"jwt_key"`

	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
		Address  string `json:"address"`
	} `json:"database"`

	BlogNotificationChannel string `json:"blog_notification_channel"`
	NatsURL                 string `json:"nats_url"`
	QueueGroup              string `json:"queue_group"`
	DurableID               string `json:"durable_id"`
	ClusterID               string `json:"cluster_id"`
	ClientID                string `json:"client_id"`
	ClientID2               string `json:"client_id2"`
	Channel                 string `json:"channel"`
}

// SetConfig load configuration tu json file
func SetConfig(config *Config) {
	// Đọc file config.local.json
	configFile, err := os.Open("config.local.json")
	defer configFile.Close()
	if err != nil {
		// Nếu không có file config.local.json thì đọc file config.development.json
		configFile, err = os.Open("config.default.json")
		defer configFile.Close()
		if err != nil {
			panic(err)
		}
	}

	// Parse dữ liệu JSON lưu vào struct blog
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		panic(err)
	}
}
