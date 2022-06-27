#!/bin/bash

ROUTE=$1
MINMAXAVE="minmaxave"
PERMINUTE="perminute"

URL="http://localhost:8000/sensors/1/stats/"
NOW=$(date +%Y-%m-%dT%H:%M:%S)
FROM=""
TO=$(date -d "5 minutes ago" +%Y-%m-%dT%H:%M:%SZ)

if [[ "$ROUTE" == "$MINMAXAVE" ]]; then
  URL+="minmaxaverage"
  FROM=$(date -d "29 days ago" +%Y-%m-%dT%H:%M:%SZ)
elif [[ "$ROUTE" == "$PERMINUTE" ]]; then
  URL+="readings"
  FROM=$(date -d "23 hours ago" +%Y-%m-%dT%H:%M:%SZ)
fi

URL+="?from=${FROM}&to=${TO}"

echo curl $URL