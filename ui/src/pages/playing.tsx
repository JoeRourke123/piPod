import {PageProps} from "./page-props";
import {
    usePlaybackState,
    usePlayerDevice,
} from "react-spotify-web-playback-sdk";
import {useEffect, useState} from "react";
import {useLocation, useParams, useSearchParams} from "react-router-dom";
import {Container} from "@chakra-ui/react";
import {PlayerView} from "../components/player-view";

export const Playing = (props: PageProps) => {
    const { spotifyUri } = useParams();
    const [isStarted, setIsStarted] = useState(false)
    const device = usePlayerDevice();
    const playbackState = usePlaybackState(true, 250);
    const [searchParams, _] = useSearchParams();

    useEffect(() => {
        if (device && !isStarted) {
            console.log(device);

            let playbackContext = searchParams.get("playback_context");
            if (!playbackContext) {
                playbackContext = playbackState?.context.uri || "";
            }

            fetch("http://localhost:9091/player", {
                method: "POST",
                body: JSON.stringify({
                    "device_id": device?.device_id,
                    "action": "START",
                    "spotify_uri": spotifyUri,
                    "playback_context": playbackContext
                }),
                headers: {
                    "Content-Type": "application/json"
                }
            });
            setIsStarted(true);
        }
    }, [device]);

    return (
        <Container>
            { playbackState && <PlayerView playbackState={playbackState} /> }
        </Container>
    )
}