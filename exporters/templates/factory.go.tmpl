package project

import (
	"context"
	"go.opentelemetry.io/collector/exporter/exporterhelper"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

const (
	// typeStr is the type of the exporter
	typeStr = "project"
)

// NewFactory creates a Datadog exporter factory
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		typeStr,
		createDefaultConfig,
		// Uncomment the exporter type that you would like, change the second parameter as you like. Available options
		//are listed below:
		// component.StabilityLevelUndefined
		// component.StabilityLevelUnmaintained
		// component.StabilityLevelDeprecated
		// component.StabilityLevelDevelopment
		// component.StabilityLevelAlpha
		// component.StabilityLevelBeta
		// component.StabilityLevelStable
		// exporter.WithMetrics(createMetricsExporter, component.StabilityLevelBeta),
		// exporter.WithTraces(createTracesExporter, component.StabilityLevelBeta),
		// exporter.WithLogs(createLogsExporter, component.StabilityLevelAlpha),
	)
}

func createDefaultConfig() component.Config {

	return &config{}
}

// createMetricsExporter creates a metrics exporter based on this config.
func createMetricsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Metrics, error) {

	return exporterhelper.NewMetricsExporter(ctx, set, cfg,
		pushMetrics,
		//	The parameters below are optional. Uncomment any as you need.
		//	exporterhelper.WithStart(start component.StartFunc),
		//exporterhelper.WithShutdown(shutdown component.ShutdownFunc),
		//exporterhelper.WithTimeout(timeoutSettings TimeoutSettings),
		//exporterhelper.WithRetry(retrySettings RetrySettings),
		//exporterhelper.WithQueue(queueSettings QueueSettings),
		//exporterhelper.WithCapabilities(capabilities consumer.Capabilities)
	)

}

// createTracesExporter creates a trace exporter based on this config.
func createTracesExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Traces, error) {

	return exporterhelper.NewTracesExporter(ctx, set, cfg,
		pushTraces,
		//	The parameters below are optional. Uncomment any as you need.
		//	exporterhelper.WithStart(start component.StartFunc),
		//exporterhelper.WithShutdown(shutdown component.ShutdownFunc),
		//exporterhelper.WithTimeout(timeoutSettings TimeoutSettings),
		//exporterhelper.WithRetry(retrySettings RetrySettings),
		//exporterhelper.WithQueue(queueSettings QueueSettings),
		//exporterhelper.WithCapabilities(capabilities consumer.Capabilities)
	)
}

func createLogsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Logs, error) {

	return exporterhelper.NewLogsExporter(ctx, set, cfg,
		pushLogs,
		//	The parameters below are optional. Uncomment any as you need.
		//	exporterhelper.WithStart(start component.StartFunc),
		//exporterhelper.WithShutdown(shutdown component.ShutdownFunc),
		//exporterhelper.WithTimeout(timeoutSettings TimeoutSettings),
		//exporterhelper.WithRetry(retrySettings RetrySettings),
		//exporterhelper.WithQueue(queueSettings QueueSettings),
		//exporterhelper.WithCapabilities(capabilities consumer.Capabilities)
	)

}
