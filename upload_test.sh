#!/bin/bash

rm uploaded/*.*

echo single file upload
echo --
curl http://localhost:8080/single -F "file=@/home/amol2/Documents/code/learn/gin/files/names.txt"
echo
echo uploaded file contents:
cat uploaded/names.txt
echo

echo multiple file upload
echo --
curl http://localhost:8080/multiple -F "upload[]=@/home/amol2/Documents/code/learn/gin/files/names.txt" -F "upload[]=@/home/amol2/Documents/code/learn/gin/files/numbers.txt"
echo
echo uploaded file#1 contents:
cat uploaded/names.txt
echo
echo uploaded file#2 contents:
cat uploaded/numbers.txt
echo