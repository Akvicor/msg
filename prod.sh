#!/bin/bash

if [ -f "/data/config.toml" ]; then
  echo "config.toml file found"
else
  ./msg example -p "/data/" -c > /data/config.toml
fi

./msg migrate -c /data/config.toml
./msg server -c /data/config.toml


