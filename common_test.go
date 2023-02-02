package main

import (
	"github.com/Chinwendu20/otel_components_generator/config"
	"github.com/Chinwendu20/otel_components_generator/exporters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func NewTestConfig() config.ConfigStruct {
	cfg := config.NewConfig()
	cfg.Output = "temp"
	cfg.Signals = "trace"
	cfg.Module = "test"
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
	template.Execute(testOutFile, cfg)
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
