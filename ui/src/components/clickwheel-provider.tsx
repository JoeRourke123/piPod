import {PageProps} from "../pages/page-props";
import React, {useCallback, useEffect} from "react";
import {PlayerDevice, usePlaybackState, usePlayerDevice} from "react-spotify-web-playback-sdk";
import {ClickWheelData, fetchClickWheelData} from "../util/clickwheel-listeners";
import {player} from "../util/service";

const triggerPlayRequest = (clickwheelData: ClickWheelData, device: PlayerDevice) => {
    if (clickwheelData.isClickWheelPressed) {
        return;
    }

    switch (clickwheelData.button) {
        case "Back":
            // Go backwards
            player({
                "device_id": device?.device_id,
                "action": "BACK",
            })
            break;
        case "Skip":
            // Go forwards
            player({
                "device_id": device?.device_id,
                "action": "SKIP",
            })
            break;
        case "Play":
            // Play/Pause
            player({
                "device_id": device?.device_id,
                "action": "TOGGLE",
            })
            break;
    }
}

const makeClickWheelDataFromKeyInput = (key: string): ClickWheelData => {
    const buttonMap: Record<string, string> = {
        "a": "Back",
        "d": "Skip",
        "s": "Play",
    }

    if (!buttonMap.hasOwnProperty(key)) {
        return {
            button: "ClickWheel",
            isClickWheelPressed: false,
            clickWheelPosition: 0
        }
    }

    return {
        button: buttonMap[key],
        isClickWheelPressed: false,
        clickWheelPosition: 0
    };
};

export const ClickwheelProvider = ({ children, socket }: PageProps & {children: React.ReactNode}) => {
    const playbackState = usePlaybackState();
    const device = usePlayerDevice();
    const onMessageCallback = useCallback((event: MessageEvent | KeyboardEvent) => {
        const clickwheelData = event instanceof MessageEvent ? fetchClickWheelData(event.data) : makeClickWheelDataFromKeyInput(event.key);
        triggerPlayRequest(clickwheelData, device!);
    }, [device, playbackState]);

    useEffect(() => {
        document.addEventListener("keyup", onMessageCallback);
        socket.addEventListener("message", onMessageCallback);

        return () => socket.removeEventListener("message", onMessageCallback);
    }, []);

    return <>{children}</>;
}