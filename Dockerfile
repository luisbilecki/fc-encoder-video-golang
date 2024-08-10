FROM golang:1.22.5-alpine3.20
ENV PATH="$PATH:/bin/bash" \
    BENTO4_BIN="/opt/bento4/bin" \    
    PATH="$PATH:/opt/bento4/utils" \
    PATH="$PATH:/opt/bento4/bin"

# Install Bento 4 on Alpine
RUN apk update && apk add bento4

# Install Needed Dependencies
RUN apk add --update --no-cache ffmpeg bash make curl py-pip g++ gcc cmake unzip curl

# Install Bento4 From Source
WORKDIR /tmp/bento4

ENV BENTO4_BASE_URL="https://www.bok.net/Bento4/source/" \
    BENTO4_VERSION="1-6-0-641" \
    BENTO4_TARGET="" \
    BENTO4_PATH="/opt/bento4" \
    BENTO4_TYPE="SRC"

RUN wget -q ${BENTO4_BASE_URL}/Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip

RUN mkdir -p ${BENTO4_PATH}

RUN unzip Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip -d ${BENTO4_PATH}

RUN rm -rf Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip

RUN cd ${BENTO4_PATH} && mkdir cmakebuild 

RUN cd ${BENTO4_PATH}/cmakebuild && cmake -DCMAKE_BUILD_TYPE=Release .. && make

RUN cp -R ${BENTO4_PATH}/Source/Python/utils ${BENTO4_PATH}/utils

RUN cp -a ${BENTO4_PATH}/Source/Python/wrappers/. ${BENTO4_PATH}/bin

WORKDIR /go/src

ENTRYPOINT ["tail", "-f", "/dev/null"]