#!/bin/sh
while getopts o:d:s: flag
do
  case "${flag}" in
    o) order=${OPTARG};;
    d) display=${OPTARG};;
    s) solver=${OPTARG};;
  esac
done

go run main.go -order=$order -show=true
