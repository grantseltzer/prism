#!/bin/bash
rm -f prism-linux-amd64.aci
acbuild begin
acbuild set-name prism
chmod +x scripts/*
acbuild copy scripts/repopulate.sh /repopulate
acbuild copy scripts/run.sh /run
acbuild copy snapOne.json /snapOne.json
acbuild copy $1 /rat
acbuild copy prism /prism
acbuild label add os linux
acbuild label add arch amd64
acbuild set-exec /run
acbuild write prism-linux-amd64.aci
acbuild end
