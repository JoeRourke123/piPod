import {PageProps} from "../pages/page-props";
import React, {useCallback, useContext, useEffect} from "react";
import {PlayerDevice, usePlaybackState, usePlayerDevice} from "react-spotify-web-playback-sdk";
import {ClickWheelData, fetchClickWheelData} from "../util/clickwheel-listeners";
import {PlayerContext, PlayerContextWrapper} from "./player/player-provider";
import {CurrentTrackContext, CurrentTrackDetails} from "./player/current-track-context";
import {sendPlayerCommand} from "./player/send-player-command";

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

export const ClickwheelProvider = ({children, socket}: PageProps & { children: React.ReactNode }) => {
    const playbackState = usePlaybackState();
    const device = usePlayerDevice();
    const playerContext = useContext(PlayerContext);
    const [currentTrack, setCurrentTrack] = useContext(CurrentTrackContext);

    const triggerPlayRequest = useCallback((clickwheelData: ClickWheelData) => {
        if (device) {
            if (clickwheelData.isClickWheelPressed) return;

            const action = {
                "Play": "TOGGLE",
                "Back": "BACK",
                "Skip": "SKIP"
            }[clickwheelData.button] || "NONE"

            sendPlayerCommand({
                deviceId: device?.device_id,
                // @ts-ignore
                action: action,
            }).then(handleOfflineBehaviour(clickwheelData));
        }
    }, [device]);


    const handleOfflineBehaviour = (clickwheelData: ClickWheelData) => async (response: CurrentTrackDetails) => {
        if (response.playerState === "OFFLINE") {
            if (clickwheelData.button === "Back" || clickwheelData.button === "Skip") {
                setCurrentTrack(response);
            } else if (clickwheelData.button === "Play") {
                playerContext.toggleAudio();
            }
        }
    }

    const onMessageCallback = useCallback((event: MessageEvent | KeyboardEvent) => {
        const clickwheelData = event instanceof MessageEvent ? fetchClickWheelData(event.data) : makeClickWheelDataFromKeyInput(event.key);
        if (clickwheelData) {
            triggerPlayRequest(clickwheelData);
        }
    }, [device, playbackState, playerContext]);

    useEffect(() => {
        document.addEventListener("keyup", onMessageCallback);
        socket.addEventListener("message", onMessageCallback);

        return () => socket.removeEventListener("message", onMessageCallback);
    }, []);

    return <>{children}</>;
}