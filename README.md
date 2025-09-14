# Licenses Service

Este proyecto implementa un servicio de gestión de licencias utilizando Go y siguiendo la arquitectura Onion.

## 🏗️ Arquitectura

El proyecto sigue los principios de **Arquitectura Onion** que separa las responsabilidades en capas concéntricas:

```
┌─────────---------------─────────────────────┐
│   Infrastructure/Presentation/Persistence   │  
│       ┌─────────────────-------────┐        │
│       │        Application         │        │ 
│       │      ┌─────────────┐       │        │
│       │      │   Domain    │       │        │ 
│       │      └─────────────┘       │        │
│       └─────────────────────-------┘        │
└────────────────────────────---------------──┘
```

### 📁 Estructura del Proyecto

```
Licenses-Service/
├── cmd/                          # Punto de entrada de la aplicación
│   └── main.go                   # Función main
├── internal/                     # Código privado de la aplicación
│   ├── domain/                   # Capa de Dominio (Core)
│   │   ├── entities/             # Entidades de negocio
|   |   |-- value_object/         # Valores objetos del dominio
│   │   ├── repositories/         # Interfaces de repositorios
│   ├── application/              # Capa de Aplicación
│   │   ├── usecases/             # Casos de uso
│   │   └── services/             # Servicios de aplicación
│   ├── infrastructure/           # Capa de Infraestructura
│   │   ├── database/             # Implementaciones de BD
│   │   ├── repositories/         # Implementaciones de repositorios
│   │   └── config/               # Configuración
│   └── presentation/             # Capa de Presentación
│       ├── handlers/             # HTTP handlers
│       └── routes/               # Definición de rutas
├── pkg/                          # Código reutilizable (público)
├── go.mod                        # Dependencias del módulo
├── go.sum                        # Checksums de dependencias
├── Dockerfile                    # Imagen Docker
└── README.md                     # Este archivo
```

## 📋 Prerrequisitos

- Go 1.21 o superior
- Docker y Docker Compose (opcional)
- PostgreSQL (si no usas Docker)

## 🛠️ Instalación y Configuración

### 1. Clonar el repositorio

```bash
git clone <repository-url>
cd Licenses-Service
```

### 2. Instalar dependencias

```bash
go mod download
go mod tidy
```


## 🏃‍♂️ Ejecución

### Desarrollo Local

```bash
go run cmd/main.go

go build -o bin/licenses-service cmd/main.go
./bin/licenses-service
```

### Con Docker

```bash
docker build -t licenses-service .

docker run -p 8080:8080 licenses-service
```

### Con Docker Compose

```bash
docker-compose up -d

docker-compose logs -f

docker-compose down
```