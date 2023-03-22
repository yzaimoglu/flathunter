package config

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/slog"
)

// SysCallHandler handles system calls.
func SysCallHandler(signal os.Signal) {
	if !fiber.IsChild() {
		if signal == syscall.SIGTERM {
			slog.Info("Got terminate signal.")
			slog.Info("Terminating the program...")
			slog.Flush()
			os.Exit(0)
		} else if signal == syscall.SIGINT {
			slog.Info("Got interrupt signal.")
			slog.Info("Gracefully shutting down the server...")
			slog.Flush()
			os.Exit(0)
		}
	}
}

// SysCallSetup sets up the system call handler.
// It should be called after the webservers start at the very end.
// Example:
//
//	func main() {
//		...
//	 SysCallSetup()
//	}
func SysCallSetup() {
	// Wait for a signal to exit.
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)
	exitchnl := make(chan int)

	go func() {
		for {
			s := <-sigchnl
			SysCallHandler(s)
		}
	}()

	exitcode := <-exitchnl
	os.Exit(exitcode)
}
