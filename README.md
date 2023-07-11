# TDL-gophers


## Configuration

### MySQL
Como se menciono en secciones previas, el proyecto hace uso de una base de datos de 
tipo MySQL, por lo que es necesario tener instalado el servidor de MySQL y actualizar
el archivo de configuración del proyecto con los datos de conexión a la base de datos. 
El archivo de configuracion se encuentra ubicado en `src/config/local.yaml`.

| Fields | Type | Description | Default |
| ------ | ---- | ----------- | ------- |
| host | string | MySQL host | localhost |
| port | int | MySQL port | 3306 |
| user | string | MySQL user | root |
| password | string | MySQL password | root |
| database | string | MySQL database | gochat |


## Getting Started

### Run

Para correr el proyecto se debe ejecutar el siguiente comando:

```bash
go run main.go
```

### Build

Para construir el proyecto se debe ejecutar el siguiente comando:

```bash
go build -o gophers
```

### API

#### Authentication

Es importante mencionar que para poder hacer uso de los endpoints de la API es necesario estar autenticado.
Para ello es necesario primero crear un usuario y luego hacer login para obtener el token de autenticación.

##### Signin

Para crear un nuevo usuario se debe hacer una petición POST al endpoint `/gophers/signin` con el siguiente body:

```json
{
    "email": "your-email",
    "password": "your-password",
    "name": "your-name"
}
```


##### Login

Para obtener el token de autenticación se debe hacer una petición POST al endpoint `/gophers/login` con el siguiente body:

```json
{
    "email": "your-email",
    "password": "your-password"
}
```

Si el login se realizo correctamente, se deberia obtener una respuesta con el siguiente 
formato:

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjIsImVtYWlsIjoiamVzdXNAZ21haWwuY29tIiwibmFtZSI6Inlpc3VzIiwiZXhwIjoxNjg5MTIzNDUwfQ.zkKp2GdFczRtsMGE3if1akHuNE8qB-Gga54S5kW33cE"
}
```

Este token debe ser enviado en el header de autenticación de todas las peticiones que se 
hagan a la API.

```bash
curl --location --request GET 'http://localhost:8080/gophers/example_endpoint' \
--header 'Authorization: Bearer <token>'
```

#### Workspaces

Los workspaces son los espacios de trabajo en los cuales se encuentran los canales y los usuarios.
Todo usuario debe pertenecer a al menos un workspace para poder interactuar con los canales y 
los usuarios.

##### Endpoints

| Method | Path | Description |
| ------ | ---- | ----------- |
| GET | /gophers/workspace | Obtiene todos los workspaces |
| POST | /gophers/workspace | Crea un nuevo workspace |
| PUT | /gophers/workspace/{key} | Actualiza un workspace por su key |
| DELETE | /gophers/workspace/{key} | Elimina un workspace por su key |
| POST | /gophers/workspace/{key}/join | Permite unirse a un workspace por su key |



#### Channels

Estos son los canales que se encuentran dentro de un workspace, los cuales pueden ser de dos tipos:
- Publicos: Cualquier usuario del workspace puede unirse a ellos.
- Privados: Solo los usuarios que con la contraseña pueden unirse a ellos.

##### Endpoints

| Method | Path | Description |
| ------ | ---- | ----------- |
| GET | /gophers/workspace/{workspaceKey}/channel | Obtiene todos los canales del workspace |
| POST | /gophers/workspace/{workspaceKey}/channel | Crea un nuevo canal |
| PUT | /gophers/workspace/{workspaceKey}/channel/{id} | Actualiza un canal por su id |
| DELETE | /gophers/workspace/{workspaceKey}/channel/{id} | Elimina un canal por su id |
| POST | /gophers/workspace/{workspaceKey}/channel/{id}/join | Permite unirse a un canal por su id |
| POST | /gophers/workspace/{workspaceKey}/channel/{id}/leave | Permite salirse de un canal por su id |
| GET | /gophers/workspace/{workspaceKey}/channel/{id}/members | Obtiene todos los miembros de un canal por su id |
| GET | /gophers/workspace/{workspaceKey}/channel/{id}/messages | Obtiene todos los mensajes de un canal por su id |


#### Direct Messages (DMs)

Los DMs son los mensajes directos que se pueden enviar entre dos usuarios. Para poder enviar un DM
es necesario que ambos usuarios pertenezcan al mismo workspace.

#### Endpoints

| Method | Path | Description |
| ------ | ---- | ----------- |
| GET | /gophers/dm | Obtiene todos los DMs |
| POST | /gophers/dm | Crea un nuevo DM |
| GET | /gophers/dm/{id}/messages | Obtiene todos los mensajes de un DM por su id |



#### Messages

Una vez que autenticado, el usuario puede enviar mensajes a los demas usuarios
que se encuentran en el mismo canal o DM. Para esto, es necesario iniciar una conexión websocket
con el servidor.

##### Endpoints

| Method | Path | Description |
| ------ | ---- | ----------- |
| POST | /gophers/workspace/{workspaceKey}/channel/{channelKey}/message | Crea un nuevo mensaje en un canal |
| POST | /gophers/workspace/{workspaceKey}/dm/{dmKey}/message | Crea un nuevo mensaje en un DM |

Es importante mencionar que el formato de los mensajes que se envian a traves de los endpoints
de mensajes es el siguiente:

```json
{
    "message": "your-message"
}
```

