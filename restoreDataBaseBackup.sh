#!/usr/bin/env bash

cd database

date=`date +%Y%m%d`
dateFormatted=`date -R`
s3Bucket="bambu-energia-en-movimiento"
fileName="GoDirector.db"
relativePath="/${s3Bucket}/Backups/GoDirector/${fileName}"
contentType="application/octet-stream"
stringToSign="PUT\n\n${contentType}\n${dateFormatted}\n${relativePath}"
s3AccessKey="AKIAIVRA64KOGYERRF2Q"
s3SecretKey="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"

stringToSign="GET\n\n${contentType}\n${dateFormatted}\n${relativePath}"

signature=`echo -en ${stringToSign} | openssl sha1 -hmac ${s3SecretKey} -binary | base64`

curl  -H "Host: ${s3Bucket}.s3.amazonaws.com" \
     -H "Date: ${dateFormatted}" \
     -H "Content-Type: ${contentType}" \
     -H "Authorization: AWS ${s3AccessKey}:${signature}" \
     https://${s3Bucket}.s3.amazonaws.com/Backups/GoDirector/${fileName} -o $fileName

cd ..