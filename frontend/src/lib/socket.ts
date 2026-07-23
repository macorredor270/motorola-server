export type ConnectionState = "connecting" | "connected" | "reconnecting" | "offline";

export interface DeviceStatus {
  host: string;
  os?: string;
  cpu: number;
  ram: number;
  battery?: number;
  status?: string;
  time?: string | Date;
}

export function connectSocket(
  onMessage: (data: DeviceStatus) => void,
  onStateChange?: (state: ConnectionState, message?: string) => void
) {
  const protocol = window.location.protocol === "https:" ? "wss" : "ws";
  const socketUrl = `${protocol}://${window.location.hostname}:8080/ws`;
  let socket: WebSocket | null = null;
  let retryTimer: number | undefined;

  const setState = (state: ConnectionState, message?: string) => {
    onStateChange?.(state, message);
  };

  const connect = () => {
    if (retryTimer) {
      window.clearTimeout(retryTimer);
    }

    setState("connecting", "Estableciendo conexión con el servidor WebSocket…");
    socket = new WebSocket(socketUrl);

    socket.addEventListener("open", () => {
      setState("connected", "Conectado al servidor.");
    });

    socket.addEventListener("message", (event) => {
      try {
        const payload = JSON.parse(event.data) as DeviceStatus;
        onMessage(payload);
      } catch {
        setState("offline", "No se pudo leer el payload del servidor.");
      }
    });

    socket.addEventListener("close", () => {
      if (socket) {
        setState("reconnecting", "Conexión perdida. Intentando reconectar...");
        retryTimer = window.setTimeout(connect, 2000);
      }
    });

    socket.addEventListener("error", () => {
      setState("offline", "No se pudo conectar al servidor.");
      socket?.close();
    });
  };

  connect();

  return {
    close() {
      if (retryTimer) {
        window.clearTimeout(retryTimer);
      }
      socket?.close();
      socket = null;
    },
  };
}
