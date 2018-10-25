#!/bin/bash
echo "{\"filter\":{\"threshold\":{\"high\":42,\"low\":24, \"event\":{\"eventKey\":\"test-event-1\"}},\"keys\":[\"test-1\"]}}"
while true; do
    echo "{\"datapointKey\":\"test-1\",\"dataType\":\"Integer\",\"value\":42,\"tsIso8601\":\"2018-10-17T07:51:07+0000\"}"
    echo "{\"datapointKey\":\"test-1\",\"dataType\":\"Integer\",\"value\":43,\"tsIso8601\":\"2018-10-17T07:51:07+0000\"}"
    sleep 1
done
