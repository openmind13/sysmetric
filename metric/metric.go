package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	Info = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "service_info",
	}, []string{"version"})

	RxBytesPerSecond = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "rx_bytes_per_second",
	}, []string{"interface"})

	TxBytesPerSecond = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "tx_bytes_per_second",
	}, []string{"interface"})
)
