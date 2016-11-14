#!/bin/bash
chmod +x buildimage.sh
chmod +x scripts/*
go build
./prism snap > snapOne.json
./buildimage.sh $1
sudo rkt --insecure-options=image run prism-linux-amd64.aci
