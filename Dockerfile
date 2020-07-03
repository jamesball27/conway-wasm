FROM golang:1.14 as development
# https://github.com/cortesi/modd
RUN env GO111MODULE=on go get github.com/cortesi/modd/cmd/modd
WORKDIR /home/app/
COPY . ./
EXPOSE 3000
CMD modd
