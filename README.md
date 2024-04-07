# Pretty Logger

This program allows you to make the output of logs to Golang more convenient visually, since color selection is applied
and the transmitted parameters are formatted in a human-readable form

> **WARNING**
> It is not recommended to use this module on production and pre-production environments, especially if data
> visualization systems such as Grafana, etc. are used, as it is intended for the convenience of local development

---

Examples

```Go
func main() {
	logger := app.InitLogger(slog.LevelDebug, os.Stdout)
	logger.Info("Test Info")
	logger.Debug("Test Debug")
	logger.Warn("Test Warning")
	logger.Error("Test", slog.String("id", "123"), slog.String("username", "test"))
}
```
In this example of logger initialization, you can see how a logger is created and parameters for its operation are
passed to it, the level of logs output and where to output them, in this example, the output occurs in stdout, and the
log level is the lowest LevelDebug
```Go
const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
)
```
Also, the second and further parameter can be passed variables for output, which will be presented as json in the log
```Go
	logger.Error("Test", slog.String("id", "123"), slog.String("username", "test"))
```
![image.png](image.png)
