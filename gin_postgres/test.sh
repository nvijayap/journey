#!/usr/bin/env bash

# test.sh

echo; curl -s localhost:8989/conn | jq

echo; curl -s localhost:8989/connect | jq

echo
