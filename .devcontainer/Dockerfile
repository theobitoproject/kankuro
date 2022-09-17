# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.238.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.18-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends protobuf-compiler \
    && curl -OL https://github.com/google/protobuf/releases/download/v21.1/protoc-21.1-linux-x86_64.zip \
    && unzip protoc-21.1-linux-x86_64.zip -d protoc3 \
    # Move protoc3/include/* to local "include"
    # "include" may change for every system/build
    # if not sure where it is, run the command "which protoc"
    # a route will be printed where protoc lives, example "/usr/bin/protoc"
    # "include" must live beside bin, in this case then, "/usr/include"
    && sudo mv protoc3/include/* /usr/include/

RUN mkdir /usr/include/google/api
RUN curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > /usr/include/google/api/annotations.proto     
RUN curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > /usr/include/google/api/http.proto

# Optional: change owner
RUN sudo chown -R vscode /usr/include/google

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
USER vscode
RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo@latest
