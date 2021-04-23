start:
	php -S localhost:8080 -t public

test:
	go test -v ./tests/...

static:
	cd public && rsync --stats -havz --exclude='.git/' ./  ${SSH_USER}@istdiestrassedes17tenjunigesperrt.de:~/istdiestrassedes17tenjunigesperrt.de/ --delete && cd ..


cloud-build:
	gcloud builds submit --tag gcr.io/istdiestrassedes17tenjunigespe/api

cloud-run:
	gcloud run deploy api --image gcr.io/istdiestrassedes17tenjunigespe/api --memory=128M --platform managed --region=europe-west1

local-run:
	docker run --env PORT=3000  -p 3000:3000 istdiestrassedes17tenjunigesperrt_api

local-build:
	docker build -t istdiestrassedes17tenjunigesperrt_api .

cloud-serviceaccount:
	gcloud iam service-accounts create serviceaccount
	gcloud projects add-iam-policy-binding istdiestrassedes17tenjunigespe --member="serviceAccount:serviceaccount@istdiestrassedes17tenjunigespe.iam.gserviceaccount.com" --role="roles/owner"
	gcloud iam service-accounts keys create .google-serviceaccount.json --iam-account=serviceaccount@istdiestrassedes17tenjunigespe.iam.gserviceaccount.com

