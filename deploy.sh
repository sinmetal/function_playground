go mod vendor
GO111MODULE=off gcloud functions deploy SampleHandler --project metal-tile-dev2 --runtime go111 --trigger-http --region asia-northeast1
rm -r vendor