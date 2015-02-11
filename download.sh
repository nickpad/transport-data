#!/bin/bash

set -e

username=$1
password=$2
script_name=`basename $0`
location="https://tdx.transportnsw.info/download/files/greater_sydney_gtfs_static.zip"

if [ -z $1 ] || [ -z $2 ]; then
  echo "usage: $script_name USERNAME PASSWORD";
  exit 1;
fi

curl --user $username:$password $location > data.zip
