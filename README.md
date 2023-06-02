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


## Getting Statred
Para poder interactuar con el API de GoChat es necesario iniciar una conexion 
websocket con el mismo. Es importante aclarar que toda interaccion con el API 
exije hacer uso del siguiente formato.

```json
{
    "action": <ActionType: string>,
    "model": <ModelType: string>,
    "id": <Identifier: string>,
    "parameters": <Parameters: dict>
}
```

NOTA: Ante una interaccion con un formato incorrecto, el API cerrara la conexion.

A continuacion se definen las opciones disponibles para cada campo del formato.

* **Action**: Define la accion a realizar sobre el recurso indicado. Las opciones 
  disponibles son:
    * signin: Permite registrar un nuevo usuario en el API.
    * login: Permite autenticar un usuario en el API.
    * create: Permite crear un nuevo recurso en el API.
    * update: Permite actualizar un recurso del API.
    * delete: Permite eliminar un recurso del API.
    * list: Permite obtener una lista de recursos del API.
    

* **Model**: Define el tipo de recurso sobre el cual se desea realizar la accion. 
  Las opciones disponibles son:
    * user: Permite interactuar con recursos de tipo User.
    * workspace: Permite interactuar con recursos de tipo Workspace.   


* **Id**: Define el identificador del recurso sobre el cual se desea realizar la
    accion. Este campo es opcional y solo es necesario en caso que la accion a 
    realizar sea 'update' o 'delete'.
  

* **Parameters**: Define los parametros necesarios para realizar la accion. Este
    campo es opcional y solo es necesario en caso que la accion a realizar sea 
    'signin', 'login' o 'create'.
  

### Authentication
Para interactuar con el API de GoChat es necesario iniciar una conexion websocket 
con el mismo y adicionalmente es necesario autenticarse. Para ello se debe enviar 
un mensaje con el siguiente formato.

```json
{
    "action": <signin | login>,      
    "parameters": {
        "email": <Email: string>,
        "name": <Name: string>,
        "password": <Password: string>
    }
}
```

NOTA: El parametro 'name' solo sera necesario en caso que la accion a realizar sea 'signin'.