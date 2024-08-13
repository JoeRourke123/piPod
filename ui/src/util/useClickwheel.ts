import {useEffect} from "react";
import {fetchClickWheelData} from "./clickwheelListeners";

export function useClickwheel(socket: WebSocket, onClickWheelChange: any) {
    const onMessageHandler = (e: MessageEvent) => {
        const clickWheelData = fetchClickWheelData(e);

        if (clickWheelData.button === "ClickWheel") {
            onClickWheelChange(clickWheelData)
        }
    }

    useEffect(() => {
        socket.addEventListener("message", onMessageHandler);

        return () => {
            socket.removeEventListener("message", onMessageHandler);
        }
    }, []);
}

export function useSelectButton(socket: WebSocket, onSelectButtonUp: any) {
    const onMessageHandler = (e: MessageEvent) => {
        const clickWheelData = fetchClickWheelData(e);

        if (clickWheelData.button === "Select" && !clickWheelData.isClickWheelPressed) {
            onSelectButtonUp()
        }
    }

    useEffect(() => {
        socket.addEventListener("message", onMessageHandler);

        return () => {
            socket.removeEventListener("message", onMessageHandler);
        }
    }, []);
}