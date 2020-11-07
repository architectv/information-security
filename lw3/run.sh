#!/bin/bash
app=./des-ecb
path=data/
for f in ${path}*; do
    ${app} ${f} > /dev/null
done
