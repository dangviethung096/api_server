package server

import (
	core "api_server_core"
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	pb "server_proto"
)

func NewServer() *Server {
	return &Server{}
}

type Server struct {
	pb.UnimplementedSearchFlightServiceServer
}

func (s *Server) SearchOneWay(context.Context, *pb.SearchOneWayFlightRequest) (*pb.SearchOneWayFlightResponse, error) {
	// TODO:
	// Convert to body request xml
	file, err := os.Open("server/request.xml")
	if err != nil {
		log.Fatalf("Open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var bodyString string = ""
	for scanner.Scan() {
		bodyString += scanner.Text()
	}

	if scanner.Err() != nil {
		log.Fatalf("Fail to read file")
	}

	url := "https://nodeA1.test.webservices.amadeus.com/1ASIWGENVN"
	headers := map[string]string{
		"Content-Type": "text/xml; charset=utf-8",
		"SOAPAction":   "http://webservices.amadeus.com/SATRQT_19_1_1A",
	}

	// Request to 1A http server
	res, err := core.HttpRequest(url, http.MethodPost, headers, []byte(bodyString))
	if err != nil {
		return nil, err
	}

	fmt.Printf("Response: \n%s\n", res.Body)
	// Handle body response xml

	return nil, nil
}
