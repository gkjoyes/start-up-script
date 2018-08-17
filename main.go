package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/g-kutty/start-up-script/jobs"
	"github.com/g-kutty/start-up-script/logger"
)

var (
	path    string
	help    bool
	version bool
)

// Version.
const (
	Version = "1.0.1"
)

// Read command line arguments before start.
func init() {
	flag.StringVar(&path, "p", "", "")
	flag.StringVar(&path, "path", "", "")
	flag.BoolVar(&help, "h", false, "")
	flag.BoolVar(&help, "help", false, "")
	flag.BoolVar(&version, "v", false, "")
	flag.BoolVar(&version, "version", false, "")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: start-up-script [options]\n")
		fmt.Fprintf(os.Stderr, "options:\n")
		fmt.Fprintf(os.Stderr, "\t-p	path 		Directory	The directory to watch.\n")
		fmt.Fprintf(os.Stderr, "\t-h	help		Help		Show this help.\n")
		fmt.Fprintf(os.Stderr, "\t-v	version		Version		Prints the version.\n")
	}
}

func main() {
	parseFlags()

	// new reader.
	r := jobs.NewReader(path)

	// read each jobs from input files.
	jobs, err := r.Read()
	if err != nil {
		logger.Error().Message(err.Error()).Log()
		os.Exit(1)
	}

	// execute jobs.
	err = jobs.Execute()
	if err != nil {
		logger.Error().Message(err.Error()).Log()
		os.Exit(1)
	}
}

// parseFlags read command line arguments.
func parseFlags() {
	flag.Parse()

	// display version.
	if version {
		fmt.Printf("start-up-script v%s\n", Version)
		os.Exit(0)
	}

	// display help guides for command line arguments.
	if help {
		flag.Usage()
		os.Exit(0)
	}

	// check path is valid.
	if path == "" {
		logger.Error().Message("Please Provide config file path", logger.FormattedMessage(path)).Log()
		os.Exit(1)
	} else if _, err := os.Stat(path); err != nil {
		logger.Error().Message("Cannot configuration file", logger.FormattedMessage(path)).Log()
		os.Exit(1)
	}
}
