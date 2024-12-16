import {CurrentTrackDetails} from "./current-track-context";
import {player} from "../../util/service";

type PlayerOptions = {
    action: "START" | "TOGGLE" | "BACK" | "SKIP" | "TRIGGER" | "PLAYING" | "NONE";
    spotifyUri?: string;
    albumId?: string;
    playbackContext?: string;
    deviceId?: string;
};

export const sendPlayerCommand = (playerOpts: PlayerOptions): Promise<CurrentTrackDetails> => {
    if (playerOpts.action === "NONE") return Promise.resolve({} as CurrentTrackDetails);

    return player(playerOpts).then(async (response) => {
        const responseJson = await response.json();
        return responseJson as CurrentTrackDetails;
    });
}