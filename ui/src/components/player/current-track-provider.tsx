import {PlayerContext} from "./player-provider";
import {CurrentTrackContext, CurrentTrackDetails} from "./current-track-context";
import {useCallback, useContext, useEffect, useState} from "react";
import {sendPlayerCommand} from "./send-player-command";
import {usePlaybackState, usePlayerDevice, useSpotifyPlayer} from "react-spotify-web-playback-sdk";

export const CurrentTrackProvider = ({ children }: { children: React.JSX.Element }) => {
    const currentTrackState = useState<CurrentTrackDetails | undefined>(undefined);

    const [currentTrack, setCurrentTrack] = currentTrackState;

    const playerContext = useContext(PlayerContext);

    const device = usePlayerDevice();

    const player = useSpotifyPlayer();

    const playbackState = usePlaybackState(true, 500);

    const [isFetching, setIsFetching] = useState(false);

    const playerStateListener = useCallback((state: any) => {
        playerContext.playing = !state.paused;

        if (playerContext.playing && device) {
            if (state.track_window.current_track.uri !== currentTrack?.spotifyUri && !isFetching) {
                setIsFetching(true);
                playerContext.playerSource = "SPOTIFY";
                sendPlayerCommand({
                    action: "PLAYING",
                    deviceId: device?.device_id
                }).then((currentTrack) => {
                    setCurrentTrack(currentTrack);
                    setIsFetching(false);
                })
            }
        }
    }, [currentTrack, isFetching, device]);

    useEffect(() => {
        player?.on("player_state_changed", playerStateListener);

        return () => {
            player?.on("player_state_changed", () => {});
        }
    }, [player, device, playerContext]);

    useEffect(() => {
        if (playbackState) {
            playerContext.setPosition(playbackState.position);
        }
    }, [playbackState]);

    useEffect(() => {
        playerContext.onTrackEnded = (_) => {
            sendPlayerCommand({
                action: "TRIGGER",
            }).then((response) => {
                setCurrentTrack(response);
            });
        };
    }, []);

    useEffect(() => {
        if (playerContext.playerSource === "OFFLINE" && currentTrack !== undefined) {
            playerContext.playAudio(currentTrack);
        }
    }, [playerContext, currentTrack]);

    return <CurrentTrackContext.Provider value={currentTrackState}>
        { children }
    </CurrentTrackContext.Provider>
}