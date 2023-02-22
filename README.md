# <img src="https://opentelemetry.io/img/logos/opentelemetry-logo-nav.png" alt="OpenTelemetry Icon" width="45" height=""> Otel Component generator

OpenTelemetry is a collection of tools, APIs, and SDKs used to instrument, generate, collect, and export telemetry data (metrics, logs, and traces) to help analyze your software’s performance and behavior.  The OpenTelemetry Collector is a constituent of opentelemetry and  can collect data from OpenTelemetry SDKs and other sources. This project is a component generator for the OpenTelemetry collector. The Opentelemetry collector has various components like Extensions, Receivers, Exporters, and Processors. The aim of the project generate skeleton code of the various collector components.

### Installation

You can download the binary [here](#). Place in your system's PATH environment variable. You can follow the tutorial [here](https://chlee.co/how-to-setup-environment-variables-for-windows-mac-and-linux/)

> Binary not released yet please refer to the [example usage without binary](### Example usage without binary)

```bash
ocg --component exporter --output example --signal trace,log --module github.com/user/sample
```

## Example usage:

```bash
ocg --component exporter --output example --signal trace,log --module github.com/user/sample
```

### Example usage without binary
Note: Ensure you ahave GO installed, if not, install [here](https://go.dev/dl/)
 
```bash
git clone github.com/Chinwendu20/otel_components_generator
cd main
go run . --component exporter --output example --signal trace,log --module github.com/user/sample
```
## Expected Output

![image](https://user-images.githubusercontent.com/59079323/220575176-67298b97-37a6-4e6b-b9f4-798cc42c9cbf.png)


### How it works

The component generator has different flags:

- component:
It is used to specify the type of component to be generated. The following are the accepted values: exporter,processor,extension,receiver. Only one can be specified at a time.

- output:
It is used to indicate the location in which the generated source code would live in. The input should be string value.

- signal:
It is used to indicate the signal(s) associated with a component. The accepted values are: trace,metric,log. More than one can be specified at a time but should be delimited with a comma and no whitespaces in between. e.g. --signal trace,metric.

- module:
It is used to indicate the name of the module to be generated

- gopath
Indicate the Go binary while executing Go commands. Default: go from the PATH"

- skipGetModules
Indicate if the generator should only generate code without calling 'go mod tidy' (default false)

### How to contribute

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue and let us know how we can make this project better. Don't forget to give the project a star! Thanks again!

- Fork the Project
- Create your Feature Branch (git checkout -b feature/AmazingFeature)
- Commit your Changes (git commit -m 'Add some AmazingFeature')
- Push to the Branch (git push origin feature/AmazingFeature)
- Open a Pull Request

### Acknowledgement.

Conceptualizing this was heavily inspired by the work done in [opentelemetry collector builder](https://github.com/open-telemetry/opentelemetry-collector/tree/main/cmd/builder)


**Please do not forget to give this project a star**
