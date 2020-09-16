package constant

const (
	//Port and other service settings
	Port      = "port"
	Env       = "env"
	Prod      = "prod"
	Dev       = "dev"
	ConfigDir = "config_dir"
)

//Logger
const (
	LogLevel  = "loglevel"
	LogOutput = "logoutput"
	//PanicLevel and other log levels
	PanicLevel string = "panic"
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel string = "fatal"
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel string = "error"
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel string = "warn"
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel string = "info"
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel string = "debug"
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel string = "trace"
)
