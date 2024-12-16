import React, {createContext, useEffect, useState} from "react";
import {WebPlaybackProvider} from "../web-playback-provider";
import {CurrentTrackProvider} from "./current-track-provider";
import {CurrentTrackDetails} from "./current-track-context";

export class PlayerContextWrapper {
    private audio: HTMLAudioElement | undefined;
    public playing: boolean = false;
    private position: number = 0;

    public playerSource: "OFFLINE" | "SPOTIFY" | null = null;
    public offlinePlayerStateChanged: (audio: HTMLAudioElement) => void = () => {};
    public onPositionUpdated: (position: number) => void = () => {};
    public onTrackEnded: (audio: HTMLAudioElement) => void = () => {};

    private constructor(audio: HTMLAudioElement | undefined = new Audio()) {
        this.audio = audio;
        this.audio.onpause = () => this.audio && this.offlinePlayerStateChanged(this.audio);
        this.audio.onplay = () => this.audio && this.offlinePlayerStateChanged(this.audio);
        this.audio.onended = () => this.audio && this.onTrackEnded(this.audio);
        this.audio?.addEventListener("timeupdate", () => this.audio && this.offlinePlayerStateChanged(this.audio));
    }

    static getInitialContext = () => new PlayerContextWrapper(undefined);
    static getInitialState = () => new PlayerContextWrapper();

    public setPosition(position: number) {
        this.position = position;
        this.onPositionUpdated(position);
    }

    public getPosition() {
        return this.position;
    }

    public playAudio = (currentTrack: CurrentTrackDetails) => {
        this.audio?.setAttribute("src", currentTrack.playerUrl);
        this.audio?.load();
        this.audio?.play();
    }

    public toggleAudio = () => {
        if (this.playing) {
            this.audio?.pause();
        } else {
            this.audio?.play();
        }
    }
}

export const PlayerContext = createContext<PlayerContextWrapper>(PlayerContextWrapper.getInitialContext());

export const PlayerProvider = ({children}: { children: React.JSX.Element }) => {
    const [offlinePlayer, _] = useState(PlayerContextWrapper.getInitialState);

    return <WebPlaybackProvider>
        <PlayerContext.Provider value={offlinePlayer}>
            <PlayerProviderSetup playerContext={offlinePlayer}>
                <CurrentTrackProvider>
                    {children}
                </CurrentTrackProvider>
            </PlayerProviderSetup>
        </PlayerContext.Provider>
    </WebPlaybackProvider>;
}

const PlayerProviderSetup = ({children, playerContext}: { children: React.JSX.Element, playerContext: PlayerContextWrapper }) => {
    useEffect(() => {
        playerContext.offlinePlayerStateChanged = (audio: HTMLAudioElement) => {
            playerContext.playing = audio.paused;
            playerContext.setPosition(audio.currentTime * 1000);
        }
    }, []);

    return <>{children}</>;
}