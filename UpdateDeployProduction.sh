#!/usr/bin/env bash

git reset HEAD --hard

git pull origin master

cd database

date=`date +%Y%m%d`
dateFormatted=`date -R`
s3Bucket="bambu-energia-en-movimiento"
fileName="GoDirector.db"
relativePath="/${s3Bucket}/Backups/GoDirector/${fileName}"
contentType="application/octet-stream"
stringToSign="PUT\n\n${contentType}\n${dateFormatted}\n${relativePath}"
s3AccessKey="AKIAIQQ3GDZI6EA5RVPQ"
s3SecretKey="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"

signature=`echo -en ${stringToSign} | openssl sha1 -hmac ${s3SecretKey} -binary | base64`

curl -X PUT -T "${fileName}" \
-H "Host: ${s3Bucket}.s3.amazonaws.com" \
-H "Date: ${dateFormatted}" \
-H "Content-Type: ${contentType}" \
-H "Authorization: AWS ${s3AccessKey}:${signature}" \
http://${s3Bucket}.s3.amazonaws.com/Backups/GoDirector/${fileName}

cd ..
