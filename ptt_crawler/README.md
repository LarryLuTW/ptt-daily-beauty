## Development

#### initialize

- go into ptt_crawler

```bash
cd ptt_crawler
```

- run develop environment in container

```bash
docker-compose up -d --build
```

- attach into container

```bash
docker exec -it ptt_crawler_functions_dev_1 bash
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
functions deploy getDailyBeauties --trigger-http
```

- testing in container

```bash
functions call getDailyBeauties --data='{"message":"Hello World"}'
```

- testing from host

```
curl http://localhost:8010/daily-beauty-209105/us-central1/getDailyBeauties
```

#### reference

- [Cloud Functions Documentation  |  Google Cloud](https://cloud.google.com/functions/docs/emulator)

## Deployment

- go into container

- deploy to cloud function

```bash
gcloud beta functions deploy getDailyBeauties --trigger-http
```

- test

```bash
curl https://us-central1-daily-beauty-209105.cloudfunctions.net/getDailyBeauties
```
