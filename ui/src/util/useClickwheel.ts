import {useEffect} from "react";
import {Socket} from "socket.io-client";
import {fetchClickwheelData} from "./clickwheelListeners";

export function useClickwheel(maxItems: number, socket: Socket) {
    useEffect(() => {
        socket.onAny(fetchClickwheelData)
    }, []);

    return 0;
}