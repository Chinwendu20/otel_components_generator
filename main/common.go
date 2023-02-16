package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/Chinwendu20/otel_components_generator/config"
	"github.com/Chinwendu20/otel_components_generator/exporters"
	"github.com/Chinwendu20/otel_components_generator/extensions"
	"github.com/Chinwendu20/otel_components_generator/processors"
	"github.com/Chinwendu20/otel_components_generator/receivers"
)

var (
	EmptyComponentErrorMessage = fmt.Sprintf("- Value for %s required, please use the component flag, --component\n", componentTypeFlag)
	EmptyModuleErrorMessage    = fmt.Sprintf("- Value for %s required, please use the module flag, --module\n", goModuleNameFlag)
	EmptyOutputErrorMessage    = fmt.Sprintf("- Value for %s required, please use the output flag, --output\n", outputDirectoryFlag)
	EmptySignalErrorMessage    = fmt.Sprintf("- Value for %s required, please use the signal flag, --signal\n", signalsFlag)
)

func ProcessOutputPath(cfg config.Struct) error {
	if _, err := os.Stat(cfg.Output); os.IsNotExist(err) {
		cfg.Logger.Info("Output path not found, creating directory")
		if err = os.Mkdir(cfg.Output, 0750); err != nil {
			return fmt.Errorf("failed to create output path: \n%w", err)
		}
	} else if err != nil {
		return fmt.Errorf(" %w", err)
	}
	cfg.Logger.Info("Processed output path")
	return nil

}

func SetGoPath(cfg *config.Struct) error {
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

func processAndWrite(cfg config.Struct, tmpl *template.Template, outFile string, tmplParams interface{}) error {
	out, err := os.Create(filepath.Clean(filepath.Join(cfg.Output, outFile)))
	if err != nil {
		return err
	}
	defer out.Close()

	return tmpl.Execute(out, tmplParams)
}

func GetModules(cfg config.Struct) error {
	if cfg.SkipGetModules {
		cfg.Logger.Info("Generating source codes only, will not update go.mod.tmpl and retrieve Go modules.")
		return nil
	}

	cfg.Logger.Info("Getting go modules")

	// #nosec G204 -- cfg.Distribution.Go is trusted to be a safe path
	cmd := exec.Command(cfg.GoPath, "mod", "tidy")
	cmd.Dir = cfg.Output
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to update go.mod.tmpl: %w. Output:\n%s", err, out)
	}

	// basic retry if error from go mod command (in case of transient network error). This could be improved
	// retry 3 times with 5 second spacing interval
	retries := 3
	failReason := "unknown"
	for i := 1; i <= retries; i++ {

		// #nosec G204 -- cfg.Distribution.Go is trusted to be a safe path
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

func obtainSourceCode(cfg config.Struct) error {
	var templates []*template.Template

	switch cfg.Component {
	case "exporter":
		templates = exporters.GenerateExporter(cfg)
	case "extension":
		templates = extensions.GenerateExtension(cfg)
	case "processor":
		templates = processors.GenerateProcessor(cfg)
	case "receiver":
		templates = receivers.GenerateReceiver(cfg)
	default:
		return errors.New("invalid value for component")

	}

	for _, tmpl := range templates {
		if err := processAndWrite(cfg, tmpl, tmpl.Name(), cfg); err != nil {
			return fmt.Errorf("failed to generate source file %q: %w", tmpl.Name(), err)
		}
	}

	cfg.Logger.Info("Sources created", zap.String("path", cfg.Output))
	return nil
}

func generateComponent(cfg config.Struct) error {
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

func validateComponent(cfg config.Struct) error {

	var errorMessage []error

	sigErr := cfg.ValidateSignal()
	compErr := cfg.ValidateComponent()
	modErr := cfg.ValidateModule()

	if sigErr == nil && compErr == nil && modErr == nil {

		return nil

	}
	errors := []error{sigErr, compErr, modErr}

	for _, error := range errors {

		if error != nil {
			errorMessage = append(errorMessage, error)
		}
	}

	return multierr.Combine(errorMessage...)
}

func checkEmptyConfigOptions(cfg config.Struct) error {
	var emptyValues []string
	if cfg.Component == "" {
		emptyValues = append(emptyValues, EmptyComponentErrorMessage)
	}
	if cfg.Module == "" {
		emptyValues = append(emptyValues, EmptyModuleErrorMessage)
	}
	if cfg.Output == "" {
		emptyValues = append(emptyValues, EmptyOutputErrorMessage)
	}
	if len(cfg.Signals) == 0 && cfg.Component != "extension" {
		emptyValues = append(emptyValues, EmptySignalErrorMessage)
	}
	if len(emptyValues) == 0 {
		return nil
	}

	emptyValues = append([]string{"\n"}, emptyValues...)
	return errors.New(strings.Join(emptyValues, ""))

}
