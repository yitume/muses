package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/yitume/muses/pkg/common"
	"github.com/yitume/muses/pkg/server/gin/plugins/ginzap"
	"time"

	"github.com/BurntSushi/toml"
)

var defaultCaller = &callerStore{
	Name: common.ModGinName,
}

type callerStore struct {
	Name         string
	IsBackground bool
	caller       *Client
	cfg          Cfg
}

type Client struct {
	*gin.Engine
}

func Register() common.Caller {
	return defaultCaller
}

func Caller() *Client {
	return defaultCaller.caller
}

func Config() Cfg {
	return defaultCaller.cfg
}

func (c *callerStore) InitCfg(cfg []byte) error {
	if err := toml.Unmarshal(cfg, &c.cfg); err != nil {
		return err
	}
	return nil
}

func (c *callerStore) InitCaller() error {
	gin.SetMode(c.cfg.Muses.Server.Gin.Mode)

	r := gin.New()
	if c.cfg.Muses.Server.Gin.EnabledLogger {
		r.Use(ginzap.Ginzap(time.RFC3339, true, c.cfg.Muses.Server.Gin.EnabledMetric))
	}

	if c.cfg.Muses.Server.Gin.EnabledRecovery {
		r.Use(ginzap.RecoveryWithZap(true))
	}

	c.caller = &Client{r}
	return nil
}
