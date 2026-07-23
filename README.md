<div align="center">

<img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/badge/Astro-4.x-FF5D01?style=flat-square&logo=astro&logoColor=white" alt="Astro">
<img src="https://img.shields.io/badge/Arch_Linux-ARM-1793D1?style=flat-square&logo=arch-linux&logoColor=white" alt="Arch Linux">
<img src="https://img.shields.io/badge/Android-Rooted-3DDC84?style=flat-square&logo=android&logoColor=white" alt="Android">
<img src="https://img.shields.io/badge/License-MIT-blue?style=flat-square" alt="License">
<img src="https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square" alt="PRs Welcome">

```text
 ██╗   ██╗ ██████╗ ██╗██████╗      ██████╗ ██████╗ ███████╗ █████╗ ██╗  ██╗ ██████╗ ██╗   ██╗████████╗
 ██║   ██║██╔═══██╗██║██╔══██╗    ██╔════╝ ██╔══██╗██╔════╝██╔══██╗██║ ██╔╝██╔═══██╗██║   ██║╚══██╔══╝
 ██║   ██║██║   ██║██║██║  ██║    ██║  ███╗██████╔╝█████╗  ███████║█████╔╝ ██║   ██║██║   ██║   ██║   
 ██║   ██║██║   ██║██║██║  ██║    ██║   ██║██╔══██╗██╔══╝  ██╔══██║██╔═██╗ ██║   ██║██║   ██║   ██║   
 ╚██████╔╝╚██████╔╝██║██████╔╝    ╚██████╔╝██║  ██║███████╗██║  ██║██║  ██╗╚██████╔╝╚██████╔╝   ██║   
  ╚═════╝  ╚═════╝ ╚═╝╚═════╝      ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝  ╚═════╝    ╚═╝   
```

# 🚀 Motorola Server

**Self-Hosted Mobile Server on Motorola Hardware**

Una plataforma todo-en-uno que combina **Android**, **Arch Linux ARM**, **Go** y **Astro** para convertir tu smartphone en un servidor Linux portátil de nivel producción.

</div>

---

## 📑 Tabla de Contenidos

