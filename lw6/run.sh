#!/bin/bash
app=./lzw-app
path=data/files/

echo "Start compressing..."
for f in ${path}*; do
    ${app} ${f} c
done
echo "End compressing!"

echo "Start decompressing..."
path=data/comp/
for f in ${path}*; do
    ${app} ${f} d
done
echo "End decompressing!"
