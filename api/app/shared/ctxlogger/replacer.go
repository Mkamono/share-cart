package ctxlogger

import "log/slog"

func GoogleCloudReplacer(groups []string, a slog.Attr) slog.Attr {
	switch a.Key {
	case slog.LevelKey:
		if a.Key == slog.LevelWarn.String() {
			return slog.String("severity", "WARNING")
		}
		a.Key = "severity"
	case slog.MessageKey:
		a.Key = "message"
	case slog.SourceKey:
		a.Key = "logging.googleapis.com/sourceLocation"
	}
	return a
}
