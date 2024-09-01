import {PageProps} from "../state/PageProps";
import {
    useErrorState,
    usePlaybackState,
    usePlayerDevice,
    useSpotifyPlayer,
    useWebPlaybackSDKReady
} from "react-spotify-web-playback-sdk";
import {useEffect, useState} from "react";
import {useLocation, useParams} from "react-router-dom";
import {Container} from "@chakra-ui/react";
import {PlayerView} from "../components/PlayerView";

export const Playing = (props: PageProps) => {
    const { spotifyUri } = useParams();
    const [isStarted, setIsStarted] = useState(false)
    const device = usePlayerDevice();
    const playbackState = usePlaybackState(true, 250);

    useEffect(() => {
        if (device && !isStarted) {
            console.log(device);
            fetch("http://localhost:9091/player", {
                method: "POST",
                body: JSON.stringify({
                    "device_id": device?.device_id,
                    "action": "START",
                    "spotify_uri": spotifyUri
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