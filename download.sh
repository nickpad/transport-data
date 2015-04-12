#!/bin/bash

set -e

username=$1
password=$2
working_dir=`dirname $0`
script_name=`basename $0`
data_dir="$working_dir/data"
location="https://tdx.transportnsw.info/download/files/greater_sydney_gtfs_static.zip"

if [ -z $1 ] || [ -z $2 ]; then
  echo "usage: $script_name USERNAME PASSWORD";
  exit 1;
fi

mkdir -p $data_dir
curl --user $username:$password $location --continue-at - --output $data_dir/data.zip
cd $data_dir
unzip -o data.zip
rm data.zip
# Remove byte order marks which cause issues with csvkit.
dos2unix *.txt
cd $working_dir
