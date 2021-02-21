#!/bin/bash

DC=co2e_etcd_1
DATA=data/co2e.ini

docker exec $DC /bin/sh -c "./bin/etcd --version"
docker exec $DC /bin/sh -c "./bin/etcdctl version"
docker exec $DC /bin/sh -c "./bin/etcdctl endpoint health"

. ./scripts/read_ini.sh
read_ini $DATA

echo "$INI__ALL_SECTIONS"
echo $INI__ALL_VARS

for section in $INI__ALL_SECTIONS; do 
  echo "Storing variables in $section:"; 
  for var in `declare | grep "^INI__"$section"__"`; do 
    set -- `echo $var | tr '=' ' '`
    full=$1
    val=$2
    cat="$(cut -d'_' -f3 <<<$full)"
    key="$(cut -d'_' -f5 <<<$full)"
    key="${cat}_${key}"
    echo "$key = $val"
    docker exec $DC /bin/sh -c "./bin/etcdctl put $key $val"
  done; 
done;