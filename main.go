package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var url = "https://filesamples.com/samples/code/json/sample2.json"

func main() {
	logFileName := "log/demo.json"
	f, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println(err)
	} else {
		log.Output(zerolog.ConsoleWriter{Out: f})
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("Established")
	log.Info().Msg("Creating output file")
	out, err := os.Create("output.json")
	if err != nil {
		log.Error().Err(err)
	}
	defer out.Close()

	log.Info().Msg("Downloading File")
	resp, err := http.Get(url)
	defer resp.Body.Close()

	log.Info().Msg("Saving output to file")
	_, err = io.Copy(out, resp.Body)
}
