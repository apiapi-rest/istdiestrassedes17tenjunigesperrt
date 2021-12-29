THIS_DIR := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))


start:
	php -S localhost:8080 -t public

test:
	go test -v ./...

static:
	cd public && rsync --stats -havz --exclude='.git/' ./  ${SSH_USER}@istdiestrassedes17tenjunigesperrt.de:~/istdiestrassedes17tenjunigesperrt.de/ --delete && cd ..


cloud-build:
	gcloud config set account ${GOOGLE_CLOUD_ACCOUNT}
	gcloud config set project ${GOOGLE_CLOUD_PROJECT}
	gcloud builds submit --tag gcr.io/${GOOGLE_CLOUD_PROJECT}/api

cloud-run:
	gcloud config set account ${GOOGLE_CLOUD_ACCOUNT}
	gcloud config set project ${GOOGLE_CLOUD_PROJECT}
	gcloud run deploy api --image gcr.io/${GOOGLE_CLOUD_PROJECT}/api --memory=128Mi --platform managed --region=europe-west1

# give the service account the permission to read secrets.
# https://cloud.google.com/secret-manager/docs/access-control?hl=de
# needed to run this once, before first use of secret in prod cloudrun.
cloud-serviceaccount:
	gcloud projects add-iam-policy-binding ${GOOGLE_CLOUD_PROJECT} --member=serviceAccount:797324418068-compute@developer.gserviceaccount.com  --role=roles/secretmanager.secretAccessor


local-build:
	docker build -t istdiestrassedes17tenjunigesperrt_api .

local-run:
	docker run --env PORT=3000 --env GOOGLE_APPLICATION_CREDENTIALS=/root/app/key.json -v ${GOOGLE_APPLICATION_CREDENTIALS}:/root/app/key.json:ro  -p 3000:3000 istdiestrassedes17tenjunigesperrt_api


# for local development, stored in key.json
local-serviceaccount:
	gcloud iam service-accounts create serviceaccount
	gcloud projects add-iam-policy-binding ${GOOGLE_CLOUD_PROJECT} --member="serviceAccount:serviceaccount@${GOOGLE_CLOUD_PROJECT}.iam.gserviceaccount.com" --role="roles/owner"
	gcloud iam service-accounts keys create key.json --iam-account=serviceaccount@${GOOGLE_CLOUD_PROJECT}.iam.gserviceaccount.com

make local-magic:
	make test && make local-build && make local-run

cloud-magic:
	make test && make cloud-build && make cloud-run
