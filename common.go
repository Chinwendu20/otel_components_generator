package main

import (
	"errors"
	"fmt"
	"github.com/Chinwendu20/otel_components_generator/config"
	"github.com/Chinwendu20/otel_components_generator/exporters"
	"github.com/Chinwendu20/otel_components_generator/extensions"
	"github.com/Chinwendu20/otel_components_generator/processors"
	"github.com/Chinwendu20/otel_components_generator/receivers"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"

	"go.uber.org/zap"
)

func ProcessOutputPath(cfg config.ConfigStruct) error {
	if _, err := os.Stat(cfg.Output); os.IsNotExist(err) {
		cfg.Logger.Info("Output path not found, creating directory")
		if err = os.Mkdir(cfg.Output, 0750); err != nil {
			return fmt.Errorf("failed to create output path: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf(" %w", err)
	}
	cfg.Logger.Info("Output path created")
	return nil

}

func SetGoPath(cfg *config.ConfigStruct) error {
	if !cfg.SkipGetModules {
		path, err := exec.LookPath("go")
		if err != nil {
			return errors.New("GO binary not found")
		}
		cfg.GoPath = path
		cfg.Logger.Info("Using go", zap.String("go-executable", cfg.GoPath))
	}
	return nil
}

func processAndWrite(cfg config.ConfigStruct, tmpl *template.Template, outFile string, tmplParams interface{}) error {
	out, err := os.Create(filepath.Clean(filepath.Join(cfg.Output, outFile)))
	if err != nil {
		return err
	}

	return tmpl.Execute(out, tmplParams)
}

func GetModules(cfg config.ConfigStruct) error {
	if cfg.SkipGetModules {
		cfg.Logger.Info("Generating source codes only, will not update go.mod.tmpl.tmpl and retrieve Go modules.")
		return nil
	}

	cmd := exec.Command(cfg.GoPath, "mod", "tidy")
	cmd.Dir = cfg.Output
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to update go.mod.tmpl.tmpl: %w. Output:\n%s", err, out)
	}

	cfg.Logger.Info("Getting go modules")
	// basic retry if error from go mod command (in case of transient network error). This could be improved
	// retry 3 times with 5 second spacing interval
	retries := 3
	failReason := "unknown"
	for i := 1; i <= retries; i++ {
		// #nosec G204
		cmd := exec.Command(cfg.GoPath, "mod", "download")
		cmd.Dir = cfg.Output
		if out, err := cmd.CombinedOutput(); err != nil {
			failReason = fmt.Sprintf("%s. Output:\n%s", err, out)
			cfg.Logger.Info("Failed modules download", zap.String("retry", fmt.Sprintf("%d/%d", i, retries)))
			time.Sleep(5 * time.Second)
			continue
		}
		return nil
	}
	return fmt.Errorf("failed to download go modules: %s", failReason)
}

func obtainSourceCode(cfg config.ConfigStruct) error {
	var templates []*template.Template

	if cfg.Component == "exporter" {
		templates = exporters.GenerateExporter(cfg)
	}
	if cfg.Component == "extension" {
		templates = extensions.GenerateExtension(cfg)
	}
	if cfg.Component == "processor" {
		templates = processors.GenerateProcessor(cfg)
	}
	if cfg.Component == "receiver" {
		templates = receivers.GenerateReceiver(cfg)
	}
	for _, tmpl := range templates {
		if err := processAndWrite(cfg, tmpl, tmpl.Name(), cfg); err != nil {
			return fmt.Errorf("failed to generate source file %q: %w", tmpl.Name(), err)
		}
	}

	cfg.Logger.Info("Sources created", zap.String("path", cfg.Output))
	return nil
}

func generateComponent(cfg config.ConfigStruct) error {
	if err := ProcessOutputPath(cfg); err != nil {
		return err
	}
	if err := obtainSourceCode(cfg); err != nil {
		return err
	}
	if err := SetGoPath(&cfg); err != nil {
		return err
	}
	if err := GetModules(cfg); err != nil {
		return err
	}
	return nil

}

func validateComponent(cfg config.ConfigStruct) error {

	sigErr := cfg.ValidateSignal()
	compErr := cfg.ValidateComponent()

	if sigErr == nil && compErr == nil {

		return nil

	}
	if sigErr == nil {

		return compErr
	}
	if compErr == nil {

		return sigErr
	}

	return fmt.Errorf("--%w\n--%w", sigErr, compErr)
}
