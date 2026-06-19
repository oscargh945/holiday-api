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