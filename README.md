# Motorola Server

<p align="center">
  <img alt="Go" src="https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go&logoColor=white">
  <img alt="Astro" src="https://img.shields.io/badge/Astro-frontend-BC52EE?style=flat-square&logo=astro&logoColor=white">
  <img alt="WebSocket" src="https://img.shields.io/badge/WebSocket-tiempo%20real-2b2b2b?style=flat-square">
  <img alt="Licencia MIT" src="https://img.shields.io/badge/licencia-MIT-2b2b2b?style=flat-square">
</p>

Panel de monitoreo en tiempo real para un teléfono Motorola reciclado como mini servidor casero. Un backend en Go lee las métricas del sistema y las transmite por WebSocket a un dashboard hecho con Astro.

## Qué hace (en la práctica)

El repo tiene dos partes independientes:

- **`backend/`** — un servidor HTTP en Go (`net/http`, sin framework) que expone:
  - `GET /api/status` → snapshot puntual en JSON con host, sistema operativo, % de CPU, % de RAM y % de batería.
  - `WS /ws` → el mismo tipo de datos, empujados por WebSocket cada 2 segundos mientras la conexión esté abierta.
  - Las métricas se leen con [`gopsutil`](https://github.com/shirou/gopsutil), que es multiplataforma (no hay nada específico de Android/root en el código Go en sí).
- **`frontend/`** — una página Astro (`src/pages/index.astro`) que abre un WebSocket contra `ws://<host>:8080/ws` y va actualizando en vivo las tarjetas de CPU, RAM y batería, más un indicador de estado de conexión (conectando / conectado / reconectando / sin conexión) con reintento automático. Estilos con Tailwind CSS v4.

Todo corre en el puerto `8080` (backend) y `4321` (frontend, servidor de desarrollo de Astro). No hay base de datos, no hay autenticación, no hay persistencia histórica: es un espejo en vivo del estado del sistema, nada más.

## Por qué existe

Es un proyecto hobby: la idea es aprovechar un Motorola viejo (rooteado, corriendo Arch Linux ARM en un chroot) como servidor casero de bajo consumo en vez de que junte polvo en un cajón, y tener una forma simple de ver desde el navegador si sigue vivo y cómo va de recursos. El setup del root/chroot del teléfono se hace aparte, a mano, en el propio dispositivo — este repositorio es solo el software (API + dashboard), no scripts de aprovisionamiento del sistema.

Nota de transparencia sobre el estado actual del código: el campo `host` que devuelve la API está fijo como `"guam"` (todavía no detecta el nombre real del dispositivo), y `scripts/start.sh` asume que el proyecto vive en `/motorola-server` dentro del sistema donde corre.

## Stack

| Parte     | Tecnología                                                                          |
|-----------|--------------------------------------------------------------------------------------|
| Backend   | Go, `net/http`, [`gorilla/websocket`](https://github.com/gorilla/websocket), [`gopsutil/v4`](https://github.com/shirou/gopsutil) |
| Frontend  | [Astro](https://astro.build), TypeScript, Tailwind CSS v4                            |

## Estructura del proyecto

```
motorola-server/
├── backend/
│   ├── cmd/server/main.go        # entrypoint, registra las rutas HTTP/WS
│   └── internal/
│       ├── api/websocket.go      # handler del WebSocket (/ws)
│       ├── handlers/status.go    # handler REST (/api/status)
│       ├── models/status.go      # struct DeviceStatus
│       └── services/system.go    # lectura de CPU/RAM/batería con gopsutil
├── frontend/
│   └── src/
│       ├── pages/index.astro     # dashboard (única página)
│       └── lib/socket.ts         # cliente WebSocket con reconexión
└── scripts/start.sh              # arranque rápido del backend
```

## Cómo correrlo

Requiere Go 1.21+ y Node 22+.

**Backend**

```bash
cd backend
go mod download
go run cmd/server/main.go
```

Levanta la API en `http://localhost:8080` (REST) y el WebSocket en `ws://localhost:8080/ws`.

**Frontend**

```bash
cd frontend
npm install
npm run dev
```

Dashboard disponible en `http://localhost:4321`. Se conecta automáticamente al backend en el puerto 8080 del mismo host.

Para acceder desde otro dispositivo de la red, usa la IP del equipo que corre el backend en vez de `localhost`.

## Endpoints

| Método | Ruta          | Descripción                                    |
|--------|---------------|--------------------------------------------------|
| `GET`  | `/api/status` | Snapshot único con host, OS, CPU, RAM, batería   |
| `WS`   | `/ws`         | Stream continuo del mismo payload cada 2s        |

## Limitaciones actuales

- Sin autenticación: cualquiera que llegue al puerto 8080 puede leer las métricas.
- Sin persistencia: no se guarda histórico, solo el valor actual.
- Puerto y nombre de host están fijos en el código; todavía no hay variables de entorno configurables.
- No incluye archivo de licencia por ahora.

---

Proyecto hobby, hecho para aprender Go, WebSockets y Astro mientras le doy una segunda vida a un teléfono viejo.

## Licencia

[MIT](LICENSE) © 2026 macorredor270.
