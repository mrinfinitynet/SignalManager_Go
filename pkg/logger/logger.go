package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Logger struct {
	logDir      string
	file        *os.File
	logger      *log.Logger
	currentDate string
	mu          sync.Mutex
}

var (
	defaultLogger *Logger
	once          sync.Once
)

// Initialize creates and returns the default logger instance
func Initialize(logDir string) error {
	var err error
	once.Do(func() {
		defaultLogger, err = NewLogger(logDir)
	})
	return err
}

// NewLogger creates a new logger instance with daily file rotation
func NewLogger(logDir string) (*Logger, error) {
	// Create log directory if it doesn't exist
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	l := &Logger{
		logDir: logDir,
	}

	if err := l.rotateLog(); err != nil {
		return nil, err
	}

	return l, nil
}

// rotateLog creates a new log file for the current date
func (l *Logger) rotateLog() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	currentDate := time.Now().Format("2006-01-02")

	// If date hasn't changed, no need to rotate
	if l.currentDate == currentDate && l.file != nil {
		return nil
	}

	// Close existing file if open
	if l.file != nil {
		l.file.Close()
	}

	// Create new log file
	logFileName := fmt.Sprintf("app-%s.log", currentDate)
	logFilePath := filepath.Join(l.logDir, logFileName)

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	l.file = file
	l.currentDate = currentDate

	// Create multi-writer to write to both file and stdout
	multiWriter := io.MultiWriter(os.Stdout, file)
	l.logger = log.New(multiWriter, "", log.LstdFlags)

	return nil
}

// checkAndRotate checks if the date has changed and rotates the log if needed
func (l *Logger) checkAndRotate() {
	currentDate := time.Now().Format("2006-01-02")
	if l.currentDate != currentDate {
		l.rotateLog()
	}
}

// Info logs an info level message
func (l *Logger) Info(format string, v ...interface{}) {
	l.checkAndRotate()
	message := fmt.Sprintf("[INFO] "+format, v...)
	l.logger.Println(message)
}

// Error logs an error level message
func (l *Logger) Error(format string, v ...interface{}) {
	l.checkAndRotate()
	message := fmt.Sprintf("[ERROR] "+format, v...)
	l.logger.Println(message)
}

// Warning logs a warning level message
func (l *Logger) Warning(format string, v ...interface{}) {
	l.checkAndRotate()
	message := fmt.Sprintf("[WARNING] "+format, v...)
	l.logger.Println(message)
}

// Debug logs a debug level message
func (l *Logger) Debug(format string, v ...interface{}) {
	l.checkAndRotate()
	message := fmt.Sprintf("[DEBUG] "+format, v...)
	l.logger.Println(message)
}

// Close closes the log file
func (l *Logger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

// Global logger functions using the default logger instance

// Info logs an info level message using the default logger
func Info(format string, v ...interface{}) {
	if defaultLogger != nil {
		defaultLogger.Info(format, v...)
	}
}

// Error logs an error level message using the default logger
func Error(format string, v ...interface{}) {
	if defaultLogger != nil {
		defaultLogger.Error(format, v...)
	}
}

// Warning logs a warning level message using the default logger
func Warning(format string, v ...interface{}) {
	if defaultLogger != nil {
		defaultLogger.Warning(format, v...)
	}
}

// Debug logs a debug level message using the default logger
func Debug(format string, v ...interface{}) {
	if defaultLogger != nil {
		defaultLogger.Debug(format, v...)
	}
}

// Close closes the default logger
func Close() error {
	if defaultLogger != nil {
		return defaultLogger.Close()
	}
	return nil
}
