deploy:
	gcloud functions deploy availability --entry-point Availability --runtime go113 --set-env-vars GOOGLE_API_KEY=${GOOGLE_API_KEY} --trigger-http --allow-unauthenticated --region europe-west3 --memory 128MB

start:
	go run cmd/main.go

test:
	go test -v ./tests/
