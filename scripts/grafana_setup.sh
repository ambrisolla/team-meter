#!/bin/bash
set -e
echo
source "${PWD}/.env"

GRAFANA_USER="admin"
GRAFANA_PASS="admin"
GRAFANA_HOST="http://localhost:3000"
CURL="curl -s -u ${GRAFANA_USER}:${GRAFANA_PASS} \
  --header Accept:application/json  --header Content-Type:application/json"

checkStatus() {
  LOOP=1
  LOOPS=10
  until [ $($CURL $GRAFANA_HOST/api/health -I | grep "^HTTP" | grep "200 OK" | wc -l) == 1 ]
  do
    echo -ne " - Waitng for Grafana startup... (${LOOP} of ${LOOPS} attempts) \r"
    if [[ ${LOOP} == ${LOOPS} ]]
    then
      echo -ne "\nerror: check if Grafana is accessible!"
      exit 1
    fi
    sleep 6
    ((LOOP++))
  done
}

createPostgresDatasource() {
  PAYLOAD="{
      \"name\": \"grafana-postgresql-datasource\",
      \"type\": \"postgres\",
      \"access\": \"proxy\",
      \"url\": \"${DATABASE_HOST}:${DATABASE_PORT}\",
      \"user\": \"postgres\",
      \"password\": \"${DATABASE_PASS}\",
      \"database\": \"${DATABASE_NAME}\",
      \"jsonData\": {
      \"sslmode\": \"disable\"
      },
      \"secureJsonData\": {
        \"password\": \"${DATABASE_PASS}\"
      }}"
  MESSAGE=$($CURL -X POST $GRAFANA_HOST/api/datasources -d "$PAYLOAD" | jq -r ".message")
  echo " - Creating postgres datasource: ${MESSAGE}"
}

checkStatus
createPostgresDatasource