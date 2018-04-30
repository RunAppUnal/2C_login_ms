# 2C_login_ms

### Montar servicio a docker

Ejecutar las siguientes instucciones:

1. Subir la bd a rancher e inicializarla:

`docker-compose up`

2. Subir el programa a Rancher

`docker build -t login-ms .`

3. Iniciar el programa

`docker run --name login-ms -p 6004:6004 login-ms`

**Aclaración:** Los comandos se deben ejecutar desde la terminal, en la ruta del proyecto (1. en una terminal y 2.,3. desde otra terminal).

### Ejecutar los microservicios

1. Ejecutar microservicio login-db

`docker start login-db`

2. Ejecutar microservicio login-ms

`docker start login-ms`
