import {Socket} from "socket.io-client";

export type PageProps = {
    socket: WebSocket,
    basePageLoaderUrl?: string
};