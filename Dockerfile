FROM golang:1.18

RUN apt-get -qq -y update
RUN apt-get -qq -y upgrade

WORKDIR /work
RUN go install honnef.co/go/tools/cmd/staticcheck@2022.1
RUN go install -v golang.org/x/lint/golint@latest
COPY . .

RUN make && make install
RUN d6 worlds chars
