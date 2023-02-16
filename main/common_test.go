package main

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Chinwendu20/otel_components_generator/config"

	"github.com/Chinwendu20/otel_components_generator/exporters"
)

func NewConfigFullOption() config.Struct {
	cfg := config.NewConfig()
	cfg.Component = "exporter"
	cfg.Output = "pop"
	cfg.Module = "github.com/user13/myexporter"
	cfg.Signals = "trace"

	return cfg
}

func NewConfigEmptyComponent() config.Struct {
	cfg := NewConfigFullOption()
	cfg.Component = ""

	return cfg

}
func NewConfigEmptySignal() config.Struct {
	cfg := NewConfigFullOption()
	cfg.Signals = ""

	return cfg

}

func NewConfigEmptySignalExtension() config.Struct {
	cfg := NewConfigFullOption()
	cfg.Signals = ""
	cfg.Component = "extension"

	return cfg

}

func NewConfigEmptyModule() config.Struct {
	cfg := NewConfigFullOption()
	cfg.Module = ""

	return cfg

}

func NewConfigEmptyOutput() config.Struct {
	cfg := NewConfigFullOption()
	cfg.Output = ""

	return cfg

}

func NewConfigEmptyModuleAndOutput() config.Struct {
	cfg := NewConfigFullOption()
	cfg.Module = ""
	cfg.Output = ""

	return cfg

}

func NewConfigEmptySignalAndComponent() config.Struct {
	cfg := NewConfigFullOption()
	cfg.Signals = ""
	cfg.Component = ""

	return cfg

}

func NewConfigEmptySignalAndOutputExtension() config.Struct {
	cfg := NewConfigFullOption()
	cfg.Signals = ""
	cfg.Component = "extension"
	cfg.Output = ""

	return cfg

}

func NewConfigEmptyModuleAndComponent() config.Struct {
	cfg := NewConfigFullOption()
	cfg.Module = ""
	cfg.Component = ""

	return cfg

}

func NewEmptyConfigComponentOutputModule() config.Struct {
	cfg := NewConfigFullOption()
	cfg.Component = ""
	cfg.Output = ""
	cfg.Module = ""

	return cfg

}

func NewEmptyConfig() config.Struct {

	return config.NewConfig()

}

func TestNewCommandEmptyConfigOptions(t *testing.T) {
	tests := []struct {
		name string
		cfg  config.Struct
		err  error
	}{
		{

			name: "Full config options",
			cfg:  NewConfigFullOption(),
			err:  nil,
		},
		{

			name: "Empty Component",
			cfg:  NewConfigEmptyComponent(),
			err:  errors.New("\n" + EmptyComponentErrorMessage),
		},
		{

			name: "Empty Signal",
			cfg:  NewConfigEmptySignal(),
			err:  errors.New("\n" + EmptySignalErrorMessage),
		},
		{

			name: "Empty Signal with extension component",
			cfg:  NewConfigEmptySignalExtension(),
			err:  nil,
		},
		{

			name: "Empty Module",
			cfg:  NewConfigEmptyModule(),
			err:  errors.New("\n" + EmptyModuleErrorMessage),
		},
		{

			name: "Empty Output",
			cfg:  NewConfigEmptyOutput(),
			err:  errors.New("\n" + EmptyOutputErrorMessage),
		},
		{

			name: "Empty Output and Module",
			cfg:  NewConfigEmptyModuleAndOutput(),
			err:  errors.New("\n" + EmptyModuleErrorMessage + EmptyOutputErrorMessage),
		},
		{

			name: "Empty Signal and Component",
			cfg:  NewConfigEmptySignalAndComponent(),
			err:  errors.New("\n" + EmptyComponentErrorMessage + EmptySignalErrorMessage),
		},
		{

			name: "Empty Signal and Output with Extension",
			cfg:  NewConfigEmptySignalAndOutputExtension(),
			err:  errors.New("\n" + EmptyOutputErrorMessage),
		},
		{

			name: "Empty Module and Component",
			cfg:  NewConfigEmptyModuleAndComponent(),
			err:  errors.New("\n" + EmptyComponentErrorMessage + EmptyModuleErrorMessage),
		},
		{

			name: "Empty Component, Output and Module",
			cfg:  NewEmptyConfigComponentOutputModule(),
			err:  errors.New("\n" + EmptyComponentErrorMessage + EmptyModuleErrorMessage + EmptyOutputErrorMessage),
		},
		{

			name: "No flag set",
			cfg:  NewEmptyConfig(),
			err:  errors.New("\n" + EmptyComponentErrorMessage + EmptyModuleErrorMessage + EmptyOutputErrorMessage + EmptySignalErrorMessage),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkEmptyConfigOptions(tt.cfg)
			assert.Equal(t, err, tt.err)

		})
	}

}

