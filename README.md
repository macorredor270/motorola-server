k# 🚀 Motorola Server

<p align="center">
  <b>Self-hosted mobile server running on Motorola hardware</b><br>
  Android + Arch Linux ARM + Go + Astro
</p>

---

## 📖 Overview

Motorola Server transforms an Android device into a complete self-hosted Linux server environment.

The project combines Android, Arch Linux ARM, Go and Astro to create a lightweight but powerful management platform running directly on mobile hardware.

Features include:

- 🐧 Arch Linux ARM environment inside Android
- ⚡ Go backend API
- 🌌 Astro web dashboard
- 🔌 Real-time WebSocket communication
- 📊 System monitoring
- 🔋 Device telemetry
- 🛠️ Automated server management

The goal is to turn a smartphone into a portable Linux server.

---

# 🏗️ Architecture


Motorola Device
│
├── Android
│ │
│ ├── Magisk
│ │ └── Startup automation
│ │
│ └── Arch Linux ARM Chroot
│ │
│ ├── Go API Server
│ │ Port: 8080
│ │
│ └── Astro Dashboard
│ Port: 4321
│
└── Network
│
└── Browser / Remote Client


---

# 📂 Project Structure


motorola-server/
│
├── backend/ # Go backend
│ │
│ ├── cmd/
│ │ └── server/
│ │ └── main.go # Application entry point
│ │
│ ├── internal/
│ │ │
│ │ ├── api/ # WebSocket API
│ │ ├── handlers/ # HTTP handlers
│ │ ├── services/ # System services
│ │ └── models/ # Data models
│ │
│ └── go.mod
│
├── frontend/ # Astro frontend
│ │
│ ├── src/
│ │ ├── pages/
│ │ ├── components/
│ │ └── lib/
│ │
│ └── package.json
│
├── database/ # Database migrations
│
├── docker/ # Container configuration
│
├── scripts/ # Automation scripts
│
├── logs/ # Runtime logs
│
└── README.md


---

# ⚙️ Features

## ✅ Implemented

- Go HTTP API
- WebSocket real-time communication
- CPU monitoring
- RAM monitoring
- Astro dashboard
- Arch Linux ARM chroot
- SSH administration
- Git based development workflow

---

# 🚧 Planned Features

## Device Monitoring

- 🔋 Battery percentage
- 🌡️ Temperature monitoring
- 💾 Storage information
- 📡 Network statistics
- 📱 Android system information
- ⚡ Charging state

---

## Server Management

- Automatic startup with Magisk
- Chroot lifecycle control
- Service health checks
- Remote command execution
- Background services
- Automatic recovery

---

## Dashboard

- Real-time graphs
- CPU history
- RAM history
- Battery charts
- System logs
- Service status
- Device information cards

---

# 🧰 Technology Stack

## Backend

- Go
- Gorilla WebSocket
- gopsutil

## Frontend

- Astro
- TypeScript
- Tailwind CSS

## System

- Android
- Magisk
- Arch Linux ARM
- OpenSSH

---

# 🚀 Running the Project

## Start Backend

```bash
cd backend

go run cmd/server/main.go
