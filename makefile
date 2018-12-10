build-image:
	docker build -t pttdb_deploy_en ./deploy-en

deploy-app:
	$(MAKE) build-image
	docker run -it --rm -v ${PWD}:/src pttdb_deploy_en gcloud app deploy -q

deploy-cron:
	$(MAKE) build-image
	docker run -it --rm -v ${PWD}:/src pttdb_deploy_en gcloud app deploy cron.yaml -q
