package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	lokihook "github.com/akkuman/logrus-loki-hook"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

var runsCounterAuto = promauto.NewCounter(prometheus.CounterOpts{
	Name: "soa_2022_runs_auto",
})

var runsCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "soa_2022_runs",
})

func init() {
	ip, _ := os.LookupEnv("LOKI_URL")

	lokiHookConfig := &lokihook.Config{
		URL:       fmt.Sprintf("%s/api/prom/push", ip),
		LevelName: "severity",
		Labels: map[string]string{
			"service": "soa-2022",
		},
	}
	hook, err := lokihook.NewHook(lokiHookConfig)
	if err != nil {
		logger.Error(err)
	} else {
		logger.AddHook(hook)
	}

	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
}

func main() {
	i := 0

	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				logger.WithFields(logrus.Fields{
					"step":  "main",
					"time":  time.Now(),
					"count": i,
				}).Info("Executing main cycle")
				i += 1
				runsCounter.Inc()
				runsCounterAuto.Inc()

				promHost, _ := os.LookupEnv("PROMETHEUS_URL")
				err := push.New(promHost, "soa_2022").
					Collector(runsCounter).
					Push()

				fmt.Println("Pushed with error: ", err)

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
