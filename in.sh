#!/bin/bash
#tr -d  "\n[:blank:]" < config.json
echo "{\"filter\":{\"threshold\":{\"high\":42,\"low\":24, \"event\":{\"eventKey\":\"test-event-1\"}}\
    ,\"keys\":[\"test-1\"]}}"
while true; do
    echo "{\"key\":\"test-1\",\"value\":42,\"timestamp\":\"2018-10-17T07:51:07+0000\"}"
    echo "{\"key\":\"test-1\",\"value\":43,\"timestamp\":\"2018-10-17T07:51:07+0000\"}"
    sleep 1
done
