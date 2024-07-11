package bootstrap

import (
	"io"
	"os"
	"path"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

var wg sync.WaitGroup

type TelegramHook struct {
}

func (t *TelegramHook) Run(
	e *zerolog.Event,
	level zerolog.Level,
	message string,
) {

	if level == zerolog.InfoLevel || level == zerolog.WarnLevel || level == zerolog.ErrorLevel {
		wg.Add(1)
		go func() {
			// Get request Uuid from context
			_ = notifyTelegram(level.String(), message)
			wg.Done()
		}()
	}

}

func notifyTelegram(title, msg string) error {
	// Send messages to telegram
	SendTelegramMessage(title, msg)

	return nil
}

// Configuration for logging
type LogOptions struct {
	// Enable console logging
	ConsoleLoggingEnabled bool

	// EncodeLogsAsJson makes the log framework log JSON
	EncodeLogsAsJson bool
	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Directory to log to to when filelogging is enabled
	Directory string
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
}

var Logger *zerolog.Logger

// Configure sets up the logging framework
//
// In production, the container logs will be collected and file logging should be disabled. However,
// during development it's nicer to see logs as text and optionally write to a file when debugging
// problems in the containerized pipeline
//
// The output log file will be located at /var/log/service-xyz/service-xyz.log and
// will be rolled according to configuration set.
func ConfigureLogger(config LogOptions) {
	var writers []io.Writer

	if config.ConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if config.FileLoggingEnabled {
		writers = append(writers, newRollingFile(config))
	}
	mw := io.MultiWriter(writers...)

	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(mw).With().Timestamp().Logger()

	logger = logger.Hook(&TelegramHook{})

	logger.Info().
		Bool("fileLogging", config.FileLoggingEnabled).
		Bool("jsonLogOutput", config.EncodeLogsAsJson).
		Str("logDirectory", config.Directory).
		Str("fileName", config.Filename).
		Int("maxSizeMB", config.MaxSize).
		Int("maxBackups", config.MaxBackups).
		Int("maxAgeInDays", config.MaxAge).
		Msg("logging configured")

	Logger = &logger
}

func newRollingFile(config LogOptions) io.Writer {
	if err := os.MkdirAll(config.Directory, 0744); err != nil {
		log.Error().Err(err).Str("path", config.Directory).Msg("can't create log directory")
		return nil
	}

	return &lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxBackups: config.MaxBackups, // files
		MaxSize:    config.MaxSize,    // megabytes
		MaxAge:     config.MaxAge,     // days
	}
}

func SendTelegramMessage(title string, message string) error {
	// Send message to telegram
	// env := NewEnv()
	// pref := tele.Settings{
	// 	Token:  env.TelegramToken,
	// 	Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	// }

	// b, err := tele.NewBot(pref)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil
	// }

	// chat := tele.Chat{ID: env.TelegramChatID}
	// _, err = b.Send(&chat, fmt.Sprintf("[%s] : %s", title, message))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	return nil
}
