# Motorola Server

Panel de monitoreo en tiempo real para sistemas. Visualiza CPU, RAM, batería y estado del SO con actualizaciones automáticas vía WebSocket.

## Características

✨ **Frontend Astro + Tailwind**
- Interfaz moderna y responsiva con diseño glassmorphism
- Panel de monitoreo en tiempo real
- Indicadores visuales de estado de conexión
- Reconexión automática con reintentos

⚡ **Backend Go**
- API REST + WebSocket para streaming de datos
- Monitoreo de CPU, RAM, batería y SO
- Actualizaciones cada 2 segundos
- Manejo robusto de conexiones

🐳 **DevOps**
- Configuración Docker lista
- Scripts de inicio automatizados
- Migraciones de base de datos

## Estructura

```
motorola-server/
├── frontend/          # Astro + Tailwind CSS
│   ├── src/pages/     # Página principal
│   ├── src/lib/       # WebSocket client
│   ├── src/styles/    # Estilos globales
│   └── tailwind.config.js
├── backend/           # Go API
│   ├── cmd/server/    # Punto de entrada
│   ├── internal/
│   │   ├── api/       # WebSocket handler
│   │   ├── handlers/  # HTTP endpoints
│   │   ├── services/  # Lógica de sistema
│   │   └── models/    # Estructuras
│   └── go.mod
├── database/          # Migraciones SQL
├── docker/            # Dockerfile y compose
└── scripts/           # Automation
```

## Instalación

### Frontend
```bash
cd frontend
npm install
npm run dev
```

### Backend
```bash
cd backend
go mod download
go run ./cmd/server/main.go
```

## API

### WebSocket
- **URL**: `ws://localhost:8080/ws`
- **Intervalo**: Datos cada 2s
- **Datos enviados**:
  ```json
  {
    "host": "guam",
    "os": "linux",
    "cpu": 25.5,
    "ram": 80.3,
    "battery": 85,
    "time": "2026-07-23T10:30:45Z"
  }
  ```

### REST
- **GET** `/api/status` - Estado actual del sistema

## Desarrollo

Con Docker Compose:
```bash
docker-compose up
```

Frontend en `http://localhost:4321`
Backend en `http://localhost:8080`

## Stack

- **Frontend**: Astro, Tailwind CSS, TypeScript
- **Backend**: Go, Gorilla WebSocket, gopsutil
- **Database**: PostgreSQL (opcional)
- **DevOps**: Docker, Docker Compose
