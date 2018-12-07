FROM google/cloud-sdk:alpine

WORKDIR /src

# install component "app-engine-go"
RUN gcloud components install app-engine-go -q

# setup configuration on startup
CMD gcloud auth activate-service-account --key-file deploy-en/credential.json && \
    gcloud config set project daily-beauty-209105 && \
    bash