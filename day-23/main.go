package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {
	log.Println("stardard logger")

	// Emit time with microseconds
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// Emit the file name and line
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// Create a custom log
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	// Output into buf
	buflog.Println("hello")
	// Output to standard output
	fmt.Print("from buflog:", buf.String())

	// Use slog (structured log)
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	myslog.Info("hello again", "key", "val", "age", 25)
	myslog.Info("hello world", "user", os.Getenv("USER"))

	logger := slog.Default()
	logger.Info("hello, world", "user", os.Getenv("USER"))
}
