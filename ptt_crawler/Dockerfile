FROM node:6.11.5

# install lsb-release
RUN apt-get update && apt-get install lsb-release -y

# install google cloud sdf
# ref: https://cloud.google.com/sdk/docs/quickstart-debian-ubuntu
RUN export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)" && \
    echo "deb http://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
    apt-get update -y && apt-get install google-cloud-sdk -y

WORKDIR /app

# init gcloud
ADD daily-beauty-209105-f184a6030753.json .
RUN gcloud auth activate-service-account --key-file daily-beauty-209105-f184a6030753.json && \
    gcloud config configurations activate default

# install functions-emulator
RUN npm i -g @google-cloud/functions-emulator

# init functions
RUN echo daily-beauty-209105 > projectId && \
    functions config set projectId daily-beauty-209105 < projectId && \
    functions config set bindHost 0.0.0.0

CMD ["bash"]