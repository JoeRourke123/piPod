import {ChakraProvider, ColorModeScript, theme} from "@chakra-ui/react"
import * as React from "react"
import {Home} from "./pages/Home"
import reportWebVitals from "./reportWebVitals"
import * as serviceWorker from "./serviceWorker"
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import {Covers} from "./pages/Covers";
import "./pipod.css";
import {ListViewProps} from "./util/ListViewTypes";
import {List} from "./pages/List";
import {Music} from "./pages/Music";
import {CategoryList} from "./pages/CategoryList";
import {Auth} from "./pages/Auth";
import {fetchAlbums, fetchAuthStatus, fetchPlaylist} from "./util/service";
import {WebPlaybackSDK} from "react-spotify-web-playback-sdk";
import {WebPlaybackProvider} from "./components/WebPlaybackProvider";
import {Playing} from "./pages/Playing";
import ReactDOM from "react-dom";


const socket = new WebSocket("ws://192.168.1.162:9091/ws");

const router = createBrowserRouter([
    {
        path: "/",
        element: <Home socket={socket} />,
        loader: fetchAuthStatus
    },
    {
      path: "/auth",
      element: <Auth socket={socket} />,
    },
    {
        path: "/covers",
        element: <Covers socket={socket} />,
    },
    {
        path: "/music",
        element: <Music socket={socket} />
    },
    {
        path: "/:type/:id",
        element: <CategoryList socket={socket} basePageLoaderUrl="/" />,
        loader: async ({  params }): Promise<ListViewProps> => {
            if (params.type === "playlists") {
                return fetchPlaylist(params.id, 0);
            } else {
                return fetchAlbums(params.id, 0);
            }
        }
    },
    {
        path: "list/:type",
        element: <CategoryList socket={socket} basePageLoaderUrl="/list/" />,
        loader: async ({ request, params }): Promise<ListViewProps> => {
            const response = await fetch(
                `http://localhost:9091/list/${ params.type }`,
            );

            return await response.json();
        },
    },
    {
        path: "/playing/:spotifyUri",
        element: <Playing socket={socket} />
    }
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

