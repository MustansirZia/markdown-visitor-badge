package configProvider

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Config - Configuration for the application.
type Config struct {
	// RedisHost - The host of the Redis server.
	RedisHost string
	// RedisPort - The port of the Redis server.
	RedisPort uint32
	// RedisUsername - The username of the Redis server.
	RedisUsername string
	// RedisPassword - The password of the Redis server.
	RedisPassword string
	// RedisPassword - The logical database inside the Redis server.
	RedisDatabase uint32
	// RedisUseTLS - Whether to enable TLS for redis.
	RedisUseTLS bool
	// Port - The port on which application server will listen if deployed using main.go.
	Port uint32
}

// ConfigProvider - Interface for providing configuration for the application.
type ConfigProvider interface {
	// Provide - Provides the configuration for the application.
	Provide() (Config, error)
}

// NewConfigProvider - Constructs and returns a new ConfigProvider.
func NewConfigProvider() ConfigProvider {
	return &envConfigProvider{}
}

type envConfigProvider struct{}

func (e envConfigProvider) Provide() (Config, error) {
	redisHost, err := getEnvVarOrError("REDIS_HOST")
	if err != nil {
		return Config{}, err
	}
	redisPortString, err := getEnvVarOrError("REDIS_PORT")
	if err != nil {
		return Config{}, err
	}
	redisPort, err := strconv.ParseUint(redisPortString, 10, 32)
	if err != nil {
		return Config{}, errors.New(("REDIS_PORT is not a valid number"))
	}
	portString := getEnvVarOrDefault("PORT", "8080")
	port, err := strconv.ParseUint(portString, 10, 32)
	if err != nil {
		return Config{}, errors.New(("PORT is not a valid number"))
	}
	redisUsername := getEnvVarOrDefault("REDIS_USERNAME", "")
	redisPassword := getEnvVarOrDefault("REDIS_PASSWORD", "")
	redisDatabaseString := getEnvVarOrDefault("REDIS_DATABASE", "0")
	redisDatabase, err := strconv.ParseUint(redisDatabaseString, 10, 32)
	if err != nil {
		return Config{}, errors.New(("REDIS_DATABASE is not a valid number"))
	}
	redisUseTLS := getEnvVarOrDefault("REDIS_USE_TLS", "false") == "true"
	return Config{RedisHost: redisHost, RedisPort: uint32(redisPort), Port: uint32(port), RedisUsername: redisUsername, RedisPassword: redisPassword, RedisDatabase: uint32(redisDatabase), RedisUseTLS: redisUseTLS}, nil
}

func getEnvVarOrError(key string) (string, error) {
	value, found := os.LookupEnv(key)
	if !found {
		return "", fmt.Errorf("%s key not found in environment", key)
	}
	return value, nil
}

func getEnvVarOrDefault(key string, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return defaultValue
	}
	return value
}
