import {PlayerInit} from "./player-init";
import {PlayerView} from "./player-view";
import {PlayerBackground} from "./player-background";

export const Player = () => {
    return <>
        <PlayerInit>
            <PlayerBackground />
            <PlayerView />
        </PlayerInit>
    </>;
}