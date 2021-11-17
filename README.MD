## MiniCurso Microservices - SEMAC 2021

Eaí! beleza? aproveite o curso!

### Pré-requisitos

- Go 1.17.xxx
- Docker
- Kubernetes

### Steps

1. Executar código golang em: ./go-microservice

```golang
go run main.go
```

2. Executar a API:

```shell
$ curl http://localhost:8080/hello
### OU
$ watch -n 0.2 curl http://localhost:8080/hello
```

3. Conteinerizar a aplicação

```shell
docker build -t hello-from-go .
```

4. Rodando a aplicação em container

```shell
docker run -p 8080:8080 hello-from-go
```

5. Implanatando no kubernetes, na pasta kubernetes, execute;

```shell
kubectl apply -f kubernetes/deployment.yaml
```

6. Escalando a aplicação

```shell
kubectl scale deployment/hello-deployment --replicas=1
```

7. Removendo implantação

```shell
kubectl delete -f kubernetes/deployment.yaml
```