package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var m = map[string]string{}

func main() {
	m["files/sample2.json"] = "https://filesamples.com/samples/code/json/sample2.json"
	m["files/sample1.json"] = "https://filesamples.com/samples/code/json/sample1.json"
	m["files/sample3.json"] = "https://filesamples.com/samples/code/json/sample3.json"
	m["files/sample4.json"] = "https://filesamples.com/samples/code/json/sample4.json"

	logFileName := "log/demo.json"
	f, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println(err)
	} else {
		log.Output(zerolog.ConsoleWriter{Out: f})
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(f).With().Timestamp().Logger()

	for key, value := range m {

		logger.Debug().Msg("Established")
		logger.Info().Msg("Creating output file")
		logger.Info().Msgf("Downloading file from: %s", value)
		out, err := os.Create(key)
		if err != nil {
			logger.Error().Err(err)
		}
		defer out.Close()

		logger.Info().Msg("Downloading File")
		resp, err := http.Get(value)
		if err != nil {
			logger.Error().Err(err)
		}
		defer resp.Body.Close()

		logger.Info().Msgf("Saving output to file: %s", key)
		_, err = io.Copy(out, resp.Body)
	}

	remove, err := exec.Command("git", "rm", "*.json").Output()
	if err != nil {
		logger.Error().Err(err)
	}
	fmt.Printf("Removed files %s\n", remove)

	add, err := exec.Command("git", "add", "*.json").Output()
	if err != nil {
		logger.Error().Err(err)
	}
	fmt.Printf("Adding %s\n", add)

	out, err := exec.Command("git", "status").Output()
	if err != nil {
		logger.Error().Err(err)
	}
	fmt.Printf("Status is %s\n", out)

	commit, err := exec.Command("git", "commit", "-m", "\"adding json files\"").Output()
	if err != nil {
		logger.Error().Err(err)
	}
	fmt.Printf("Adding %s\n", commit)

	push, err := exec.Command("git", "push", "origin", "tmp/kb").Output()
	if err != nil {
		logger.Error().Err(err)
	}
	fmt.Printf("Pushed to Github %s\n", push)

}
