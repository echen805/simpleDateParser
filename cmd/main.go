package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simpleDateParser"
	"syscall"
)

func main(){
	var (
		httpAddr = flag.String("http",":8080","http port to listen on")
	)
	flag.Parse()
	ctx := context.Background()
	// our simpleDateParser service
	srv := simpleDateParser.NewService()
	errChan := make(chan error)

	// This goroutine is to stop the server when a user presses CTRL + C
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := simpleDateParser.Endpoints{
		GetEndpoint: simpleDateParser.MakeGetEndpoint(srv),
		StatusEndpoint: simpleDateParser.MakeStatusEndpoint(srv),
		ValidateEndpoint: simpleDateParser.MakeValidateEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("simpleDateParser is listening on port: ", *httpAddr)
		handler := simpleDateParser.NewHTTPServer(ctx,endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
