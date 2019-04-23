From busybox:latest
MAINTAINER frank@linux.cn
COPY kugo /
COPY server.crt /
COPY server.key /
RUN chmod +x kugo
EXPOSE 8080
EXPOSE 6443
CMD ["kugo","https=true"]
