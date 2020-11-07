#!/bin/bash
app=./des-ecb
path=data/in/
for f in ${path}*; do
    ${app} ${f} > /dev/null
done
