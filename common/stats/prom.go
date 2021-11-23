package stats

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sync"
	"time"
)

type hash struct {
	gauges       map[string]prometheus.Gauge
	gaugesVec    map[string]*prometheus.GaugeVec
	counters     map[string]prometheus.Counter
	countersVec  map[string]*prometheus.CounterVec
	histogramVec map[string]*prometheus.HistogramVec
	//registerer *prometheus.Registry
	registerer prometheus.Registerer
	handler    http.Handler
}

type PromHashing interface {
	GetHandler() http.Handler
	IncVec(name string, types string)
	DecVec(name string, types string)
	IncHistogram(types string, requestType string, startTime time.Time)
	IncHistogramQuery(types string, requestType string, table string, caller string, startTime time.Time)
}

var once sync.Once
var v PromHashing

func Initialize() PromHashing {
	once.Do(func() {
		h := hash{
			gauges:       make(map[string]prometheus.Gauge),
			gaugesVec:    make(map[string]*prometheus.GaugeVec),
			counters:     make(map[string]prometheus.Counter),
			countersVec:  make(map[string]*prometheus.CounterVec),
			histogramVec: make(map[string]*prometheus.HistogramVec),
			registerer:   prometheus.NewRegistry(),
			//registerer: prometheus.NewRegistry(),
		}

		h.register()

		gatherer := prometheus.DefaultGatherer
		h.handler = promhttp.InstrumentMetricHandler(h.registerer, promhttp.HandlerFor(gatherer, promhttp.HandlerOpts{}))

		v = &h
	})
	return v
}

func (h *hash) GetHandler() http.Handler {
	return h.handler
}

func (h *hash) IncHistogram(types string, requestType string, startTime time.Time) {
	switch types {
	case TypeDurationTs:
		h.histogramVec[types].WithLabelValues(requestType).Observe(time.Since(startTime).Seconds())
	}
}

func (h *hash) IncHistogramQuery(types string, requestType string, table string, caller string, startTime time.Time) {
	switch types {
	case TypeDurationDbQuery:
		h.histogramVec[types].WithLabelValues(requestType, table, caller).Observe(time.Since(startTime).Seconds())
	}
}

//func (h *hash) Inc(types string) {
//	switch types {
//	case TYPE_TRAFFIC_OUTBOUND_TOTAL, TYPE_TRAFFIC_INBOUND_TOTAL,
//		TYPE_QUEUE_INBOUND_CHAT, TYPE_QUEUE_INBOUND_ROUTER, TYPE_QUEUE_INBOUND_AUTH, TYPE_QUEUE_INBOUND_META, TYPE_QUEUE_INBOUND_EMIT,
//		TYPE_QUEUE_OUTBOUND_CHAT, TYPE_QUEUE_OUTBOUND_ROUTER, TYPE_QUEUE_OUTBOUND_AUTH, TYPE_QUEUE_OUTBOUND_META, TYPE_QUEUE_OUTBOUND_EMIT:
//		h.counters[types].Inc()
//	case TYPE_SESSION_LIVE, TYPE_SESSION_WEBSOCKET:
//		h.gauges[types].Inc()
//	}
//}

func (h *hash) IncVec(name string, types string) {
	switch name {
	case TypeTraffic, TypeQueue:
		h.countersVec[name].WithLabelValues(types).Inc()
	case TypeSession, TypeQueueStatus, TypeCache:
		h.gaugesVec[name].WithLabelValues(types).Inc()
	}
}

//func (h *hash) Dec(types string) {
//	switch types {
//	case TYPE_TRAFFIC_OUTBOUND_TOTAL, TYPE_TRAFFIC_INBOUND_TOTAL:
//
//	case TYPE_SESSION_LIVE, TYPE_SESSION_WEBSOCKET:
//		h.gauges[types].Dec()
//	}
//}

func (h *hash) DecVec(name string, types string) {
	switch name {
	case TypeSession, TypeQueueStatus, TypeCache:
		h.gaugesVec[name].WithLabelValues(types).Dec()
	}
}

