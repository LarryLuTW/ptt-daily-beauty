FROM google/cloud-sdk:alpine

WORKDIR /src

# install component "app-engine-go"
RUN gcloud components install app-engine-go -q

# setup configuration on startup
COPY docker-entrypoint.sh /usr/local/bin/
ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["bash"]