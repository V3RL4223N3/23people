#!/bin/bash

# To deploy this app it's as simple as running the docker run command on any gcp instance
# the app will bind to 0.0.0.0 on port 80, 
# This can be configured but I left it static since this is an example



## Created the gcp instance with this command
gcloud beta compute --project=interviews-282818 instances create instance-1 --zone=us-central1-a --machine-type=f1-micro --subnet=default --network-tier=PREMIUM --maintenance-policy=MIGRATE --service-account=832296704814-compute@developer.gserviceaccount.com --scopes=https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring.write,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/trace.append --image=debian-10-buster-v20200618 --image-project=debian-cloud --boot-disk-size=10GB --boot-disk-type=pd-standard --boot-disk-device-name=instance-1 --no-shielded-secure-boot --shielded-vtpm --shielded-integrity-monitoring --reservation-affinity=any

## Firewall, Allow only the IP from APIGEE through port 80

gcloud compute --project=interviews-282818 firewall-rules create allow-80 --direction=INGRESS --priority=1000 --network=default --action=ALLOW --rules=tcp:80 --source-ranges=35.225.242.186

 ## SSH, if need to 
gcloud beta compute ssh --zone "us-central1-a" "instance-1" --project "interviews-282818"

