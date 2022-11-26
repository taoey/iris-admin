package sysinit

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	prometheusGaugeTest = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "global_flow",
		Help: "prometheus-gauge测试",
	})
)

// 测试使用，正式环境勿用
func RecordMetrics() {
	go func() {
		count := 1
		for {
			prometheusGaugeTest.Set(float64(count))
			count += 1
			time.Sleep(time.Second)
		}
	}()
}
