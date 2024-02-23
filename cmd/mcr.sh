#!/bin/bash

if [ -z "$1" ]; then 
    echo "please provde entity name"
    exit 1
fi 

repo="repository"
svc="service"
ctr="controller"

suffr="_$repo"
suffs="_$svc"
suffc="_$ctr"

entity="$1"

touch "$repo/${entity}${suffr}.go"
touch "$svc/${entity}${suffs}.go"
touch "$ctr/${entity}${suffc}.go"