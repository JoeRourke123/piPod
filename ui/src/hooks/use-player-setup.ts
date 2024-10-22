import {useEffect, useState} from "react";
import {usePlaybackState, usePlayerDevice} from "react-spotify-web-playback-sdk";
import {useParams, useSearchParams} from "react-router-dom";
import {player} from "../util/service";

export const usePlayerSetup = () => {
    const playbackState = usePlaybackState(true, 250);
    const device = usePlayerDevice();
    const { spotifyUri } = useParams();
    const [isStarted, setIsStarted] = useState(false);
    const [searchParams, _] = useSearchParams();

    useEffect(() => {
        if (device && !isStarted) {
            console.log(device);

            let playbackContext = searchParams.get("playback_context");
            if (!playbackContext) {
                playbackContext = playbackState?.context.uri || "";
            }

            player({
                "device_id": device?.device_id,
                "action": "START",
                "spotify_uri": spotifyUri,
                "playback_context": playbackContext
            });

            setIsStarted(true);
        }
    }, [device]);

    return playbackState;
}