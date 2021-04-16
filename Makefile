gcloud:
	gcloud functions deploy availability --entry-point Availability --runtime go113 --set-env-vars GOOGLE_API_KEY=${GOOGLE_API_KEY} --trigger-http --allow-unauthenticated --region europe-west3 --memory 128MB

start:
	go run cmd/main.go & php -S localhost:8080 -t public

test:
	go test -v ./tests/

static:
	cd public && rsync --stats -havz --exclude='.git/' ./  ${SSH_USER}@istdiestrassedes17tenjunigesperrt.de:~/istdiestrassedes17tenjunigesperrt.de/ --delete && cd ..

magic:
	make static; make gcloud;
