FROM golang:1.14.2
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
RUN go get github.com/kyawmyintthein/revel_mgo

ADD app /go/src/github.com/V3RL4223N3/23people/app
ADD conf /go/src/github.com/V3RL4223N3/23people/conf
ADD messages /go/src/github.com/V3RL4223N3/23people/messages
ADD public /go/src/github.com/V3RL4223N3/23people/public
ADD tests /go/src/github.com/V3RL4223N3/23people/tests


RUN cd /go/src/github.com/V3RL4223N3/23people
RUN revel_mgo mgo:setup

WORKDIR /go/src 
RUN revel build  github.com/V3RL4223N3/23people dev
CMD ./dev/run.sh
EXPOSE 9000
