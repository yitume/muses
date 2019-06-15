package stat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yitume/muses/pkg/common"
	"github.com/yitume/muses/pkg/logger"
	"github.com/zsais/go-gin-prometheus"
	"net/http"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
)

var defaultCaller = &callerStore{
	Name:         "stat",
	IsBackground: true,
}

type callerStore struct {
	Name         string
	IsNecessary  bool
	IsBackground bool
	caller       *Client
	cfg          Cfg
}

type Client struct {
	*http.Server
}

func Register() common.Caller {

	return defaultCaller
}

func (c *callerStore) InitCfg(cfg []byte) error {
	if err := toml.Unmarshal(cfg, &c.cfg); err != nil {
		return err
	}
	return nil
}

func (c *callerStore) InitCaller() error {
	serverStats := &http.Server{
		Addr:         c.cfg.Muses.Server.Stat.Addr,
		Handler:      initStat(),
		ReadTimeout:  c.cfg.Muses.Server.Stat.ReadTimeout.Duration,
		WriteTimeout: c.cfg.Muses.Server.Stat.WriteTimeout.Duration,
	}
	defer func() {
		serverStats.Close()
	}()
	c.caller = &Client{serverStats}
	if err := serverStats.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
		logger.DefaultLogger().Error("ServerApi err", zap.String("err", err.Error()))
	}

	return nil
}

func initStat() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())
	p := ginprometheus.NewPrometheus(common.MetricPrefix)
	p.Use(r)
	r.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome ServerApi Stats",
			},
		)
	})

	return r
}