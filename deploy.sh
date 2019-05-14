go mod vendor
GO111MODULE=off gcloud functions deploy HelloFunction --project sinmetal-lab --runtime go111 --trigger-http --region us-central1
rm -r vendor