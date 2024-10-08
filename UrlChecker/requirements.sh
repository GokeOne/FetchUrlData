#!/bin/bash


apt update

apt install -y golang-go
apt install -y curl
go install github.com/lc/gau/v2/cmd/gau@latest


