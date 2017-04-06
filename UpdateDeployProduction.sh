#!/usr/bin/env bash

cd /opt/GoLang_work/src/github.com/BambuSolar/GoDirector

git reset HEAD --hard

git pull origin master

rm database/GoDirector.db

cd database

wget https://s3-us-west-2.amazonaws.com/bambu-energia-en-movimiento/Backups/GoDirector/GoDirector.db

cd ..
