package project

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

const (
	// typeStr is the type of the processor
	typeStr = "project"
)

// NewFactory creates a processor factory
func NewFactory() processor.Factory {
	return processor.NewFactory(
		typeStr,
		createDefaultConfig,
		// Uncomment the processor type that you would like, change the second parameter as you like
		// component.StabilityLevelUndefined
		// component.StabilityLevelUnmaintained
		// component.StabilityLevelDeprecated
		// component.StabilityLevelDevelopment
		// component.StabilityLevelAlpha
		// component.StabilityLevelBeta
		// component.StabilityLevelStable
		//processor.WithMetrics(createMetricsProcessor, component.StabilityLevelBeta),
		// processor.WithTraces(createTracesProcessor, component.StabilityLevelBeta),
		// processor.WithLogs(createLogsProcessor, component.StabilityLevelAlpha),
	)
}

func createDefaultConfig() component.Config {

	return &config{}
}

// createMetricsProcessor creates a metrics processor based on this config.
func createMetricsProcessor(
	ctx context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (processor.Metrics, error) {

	return processorhelper.NewMetricsProcessor(
		ctx,
		set,
		cfg,
		nextConsumer,
		processMetrics,
		//	The parameters below are optional. Uncomment any as you need.
		//	processorhelper.WithStart(start component.StartFunc),
		//processorhelper.WithShutdown(shutdown component.ShutdownFunc),
		//processorhelper.WithCapabilities(capabilities consumer.Capabilities)
	)

}

// createTracesProcesor creates a trace processor based on this config.
func createTracesProcessor(
	ctx context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {

	return processorhelper.NewTracesProcessor(
		ctx,
		set,
		cfg,
		nextConsumer,
		processTraces,
		//	The parameters below are optional. Uncomment any as you need.
		//	processorhelper.WithStart(start component.StartFunc),
		//processorhelper.WithShutdown(shutdown component.ShutdownFunc),
		//processorhelper.WithCapabilities(capabilities consumer.Capabilities)
	)

}

func createLogsProcessor(
	ctx context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (processor.Logs, error) {

	return processorhelper.NewLogsProcessor(
		ctx,
		set,
		cfg,
		nextConsumer,
		processLogs,
		//	The parameters below are optional. Uncomment any as you need.
		//	processorhelper.WithStart(start component.StartFunc),
		//processorhelper.WithShutdown(shutdown component.ShutdownFunc),
		//processorhelper.WithCapabilities(capabilities consumer.Capabilities)
	)

}
