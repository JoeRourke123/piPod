import {ChakraProvider, defaultSystem} from "@chakra-ui/react"
import * as React from "react"
import * as serviceWorker from "./serviceWorker"
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import {Covers} from "./pages/covers";
import "./pipod.css";
import {Auth} from "./pages/auth";
import ReactDOM from "react-dom";
import {DesktopSettings} from "./pages/desktop-settings";
import {Actions} from "./pages/actions";
import {fetchListView, liveListView, simpleView} from "./util/router-types";
import {ClickwheelProvider} from "./components/clickwheel-provider";
import {socket} from "./util/socket";
import {StatusContextProvider} from "./components/status-context-provider";
import {Brickbreaker} from "./pages/games/brickbreaker";
import {PlayerProvider} from "./components/player/player-provider";
import {Player} from "./components/player/player";
import {Database} from "./pages/database";

const router = createBrowserRouter([
    liveListView("/",  true),
    fetchListView("/music", false, false),
    fetchListView("/:type"),
    fetchListView("/:type/:id", false, false),
    liveListView("/queue"),
    simpleView("/auth", (s) => <Auth socket={s}/>),
    simpleView("/covers", (s) => <Covers socket={s}/>),
    simpleView("/playing/:spotifyUri/:playbackContext", (s) => <Player/>),
    simpleView("/settings", (s) => <DesktopSettings socket={s}/>),
    simpleView("/settings/db", (s) => <Database socket={s}/>),
    simpleView("/actions", (s) => <Actions socket={s}/>),
    fetchListView("/games", false, false),
    simpleView("/game/brickbreaker", (s) => <Brickbreaker socket={s}/>)
]);

ReactDOM.render(
    <React.StrictMode>
        <ChakraProvider value={defaultSystem}>
            <PlayerProvider>
                <ClickwheelProvider socket={socket}>
                    <StatusContextProvider socket={socket}>
                        <RouterProvider router={router}/>
                    </StatusContextProvider>
                </ClickwheelProvider>
            </PlayerProvider>
        </ChakraProvider>
    </React.StrictMode>,
    document.getElementById("root")
)

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://cra.link/PWA
serviceWorker.unregister()