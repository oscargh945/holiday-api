# Holiday API

Esta es una API desarrollada en Go usando Gin Gonic para cumplir con los requisitos de un Desafio Tecnico.

La idea principal es consumir una API externa de feriados al iniciar la aplicación y guardar esa información en memoria, para no tener que llamar el servicio externo cada vez que se hace una petición.

## Tecnologías usadas

- Go
- Gin Gonic
- Docker
- OpenAPI
- slog para logs básicos
- Pruebas unitarias

## Funcionalidades

- Obtener la lista de feriados.
- Filtrar feriados por tipo.
- Filtrar feriados por rango de fechas.
- Responder en formato JSON.
- Responder en formato XML.
- Cargar los feriados una sola vez al iniciar la aplicación.
- Tener documentación OpenAPI.
- Ejecutar el proyecto con Docker.

## Arquitectura del proyecto

Se busco seguir una arquitectura tipo hexagonal, pero adaptada a una API pequeña.

La estructura principal es:

- `domain/entities`: contiene las entidades del negocio.
- `domain/repositories`: contiene las interfaces de los repositorios.
- `domain/usecase`: contiene la lógica principal de la aplicación.
- `infrastructure/client`: contiene el cliente que llama a la API externa.
- `infrastructure/repositories`: contiene la implementación del repositorio.
- `transport/http`: contiene los handlers y las rutas HTTP.
- `cmd/api`: contiene el punto de entrada de la aplicación.


## Servicio externo usado

La información de los feriados se obtiene desde este servicio:

```text
https://api.victorsanmartin.com/feriados/en.json
```


## Nota 

La aplicación intenta cargar los feriados desde la URL original del desafío:

```text
https://api.victorsanmartin.com/feriados/en.json
```

Pero el servicio respondía con error `403 Forbidden` y durante el desarrollo, esa URL redirigía a:

```text
https://api.boostr.cl/feriados/en.json
```


Entonces decidi agregar una estrategia de respaldo, la carga de feriados funciona así:

1. Intenta cargar los feriados desde la URL original.
2. Si hay redirección, intenta cargar desde la nueva URL.
3. Si todos los servicios externos fallan, se registra un warning y se cargan datos de respaldo en memoria.


## Instrucciones de uso

### Levantar la aplicación con Docker

```bash
docker compose up --build
```


## Curls de prueba

```bash
# Estado de la API
curl "http://localhost:8080/"

# Obtener todos los feriados en JSON
curl -H "Accept: application/json" "http://localhost:8080/api/v1/holidays"

# Obtener todos los feriados en XML
curl -H "Accept: application/xml" "http://localhost:8080/api/v1/holidays"

# Filtrar por tipo Civil
curl -H "Accept: application/json" "http://localhost:8080/api/v1/holidays?type=Civil"

# Filtrar por tipo Religioso
curl -H "Accept: application/json" "http://localhost:8080/api/v1/holidays?type=Religioso"

# Filtrar por rango de fechas
curl -H "Accept: application/json" "http://localhost:8080/api/v1/holidays?from=2024-01-01&to=2024-12-31"

# Filtrar por tipo y rango de fechas
curl -H "Accept: application/json" "http://localhost:8080/api/v1/holidays?type=Civil&from=2024-01-01&to=2024-12-31"

# Validar error por fecha inválida
curl -i -H "Accept: application/json" "http://localhost:8080/api/v1/holidays?from=fecha-invalida"

# Consultar OpenAPI si está servido por la app
curl "http://localhost:8080/docs/openapi.yaml"
```