package configs

import (
	"flag"
	"strings"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

var configPath = flag.String("config-path", "config/cfg.env", "path to config file (.yaml or .env file)")

type Config struct {
	HTTPServer HTTPServer `yaml:"http_server"`
	TaskStore  TaskStore  `yaml:"task_store"`
	Loglvl     string     `yaml:"log-lvl" env:"LOG_LEVEL" env-default:"info"`
	Cors       Cors       `yaml:"cors"`
}
type HTTPServer struct {
	Address  string        `yaml:"address" env:"ADDRESS" env-default:":8081"`
	Timeout  time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"5s"`
	Loglvl   string        `yaml:"log-lvl" env:"LOG_LVL" env-default:"info"`
	RpcLimit int           `yaml:"rpc-limit" env:"RPC_LIMIT" env-default:"100"`
}

type Cors struct {
	AllowAllOrigins  bool          `yaml:"allowAllOrigins" env:"ALLOW_ALL_ORIGINS" env-default:"true"`
	AllowMethods     []string      `yaml:"allowMethods" env-default:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowHeaders     []string      `yaml:"allowHeaders" env-default:"Content-Type,Accept,Authorization,Origin"`
	ExposeHeaders    []string      `yaml:"exposeHeaders" env-default:"Content-Length"`
	AllowCredentials bool          `yaml:"allowCredentials" env-default:"true"`
	MaxAge           time.Duration `yaml:"maxAge" env-default:"48h"`
}

type TaskStore struct {
	Url    string `yaml:"url" env:"URL" env-default:"mongodb://admin:admin@localhost:27017"`
	NameDB string `yaml:"name_db" env:"NAME_DB" env-default:"db_tasks"`
}

func getConfigType(configPath string) string {
	return configPath[strings.LastIndex(configPath, ".")+1:]
}
func GetConfig(configPath string) (*Config, error) {
	switch getConfigType(configPath) {
	case "yaml":
		return parseYamlCfg(configPath)
	default:
		return parseEnvCfg(configPath)
	}
}
func ParseFlags() (*Config, error) {
	flag.Parse()

	return GetConfig(*configPath)
}

func parseYamlCfg(configPath string) (*Config, error) {
	config := &Config{}
	err := cleanenv.ReadConfig(configPath, config)
	return config, err
}

func parseEnvCfg(configPath string) (*Config, error) {
	config := &Config{}
	_ = godotenv.Load(configPath)
	err := cleanenv.ReadEnv(config)
	return config, err
}
func (cfg Config) ZeroLogLevel() zerolog.Level {
	switch cfg.HTTPServer.Loglvl {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}