- [Visión General](#-visión-general)
- [Arquitectura](#-arquitectura)
- [Características Clave](#-características-clave)
- [Stack Tecnológico](#-stack-tecnológico)
- [Endpoints de la API](#-endpoints-de-la-api)
- [Estructura del Proyecto](#-estructura-del-proyecto)
- [Inicio Rápido](#-inicio-rápido)
- [Solución de Problemas](#-solución-de-problemas)
- [Roadmap](#-roadmap)
- [Contribuir](#-contribuir)
- [Licencia](#-licencia)

---

## 📖 Visión General

**Motorola Server** da nueva vida a los smartphones Android convirtiéndolos en entornos de servidor Linux autoalojados y eficientes energéticamente.

Aprovechando un contenedor **chroot** con Arch Linux ARM ejecutándose dentro de Android, este proyecto conecta la eficiencia del hardware móvil con capacidades nativas de servidor — todo gestionado a través de un **backend en Go de alto rendimiento** y un **dashboard moderno en Astro**.

A diferencia de soluciones como Termux o UserLAnd, este enfoque proporciona un entorno Linux **nativo en ARM** con acceso completo al sistema de archivos, sin capas de emulación ni abstracciones innecesarias.

> 💡 **¿Por qué un smartphone como servidor?** Consumo energético de ~2-5W frente a 30-60W de un VPS tradicional. Perfecto para homelabs, monitoreo 24/7, automatizaciones ligeras y aprendizaje de infraestructura.

---

## 🏗️ Arquitectura

```text
┌─────────────────────────────────────────────────────────────────────┐
│                       DISPOSITIVO MOTOROLA                       │
│                                                                   │
│  ┌─────────────────────────────────────────────────────────────┐  │
│  │                        Android OS                          │  │
│  │                                                             │  │
│  │  ┌──────────────────┐                                       │  │
│  │  │   Magisk (Root)  │ ──▶ Persistencia & arranque automático │  │
│  │  └────────┬─────────┘                                       │  │
│  │           │                                                 │  │
│  │           ▼                                                 │  │
│  │  ┌───────────────────────────────────────────────────────┐  │  │
│  │  │            Arch Linux ARM (Chroot)                   │  │  │
│  │  │                                                       │  │  │
│  │  │  ┌─────────────────┐    ┌────────────────────────┐   │  │  │
│  │  │  │  Go API Server  │    │   Astro Dashboard      │   │  │  │
│  │  │  │   :8080 (REST)  │◄──►│   :4321 (SSR/Static)  │   │  │  │
│  │  │  │   :8080 (WS)    │    │                        │   │  │  │
│  │  │  └────────┬────────┘    └────────────────────────┘   │  │  │
│  │  │           │                                          │  │  │
│  │  │  ┌────────▼────────┐    ┌────────────────────────┐   │  │  │
│  │  │  │ OpenSSH :22     │    │  Systemd Services      │   │  │  │
│  │  │  └─────────────────┘    └────────────────────────┘   │  │  │
│  │  └───────────────────────────────────────────────────────┘  │  │
│  └─────────────────────────────────────────────────────────────┘  │
└───────────────────────────┬─────────────────────────────────────────┘
                            │  REST / WebSocket (LAN o Tailscale)
                            ▼
                    ┌───────────────┐
                    │ Remote Client │
                    │  (Navegador)  │
                    └───────────────┘
```

---

## ✨ Características Clave

### 🟢 Implementado

| Feature | Descripción |
|---------|-------------|
| ⚡ **Go REST API** | Backend ligero y concurrente para control del sistema con latencia < 5ms en respuestas locales |
| 🔌 **WebSockets en Tiempo Real** | Streaming continuo de métricas del sistema hacia clientes web sin polling |
| 📊 **Monitorización Core** | Métricas en tiempo real de uso de CPU, asignación de RAM y procesos activos |
| 🌌 **Dashboard Astro** | UI web rápida, responsiva y accesible construida con Islands Architecture |
| 🐧 **Chroot Arch Linux ARM** | Ecosistema Linux aislado ejecutándose de forma nativa sobre arquitectura ARM |
| 🔒 **OpenSSH Integrado** | Acceso remoto seguro por línea de comandos desde cualquier dispositivo de la red |

### 🟡 Planificado

| Feature | Descripción |
|---------|-------------|
| 🔋 **Gestión de Energía** | Monitorización de batería, salud, temperatura y estado de carga en tiempo real |
| 💾 **Métricas de Almacenamiento** | Análisis en tiempo real de almacenamiento interno y tarjetas SD externas |
| ⚙️ **Gestor de Servicios** | Controles en la UI para iniciar, detener y reiniciar servicios del chroot |

---

## 🧰 Stack Tecnológico

| Componente | Tecnología | Descripción |
|------------|------------|-------------|
| **Backend** | Go 1.21+ | API REST & WebSockets |
| **Frontend** | Astro 4.x | Dashboard SSR/Static |
| **Sistema** | Arch Linux ARM | Entorno chroot nativo |
| **Host** | Android + Magisk | Root & persistencia |
| **Red** | OpenSSH | Acceso remoto seguro |
| **Protocolo** | WebSocket | Streaming en tiempo real |

---

## 🔌 Endpoints de la API

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| `GET` | `/api/v1/status` | Estado general del servidor y servicios activos |
| `GET` | `/api/v1/metrics/cpu` | Uso de CPU por núcleo y porcentaje global |
| `GET` | `/api/v1/metrics/memory` | RAM usada, disponible y porcentaje de utilización |
| `WS` | `/api/v1/ws/metrics` | Stream continuo de métricas del sistema (JSON cada 1s) |

---

## 📂 Estructura del Proyecto

```text
motorola-server/
├── backend/                    # Aplicación Go
│   ├── cmd/
│   │   └── server/              # Punto de entrada principal
│   │       └── main.go
│   ├── internal/
│   │   ├── api/               # Handlers WebSocket & rutas
│   │   ├── handlers/           # Handlers HTTP estándar
│   │   ├── models/            # Estructuras de datos & tipos
│   │   └── services/          # Lógica de monitorización
│   └── go.mod
│
├── frontend/                   # Aplicación Astro
│   ├── src/
│   │   ├── components/        # Componentes UI reutilizables
│   │   ├── lib/               # Cliente API & helpers
│   │   └── pages/             # Vistas de la aplicación
│   └── package.json
│
├── database/                   # Migraciones & esquemas
├── docker/                     # Configuraciones de contenedores
├── scripts/                    # Setup automatizado & gestión chroot
├── logs/                       # Logs de la aplicación
└── README.md
```

---

## 🚀 Inicio Rápido

> ⚠️ **Requisitos Previos**
> Dispositivo Android rooteado con Magisk · Arch Linux ARM configurado en directorio chroot · Go 1.21+ y Node.js 18+ instalados dentro del chroot

### 1. Clonar el repositorio

```bash
git clone https://github.com/tu-usuario/motorola-server.git
cd motorola-server
```

### 2. Iniciar el Backend

```bash
cd backend
go mod download
go run cmd/server/main.go
```
> El servidor API arrancará en `http://localhost:8080`

### 3. Iniciar el Frontend

```bash
cd frontend
npm install
npm run dev
```
> El dashboard web será accesible en `http://localhost:4321`

### 4. Acceder desde la red

```bash
# Desde otro dispositivo en la misma LAN:
http://<IP_DEL_DISPOSITIVO>:4321

# O vía SSH directamente al chroot:
ssh usuario@<IP_DEL_DISPOSITIVO> -p 22
```

> 💡 **Acceso externo** — Usa [Tailscale](https://tailscale.com) o Cloudflare Tunnel para acceder a tu servidor desde fuera de tu red local sin abrir puertos en el router.

---

## 🔧 Solución de Problemas

| Problema | Solución |
|----------|----------|
| **"Permission denied" al montar chroot** | Verifica que Magisk esté instalado y ejecuta `su` antes de los scripts de montaje. Comprueba que los paths en `scripts/` coincidan con tu dispositivo |
| **El backend no arranca en puerto 8080** | Otro proceso puede estar usando el puerto. Ejecuta `ss -tlnp | grep 8080` para identificarlo, o cambia el puerto en las variables de entorno |
| **WebSocket se desconecta constantemente** | Revisa que no haya un proxy inverso cortando conexiones largas. Si usas nginx, añade `proxy_read_timeout 86400s;` al bloque location |

---

## 🗺️ Roadmap

### Fase 1 — Estabilidad Core ✅

- [x] API REST funcional con endpoints de métricas
- [x] WebSocket para streaming de datos en tiempo real
- [x] Dashboard Astro con visualización de CPU y RAM
- [x] Servidor SSH accesible desde la LAN

### Fase 2 — Monitorización Avanzada

- [ ] 🔋 Métricas de batería: porcentaje, salud, temperatura, estado de carga
- [ ] 💾 Análisis de almacenamiento: interno + SD externa con I/O stats
- [ ] 🌐 Monitorización de red: interfaces, tráfico, conexiones activas
- [ ] 🌡️ Sensores térmicos del dispositivo con alertas configurables

### Fase 3 — Gestión & Automatización

- [ ] ⚙️ Gestor de servicios: start/stop/restart desde el dashboard
- [ ] 🔄 Auto-boot: scripts Magisk para arranque automático al encender
- [ ] 📈 Gráficos históricos: persistencia de métricas con análisis temporal
- [ ] 🔐 Autenticación: login seguro para el dashboard web

### Fase 4 — Producción

- [ ] 📦 Build estático del dashboard para servir desde Go
- [ ] 🐳 Containerización completa con Docker
- [ ] 📝 Documentación completa de API con OpenAPI/Swagger
- [ ] 🧪 Suite de tests unitarios e integración con CI/CD

---

## 🤝 Contribuir

Las contribuciones son bienvenidas. Para participar:

1. Haz un **Fork** del repositorio
2. Crea una rama feature: `git checkout -b feature/nombre-funcionalidad`
3. Realiza tus cambios y haz commit: `git commit -m 'feat: descripción del cambio'`
4. Sube la rama: `git push origin feature/nombre-funcionalidad`
5. Abre un **Pull Request** detallando el problema que resuelve

Por favor, sigue las convenciones de [Conventional Commits](https://www.conventionalcommits.org/) para los mensajes de commit.

---

## 📜 Licencia

Este proyecto está bajo la licencia **MIT**. Puedes consultar el archivo `LICENSE` para más detalles.

---

<div align="center">

Construido con ❤️ usando Go, Astro y un Motorola que se negaba a jubilarse.

</div>
