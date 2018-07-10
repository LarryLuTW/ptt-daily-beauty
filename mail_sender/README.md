## Development

#### initialize

- go into project

```bash
cd mail_sender
```

- run develop environment in container

```bash
docker-compose up -d --build
```

- attach into container

```bash
docker exec -it mail_sender_app_engine_py_dev_1 bash
```

- install dependencies

```bash
pip install -r requirements.txt
```

#### develop

- start develop server

```bash
python main.py
```

- test on port 8080
```bash
curl http://127.0.0.1:8080/
```

- how to install new dependencies

```bash
pip install flask
pip freeze > requirements.txt
```

- how to get new sender credentials

```bash
rm credentials.json
python get_credentials.py --noauth_local_webserver
```

## Deployment

### app

go into container

- deploy to app engine(about 8 ~ 10 mins)

```bash
gcloud app deploy -q
```

- open with browser

```bash
gcloud app browse
```

wait for deploy finish and goto [https://daily-beauty-209105.appspot.com/](https://daily-beauty-209105.appspot.com/)

- get logs

```
gcloud app logs tail -s default
```

### cron job

```bash
gcloud app deploy cron.yaml -q
```

goto [https://console.cloud.google.com/appengine/taskqueues/cron?project=daily-beauty-209105](https://console.cloud.google.com/appengine/taskqueues/cron?project=daily-beauty-209105)

## Delete old version

- list all versions

```
gcloud app versions list
```

- stop a version
```
gcloud app versions stop <version>
```


