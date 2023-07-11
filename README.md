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

### Models (no estoy seguro del nombre, lo puedes modificar jesus)

#### Authentication

Para la autenticación se hace uso de JWT. Para ello se hace uso de dos endpoints:

##### Login

Para obtener el token de autenticación se debe hacer una petición POST al endpoint `/gophers/login` con el siguiente body:

```json
{
    "email": "your-email",
    "password": "your-password"
}
```

##### Signin

Para crear un nuevo usuario se debe hacer una petición POST al endpoint `/gophers/signin` con el siguiente body:

```json
{
    "email": "your-email",
    "password": "your-password",
    "name": "your-name"
}
```

#### Users


#### Workspaces


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



#### Messages

