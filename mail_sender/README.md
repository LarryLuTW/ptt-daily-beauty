## Deployment

#### First, go into deploy environment
```
docker-compose up -d --build
docker exec -it mail_sender_deploy_en bash
```

### app

- deploy to app engine(about 8 ~ 10 mins)

```bash
gcloud app deploy -q
```

- open with browser

```bash
gcloud app browse
```

wait for deploy finish and goto [https://daily-beauty-209105.appspot.com/test](https://daily-beauty-209105.appspot.com/test)

- get logs

```
gcloud app logs tail -s default
```

### cron job

```bash
gcloud app deploy cron.yaml -q
```

goto [https://console.cloud.google.com/appengine/taskqueues/cron?project=daily-beauty-209105](https://console.cloud.google.com/appengine/taskqueues/cron?project=daily-beauty-209105)

