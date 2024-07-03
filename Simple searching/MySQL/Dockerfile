# MySQl Dockerfile
FROM mysql:latest

# Setting mysql
ENV MYSQL_DATABASE=db
ENV MYSQL_USER=may
ENV MYSQL_PASSWORD=12345
ENV MYSQL_ROOT_PASSWORD=12345

# Using my.cnf
COPY my.cnf /etc/my.cnf

# Set workdir
WORKDIR /var/lib/mysql

# Set port
EXPOSE 3306

CMD ["mysqld"]
