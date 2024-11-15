import {PlayerContext} from "./player-provider";
import {CurrentTrackContext, CurrentTrackDetails} from "./current-track-context";
import {useContext, useEffect, useState} from "react";
import {sendPlayerCommand} from "./send-player-command";

export const CurrentTrackProvider = ({ children }: { children: React.JSX.Element }) => {
    const currentTrackState = useState<CurrentTrackDetails | undefined>(undefined);

    const [currentTrack, setCurrentTrack] = currentTrackState;

    const playerContext = useContext(PlayerContext);

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
    }, [currentTrack]);

    return <CurrentTrackContext.Provider value={currentTrackState}>
        { children }
    </CurrentTrackContext.Provider>
}