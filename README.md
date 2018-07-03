# dailt-beauty

## PTT Crawler

- Nodejs or Go
- Cloud Function

### development

#### initialize

- `cd ptt_crawler`
- run develop environment in container: `docker-compose up -d --build`
- attach into container: `docker exec -it ptt_crawler_functions_dev_1 bash`
- install dependencies: `npm i`
- start functions-emulator: `functions start`

#### run in emualator

- deploy to emulator: `functions deploy helloGET --trigger-http`
- testing in container: `functions call helloGET --data='{"message":"Hello World"}'`
- testing from host: `http://localhost:8010/daily-beauty-209105/us-central1/helloGET`

#### reference

- [Cloud Functions Documentation  |  Google Cloud](https://cloud.google.com/functions/docs/emulator)

## Mail Sender

- Gmail API
- Cloud Function

## Database

- Datastore

## Subscribe API

- PubSub(optional)
- Cloud Function

## Subscribe Website

- Github Page
