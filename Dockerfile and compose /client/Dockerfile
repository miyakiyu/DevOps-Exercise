FROM ubuntu:latest

RUN apt-get update && apt-get install -y netcat

COPY start.sh /start.sh

RUN chmod +x /start.sh

CMD ["/start.sh"]
