#!/bin/bash

ROUTE=$1
MINMAXAVE="minmaxave"
PERMINUTE="perminute"

METHOD="POST"
URL="http://localhost:8000/stats/"
DATA="'{\"sensor_id\": 1, "
NOW=$(date +%Y-%m-%dT%H:%M:%S)
FROM=""

if [[ "$ROUTE" == "$MINMAXAVE" ]]; then
  URL+="minmaxaverage"
  FROM=$(date -d "30 days ago" +%Y-%m-%dT%H:%M:%S)
elif [[ "$ROUTE" == "perminute" ]]; then
  URL+="readings"
  FROM=$(date -d "1 day ago" +%Y-%m-%dT%H:%M:%S)
fi

DATA+="\"from\": ${FROM}, \"to\": ${NOW}}'"

echo curl $URL -x $METHOD -d $DATA