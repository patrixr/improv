#!/bin/bash

set -e

NAME="improv-server"
GO_PATH="$GOPATH"
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
BIN="server"
SERVICE_FILE="$DIR/$NAME.service"

sudo apt-get update
sudo apt-get install sox

sudo sh -c "sed -e 's;%DIR%;$DIR;g' -e 's;%BIN%;$BIN;g' $SERVICE_FILE > /etc/systemd/system/$NAME.service"

sudo systemctl enable $NAME
sudo systemctl start $NAME
sudo systemctl status $NAME