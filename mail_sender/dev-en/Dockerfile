FROM python:3.6

WORKDIR /app

# install lsb-release
RUN apt-get update && apt-get install lsb-release -y

# install google cloud sdd
# ref: https://cloud.google.com/sdk/docs/quickstart-debian-ubuntu
RUN export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)" && \
    echo "deb http://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
    apt-get update -y && apt-get install google-cloud-sdk -y

# init gcloud
ADD daily-beauty-209105-f184a6030753.json .
RUN gcloud auth activate-service-account --key-file daily-beauty-209105-f184a6030753.json && \
    gcloud config set project daily-beauty-209105 && \
    gcloud config configurations activate default

CMD ["bash"]