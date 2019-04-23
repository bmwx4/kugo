From  registry.xiaojukeji.com/didionline/sre-didi-centos7-init:stable
MAINTAINER frank@linux.cn
copy ./kugo /
copy ./server.key /
copy ./server.crt /
RUN chmod +x /kugo
#EXPOSE 8080
#EXPOSE 6443
CMD ["/kugo", "httpd", "--https=true"]