func NewTestConfig() config.Struct {
	cfg := config.NewConfig()
	cfg.Output = "temp"
	cfg.Signals = "trace"
	cfg.Module = "github.com/user13/myexporter"
	cfg.Component = "exporter"
	return cfg
}

func TestProcessOutputPathFileExists(t *testing.T) {
	defer os.Remove("./temp")
	err := os.Mkdir("./temp", 0750)
	require.NoError(t, err)
	cfg := config.NewConfig()
	cfg.Output = "./temp"
	err = ProcessOutputPath(cfg)
	assert.NoError(t, err)

}
func TestProcessOutputPathFileDoesNotExist(t *testing.T) {
	defer os.Remove("./temp")
	cfg := config.NewConfig()
	cfg.Output = "./temp"
	err := ProcessOutputPath(cfg)
	assert.NoError(t, err)
}
func TestSetGoPathSkipGetModulesTrue(t *testing.T) {
	cfg := config.NewConfig()
	cfg.SkipGetModules = true
	err := SetGoPath(&cfg)
	require.NoError(t, err)
	assert.Equal(t, cfg.GoPath, "")

}

func TestSetGoPathSkipGetModulesFalse(t *testing.T) {
	cfg := config.NewConfig()
	err := SetGoPath(&cfg)
	require.NoError(t, err)
	path, err := exec.LookPath("go")
	require.NoError(t, err)
	assert.Equal(t, cfg.GoPath, path)

}

func TestProcessAndWrite(t *testing.T) {
	cfg := NewTestConfig()
	err := os.Mkdir(cfg.Output, 0750)
	require.NoError(t, err)
	template := exporters.GenerateExporter(cfg)[0]
	outFile := template.Name()
	err = processAndWrite(cfg, template, outFile, cfg)
	require.NoError(t, err)
	outputFilePath := filepath.Join(cfg.Output, outFile)
	assert.FileExists(t, outputFilePath)

	testOutputFilePath := filepath.Join(cfg.Output, "test2")
	testOutFile, _ := os.Create(filepath.Clean(testOutputFilePath))
	err = template.Execute(testOutFile, cfg)
	require.NoError(t, err)
	testOutFile.Close()

	content, err := os.ReadFile(outputFilePath)
	require.NoError(t, err)
	contentOfTest, err := os.ReadFile(testOutputFilePath)
	require.NoError(t, err)

	assert.Equal(t, content, contentOfTest)

	err = os.RemoveAll(cfg.Output)
	require.NoError(t, err)

}

func TestGenerateSourceCode(t *testing.T) {
	cfg := NewTestConfig()
	assert.NoError(t, ProcessOutputPath(cfg))
	assert.NoError(t, obtainSourceCode(cfg))
	assert.NoError(t, ProcessOutputPath(cfg))
	assert.NoError(t, SetGoPath(&cfg))
	assert.NoError(t, GetModules(cfg))
	defer func() {
		os.RemoveAll(cfg.Output)

	}()

}
