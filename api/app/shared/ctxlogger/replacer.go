package ctxlogger

import "log/slog"

func GoogleCloudReplacer(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		if a.Key == slog.LevelWarn.String() {
			slog.String("severity", "WARNING")
		}
		a.Key = "severity"
	}
	if a.Key == slog.MessageKey {
		a.Key = "message"
	}
	if a.Key == slog.SourceKey {
		a.Key = "logging.googleapis.com/sourceLocation"
	}
	return a
}
