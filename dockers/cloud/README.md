# Deploy cloud

Compile the cloud application
```
cmd/cloud$ go build
cmd/cloud$ mv cloud ../../dockers/cloud
```

```
docker build -t cloud .
docker run -it --rm --name cloudApp -v /home/diako/projects/sigma-mono/dockers/cloud/volume:/opt/volume -p 7173:7173 cloud



```


