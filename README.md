# KUGO
A web server image for k8s testing.

#### URL
/v1/

/healthz
...

#### Usage
```bash
# go build 
# docker build -t bmwx4/kugo .
# docker run -itd  bmwx4/kugo
# curl -k -sL -v https://172.17.0.2:6443/v1
# curl -v http://172.17.0.2:8080/v1
```
