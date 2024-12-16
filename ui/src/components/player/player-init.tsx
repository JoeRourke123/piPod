import {useContext, useEffect} from "react";
import {PlayerContext} from "./player-provider";
import {player} from "../../util/service";
import {useParams, useSearchParams} from "react-router-dom";
import {usePlayerDevice} from "react-spotify-web-playback-sdk";
import {CurrentTrackContext} from "./current-track-context";
import {sendPlayerCommand} from "./send-player-command";

export const PlayerInit = ({children}: { children: React.JSX.Element[] }) => {
    const playerContext = useContext(PlayerContext);
    const [currentTrack, setCurrentTrack] = useContext(CurrentTrackContext);

    const device = usePlayerDevice();

    const {spotifyUri, playbackContext} = useParams();

    useEffect(() => {
        if (!playerContext.playing || currentTrack?.spotifyUri !== spotifyUri) {
            if (device) {
                sendPlayerCommand({
                    action: "START",
                    playbackContext,
                    spotifyUri,
                    deviceId: device.device_id,
                }).then(async (response) => {
                    setCurrentTrack(response);
                    playerContext.playerSource = response.playerState as "OFFLINE" | "SPOTIFY" | null;
                });
            }
        }
    }, [device]);

    return <>{children}</>;
}