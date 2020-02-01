FROM golang:latest

#  See https://code.visualstudio.com/docs/remote/containers-advanced#_reducingg-dockerfile-build-warnings.
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
    && apt install -y --no-install-recommends apt-utils 2>&1

# Verify git, iproute2, procps and lsb-release are installed
RUN apt install -y git iproute2 procps lsb-release

# Install Go tools
RUN apt update \
    && go get -x -d github.com/stamblerre/gocode 2>&1 \
    && go build -o gocode-gomod github.com/stamblerre/gocode \
    && mv gocode-gomod $GOPATH/bin/
    
# Install other tools
ENV GO111MODULE=on
RUN go get golang.org/x/tools/gopls@latest
RUN go get github.com/mdempsky/gocode
RUN go get github.com/uudashr/gopkgs/cmd/gopkgs
RUN go get github.com/ramya-rao-a/go-outline
RUN go get github.com/acroca/go-symbols
RUN go get golang.org/x/tools/cmd/guru
RUN go get golang.org/x/tools/cmd/gorename
RUN go get github.com/go-delve/delve/cmd/dlv
RUN go get github.com/stamblerre/gocode
RUN go get github.com/rogpeppe/godef
RUN go get golang.org/x/tools/cmd/goimports
RUN go get golang.org/x/lint/golint

    # Cleanup
RUN apt-get autoremove -y
RUN apt-get clean -y
RUN rm -rf /var/lib/apt/lists/*

# Revert workaround above
ENV DEBIAN_FRONTEND=dialog

# Expose service port
EXPOSE 8001