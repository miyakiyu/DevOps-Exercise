#Dockerfile for pull ubuntu image and install netcat
FROM ubuntu:latest

#install netcat
RUN apt-get update  && apt-get install -y netcat

#writing shell script
#copy (local file) to (container)
COPY start.sh /start.sh

RUN chmod +x /start.sh

CMD ["/start.sh"]
