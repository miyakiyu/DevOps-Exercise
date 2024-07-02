#!/bin/bash
echo "Client on"
for i in {1..3}
do
  echo  message = $(nc server 9527)
done
#echo "I got you" | nc server 9527
exit

