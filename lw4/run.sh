#!/bin/bash
app=./rsa-ecb.out
path=data/in/
for f in ${path}*; do
    ${app} ${f} $1 > /dev/null
done
