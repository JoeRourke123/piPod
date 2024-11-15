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

    const {spotifyUri} = useParams();
    const [searchParams, _] = useSearchParams();

    useEffect(() => {
        if (!playerContext.playing || currentTrack?.spotifyUri !== spotifyUri) {
            const playbackContext = searchParams.get("playback_context") || undefined;
            const albumId = searchParams.get("album_id") || undefined;

            if (device) {
                sendPlayerCommand({
                    action: "START",
                    playbackContext,
                    albumId,
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