func (h *hash) AddGauge(name string) {
	newGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: TypeNamespaceLink,
		Name:      name,
	})
	newGauge.Set(float64(0))
	prometheus.MustRegister(newGauge)
	h.gauges[name] = newGauge
	h.registerer.MustRegister(newGauge)
}

func (h *hash) AddGaugeVec(name string, types []string) {
	newGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: TypeNamespaceLink,
		Name:      name,
	}, types)

	prometheus.MustRegister(newGauge)
	h.gaugesVec[name] = newGauge
	h.registerer.MustRegister(newGauge)
}

func (h *hash) AddCounter(name string) {
	newCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: TypeNamespaceLink,
		Name:      name,
	})

	prometheus.MustRegister(newCounter)
	h.counters[name] = newCounter
	h.registerer.MustRegister(newCounter)
}

func (h *hash) AddCounterVec(name string, types []string) {
	newCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: TypeNamespaceLink,
		Name:      name,
	}, types)
	prometheus.MustRegister(newCounter)
	h.countersVec[name] = newCounter
	h.registerer.MustRegister(newCounter)
}

func (h *hash) AddHistogramVec(name string, buckets []float64, types []string) {
	newHistogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: TypeNamespaceLink,
		Name:      name,
		Buckets:   buckets,
	}, types)

	prometheus.MustRegister(newHistogram)
	h.histogramVec[name] = newHistogram
	h.registerer.MustRegister(newHistogram)
}

//func (h *hash) AddSummary(name string, buckets []float64, types []string) {
//	//newSummary := prometheus.NewSummary(prometheus.SummaryOpts{
//	//	Namespace: TYPE_NAMESPACE_LINK,
//	//	Name:      name,
//	//})
//	//
//	//prometheus.MustRegister(newHistogram)
//	//h.histogramVec[name] = newHistogram
//	//h.registerer.MustRegister(newHistogram)
//}

func (h *hash) register() {
	//h.AddGaugeVec(TYPE_SESSION, []string{TYPE_SESSION_LIVE, TYPE_SESSION_WEBSOCKET})
	h.AddGaugeVec(TypeSession, []string{"type"})
	h.AddGaugeVec(TypeQueueStatus, []string{"type"})
	//h.AddCounterVec(TYPE_TRAFFIC, []string{TYPE_TRAFFIC_INBOUND_TOTAL, TYPE_TRAFFIC_OUTBOUND_TOTAL, TYPE_TRAFFIC_FAIL_OUT, TYPE_TRAFFIC_SUCCESS_OUT})
	h.AddCounterVec(TypeTraffic, []string{"type"})
	//h.AddCounterVec(TYPE_QUEUE, []string{TYPE_QUEUE_INBOUND_ROUTER, TYPE_QUEUE_INBOUND_EMIT, TYPE_QUEUE_INBOUND_CHAT, TYPE_QUEUE_INBOUND_AUTH, TYPE_QUEUE_INBOUND_META,
	//	TYPE_QUEUE_OUTBOUND_ROUTER, TYPE_QUEUE_OUTBOUND_EMIT, TYPE_QUEUE_OUTBOUND_CHAT, TYPE_QUEUE_OUTBOUND_AUTH, TYPE_QUEUE_OUTBOUND_META})
	h.AddCounterVec(TypeQueue, []string{"type"})
	h.AddGaugeVec(TypeCache, []string{"type"})

	h.AddHistogramVec(TypeDurationTs, []float64{0.01, 0.02, 0.04, 0.08, 0.16, 0.32, 0.5, 1, 5, 10, 20, 30}, []string{"rt"})
	h.AddHistogramVec(TypeDurationDbQuery, []float64{0.01, 0.02, 0.04, 0.08, 0.16, 0.32, 0.5, 1, 5, 10, 20, 30}, []string{"crud", "table", "caller"})
}
