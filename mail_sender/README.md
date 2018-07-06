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

## Deployment new version

go into container

- deploy to app engine

```bash
gcloud app deploy
```

- open on browser

```bash
# wait for deploy finish
gcloud app browse
```

## Delete old version

- list all instances

```
gcloud app instance list
```

- delete a instance
```
gcloud app instances delete <instance id> -s <service> -v <version>
```


