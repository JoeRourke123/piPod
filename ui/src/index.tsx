import {ChakraProvider, ColorModeScript, theme} from "@chakra-ui/react"
import * as React from "react"
import reportWebVitals from "./reportWebVitals"
import * as serviceWorker from "./serviceWorker"
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import {Covers} from "./pages/covers";
import "./pipod.css";
import {Auth} from "./pages/auth";
import {WebPlaybackProvider} from "./components/web-playback-provider";
import {Playing} from "./pages/playing";
import ReactDOM from "react-dom";
import {DesktopSettings} from "./pages/desktop-settings";
import {Actions} from "./pages/actions";
import {fetchListView, liveListView, simpleView} from "./util/router-types";

const router = createBrowserRouter([
    liveListView("/", "/views/home", true),
    fetchListView("/music", "/views/music", false, false),
    fetchListView("/list/:type", "/list/:type"),
    fetchListView("/:type/:id", "/:type/:id"),
    liveListView("/queue", "/queue"),
    simpleView("/auth", (s) => <Auth socket={s} />),
    simpleView("/covers", (s) => <Covers socket={s} />),
    simpleView("/playing/:spotifyUri", (s) => <Playing socket={s} />),
    simpleView("/settings", (s) => <DesktopSettings socket={s} />),
    simpleView("/actions", (s) => <Actions socket={s} />),
]);


ReactDOM.render(
  <React.StrictMode>
    <ColorModeScript />
      <ChakraProvider theme={theme}>
          <WebPlaybackProvider>
              <RouterProvider router={router} />
          </WebPlaybackProvider>
      </ChakraProvider>
  </React.StrictMode>,
    document.getElementById("root")
)

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://cra.link/PWA
serviceWorker.unregister()

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()

