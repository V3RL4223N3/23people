#!/bin/bash

docker build . -t snbsniper29/23people:prod
docker push snbsniper29/23people:prod

docker run -p 80:80 --net=host -d snbsniper29/23people:prod

