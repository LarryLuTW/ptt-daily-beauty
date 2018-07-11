## Development

#### initialize

- go into `subscribe_api` directory

```bash
cd subscribe_api
```

- run develop environment in container

```bash
docker-compose up -d --build
```

- attach into container

```bash
docker exec -it subscribe_api_functions_dev_1 bash
```

- install dependencies from package.json

```bash
npm i
```

- start functions-emulator

```bash
functions start
```

#### run in emualator

- install new dependencies

```bash
# because npm version is 3.10.10
# use -S to save into package.json
npm i axios -S
```

- deploy to emulator

```bash
functions deploy subscribe --trigger-http
functions deploy unsubscribe --trigger-http
```

- testing in container

```bash
functions call subscribe --data='{"email":"pudding850806@gmail.com"}'
functions call unsubscribe --data='{"email":"pudding850806@gmail.com"}'
```

- testing from host

```
curl http://localhost:8010/daily-beauty-209105/us-central1/subscribe
curl http://localhost:8010/daily-beauty-209105/us-central1/unsubscribe
```

#### reference

- [Cloud Functions Documentation  |  Google Cloud](https://cloud.google.com/functions/docs/emulator)

## Deployment

- go into container

- deploy to cloud function

```bash
gcloud beta functions deploy subscribe --trigger-http
gcloud beta functions deploy unsubscribe --trigger-http
```

- test

```bash
curl https://us-central1-daily-beauty-209105.cloudfunctions.net/subscribe
curl https://us-central1-daily-beauty-209105.cloudfunctions.net/unsubscribe
```
