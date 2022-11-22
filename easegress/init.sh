#!/usr/bin/env sh

CONFIG_PATH=/opt/config

for CONFIG_FILE in ${CONFIG_PATH}/*.yaml
do
	echo "create object: ${CONFIG_FILE}"
	egctl object create -f ${CONFIG_FILE} 
done