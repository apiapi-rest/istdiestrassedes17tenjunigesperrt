package main

import (
	"context"
	"log"
	"os"

	"apiapi.rest/istdiestrassedes17tenjunigesperrt"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	ctx := context.Background()
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/availability", istdiestrassedes17tenjunigesperrt.Availability); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}

	// Use PORT environment variable, or default to 8081.
	port := "8081"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}

// func main() {
// 	data := availability.AvailabilityResponse()

// 	pretty.Println(data)

// 	json, err := json.MarshalIndent(data, "", "	")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(json))

// }
