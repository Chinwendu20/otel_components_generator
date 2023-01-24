package otel_components_generator

import (
	"fmt"
	"github.com/Chinwendu20/otel_components_generator/exporters"
	"github.com/Chinwendu20/otel_components_generator/internal"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"

	"go.uber.org/zap"
)

func processAndWrite(cfg internal.ConfigStruct, tmpl *template.Template, outFile string, tmplParams interface{}) error {
	out, err := os.Create(filepath.Clean(filepath.Join(cfg.Output, outFile)))
	if err != nil {
		return err
	}

	return tmpl.Execute(out, tmplParams)
}

func GetModules(cfg internal.ConfigStruct) error {
	if cfg.SkipGetModules {
		cfg.Logger.Info("Generating source codes only, will not update go.mod and retrieve Go modules.")
		return nil
	}

	cmd := exec.Command(cfg.GoPath, "mod", "tidy", "-compat=1.18")
	cmd.Dir = cfg.Output
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to update go.mod: %w. Output:\n%s", err, out)
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

func obtainSourceCode(cfg internal.ConfigStruct) error {
	var templates []*template.Template

	if cfg.Component == "exporter" {
		templates = exporters.GenerateExporter(cfg)
	}
	//if cfg.component == "extension" {
	//	templates := extensions.GenerateExtension()
	//}
	//if cfg.component == "processor" {
	//	templates := processors.GenerateProcessor()
	//}
	//if cfg.component == "receiver" {
	//	templates := receivers.GenerateReceiver()
	//}
	for _, tmpl := range templates {
		if err := processAndWrite(cfg, tmpl, tmpl.Name(), cfg); err != nil {
			return fmt.Errorf("failed to generate source file %q: %w", tmpl.Name(), err)
		}
	}

	cfg.Logger.Info("Sources created", zap.String("path", cfg.Output))
	return nil
}

func GenerateComponent(cfg internal.ConfigStruct) error {
	if err := cfg.ProcessOutputPath(); err != nil {
		return err
	}
	if err := obtainSourceCode(cfg); err != nil {
		return err
	}
	if err := cfg.SetGoPath(); err != nil {
		return err
	}
	if err := GetModules(cfg); err != nil {
		return err
	}
	return nil

}
