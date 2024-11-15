import {createContext, Dispatch, SetStateAction} from "react";

export type CurrentTrackDetails = {
    spotifyUri: string;
    trackName: string;
    artist: string;
    album: string;
    albumId: string;
    imageUrl: string;
    duration: number;
    playerUrl: string;
    playerState: string;
}

type CurrentTrackState = [CurrentTrackDetails | undefined, Dispatch<SetStateAction<CurrentTrackDetails | undefined>>];

export const CurrentTrackContext = createContext<CurrentTrackState>([undefined, () => {}]);
