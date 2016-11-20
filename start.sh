#!/bin/bash

display_usage() {
    echo -e "\nUsage:\n$0 PATH_TO_EXECUTABLE\n"
}

if [[ $# -ne 1 ]]; then
    display_usage
    exit 1
fi

if [[ $@ == "--help" || $@ == "-h" ]]; then
    display_usage
    exit 0
fi

chmod +x buildimage.sh
chmod +x scripts/*
go build
./prism snap > snapOne.json
./buildimage.sh $1
sudo rkt --insecure-options=image run prism-linux-amd64.aci
