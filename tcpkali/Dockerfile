###############################
# Build the IMSModule-build image.
FROM ubuntu:16.04 as build-image
ARG APP_MODULE=metrics_agent
ENV APP_MODULE_PATH=/tmp/${APP_MODULE}

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
      build-essential software-properties-common python-software-properties \
      gcc-5 g++-5 libtool automake bison flex make autoconf

#RUN apk add --update \
#  build-base \
#  coreutils \
#  libevent-dev \
#  sparsehash \
#  libtool \
#  automake \
#  bison \
#  flex \
#  make \
#  pkgconfig \
#  openssl-dev \
#  autoconf \
#  cmake \
#  bsd-compat-headers  \
#  git \
#  python

ADD ./src /opt/src
RUN cd /opt/src &&\
   autoreconf -iv &&\
   ./configure &&\
   make &&\
   make install

FROM ubuntu:16.04
COPY ./app /app
COPY --from=build-image /opt/src/src/tcpkali /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    openssl \
    bash

#CMD ["/app/tcpkali", "-c", "10000", "-w", "8", "-T", "180", "localhost:3540"]
CMD ["tail", "-f", "/dev/null"]

