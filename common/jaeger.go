package common

//creat Link tracking
import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewTracker(serviceName string, addr string) (opentracing.Tracer, io.Closer, error)  {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type: jaeger.SamplerTypeConst,
			Param: 1,
			//SamplingServerURL: "",
			//SamplingRefreshInterval: 0,
			//MaxOperations: 0,
			//OperationNameLateBinding: false,
			//Options: nil,
		},
		Reporter: &config.ReporterConfig{
			BufferFlushInterval: 1*time.Second,
			LogSpans: true,
			LocalAgentHostPort: addr,
		},
	}
	return cfg.NewTracer()
}