#!/bin/bash
echo "Server on"
for i in {1..3}
do 
  echo "Hello World!" | nc -l -w 10 -p 9527
done
#echo ":D" | nc -l -w 5 -p 9527 
exit

