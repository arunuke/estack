#!/bin/sh

# Change these 3 values as required 
PROFILE=arunt
API_KEY_FILE=api.key
API_SERVER=localhost
MAAS_URL=http://$API_SERVER/MAAS/api/2.0
sudo maas-region apikey --username=$PROFILE > $API_KEY_FILE
maas login $PROFILE $MAAS_URL - < $API_KEY_FILE
