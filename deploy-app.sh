#!/bin/bash

docker build . -t snbsniper29/23people:prod
docker push snbsniper29/23people:prod

docker run -p 80:80 --net=host -d snbsniper29/23people:prod


#API-HEADER: apikey
#API-KEY: GU5Ldi5IvmAxihA13RV2AKQkQui9cL7d
curl --location --request GET 'https://danielftapiar-eval-prod.apigee.net/people?apikey=GU5Ldi5IvmAxihA13RV2AKQkQui9cL7d' \
--header 'Cookie: REVEL_FLASH='
