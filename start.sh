#!/bin/bash
chmod +x buildimage.sh
chmod +x scripts/*
go build
./buildimage.sh
sudo rkt --insecure-options=image run catseye-linux-amd64.aci
