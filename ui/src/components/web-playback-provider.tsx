import React, {useCallback, useEffect} from "react";
import {fetchSpotifyToken} from "../util/service";
import {WebPlaybackSDK} from "react-spotify-web-playback-sdk";

export const WebPlaybackProvider = (props: { children: React.ReactElement }) => {
    const getOAuthToken = useCallback(async (callback: (_: string) => void) => {
        const token = await fetchSpotifyToken();
        callback(token);
    }, []);

    return <WebPlaybackSDK
        initialDeviceName={"PiPod"}
        connectOnInitialized={true}
        getOAuthToken={getOAuthToken}
    >{ props.children }</WebPlaybackSDK>;
}