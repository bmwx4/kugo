From  centos:latest
MAINTAINER frank@linux.cn
copy ./kugo /
copy ./server.key /
copy ./server.crt /
RUN chmod +x /kugo
#EXPOSE 8080
#EXPOSE 6443
CMD ["/kugo", "httpd", "--https=true"]
