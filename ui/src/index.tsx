import {ChakraProvider, ColorModeScript, theme} from "@chakra-ui/react"
import * as React from "react"
import * as ReactDOM from "react-dom/client"
import { Home } from "./pages/Home"
import reportWebVitals from "./reportWebVitals"
import * as serviceWorker from "./serviceWorker"
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import {Covers} from "./pages/Covers";
import "./pipod.css";
import {io} from "socket.io-client";
import {ListViewProps} from "./util/ListViewTypes";
import {VinylRecord} from "@phosphor-icons/react";
import {List} from "./pages/List";
import {Music} from "./pages/Music";
import {CategoryList} from "./pages/CategoryList";


const container = document.getElementById("root")
if (!container) throw new Error('Failed to find the root element');
const root = ReactDOM.createRoot(container);

const socket = new WebSocket("ws://192.168.1.162:9091/ws");

const router = createBrowserRouter([
    {
        path: "/",
        element: <Home socket={socket} />,
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
        path: "list/:type",
        element: <CategoryList socket={socket} />,
        loader: async ({ request, params }) => {
            const response = await fetch(
                `http://localhost:9091/list/${ params.type }`,
            );

            const listViewProps: ListViewProps = await response.json();

            return listViewProps
        },
    },
    {
        path: "/list/:type/:id",
        element: <List socket={socket} />,
        loader: ({  params }): ListViewProps => {
            return {
                title: "brat",
                fallbackIcon: (c: string) => <VinylRecord scale={12} color={c} />,
                items: [
                    { title: "360" },
                    { title: "Club classics" },
                    { title: "Sympathy is a knife" },
                    { title: "I may say something stupid" },
                    { title: "Talk talk" },
                    { title: "Von dutch" },
                ]
            }
        }
    }
]);

root.render(
  <React.StrictMode>
    <ColorModeScript />
      <ChakraProvider theme={theme}>
        <RouterProvider router={router} />
      </ChakraProvider>
  </React.StrictMode>,
)

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://cra.link/PWA
serviceWorker.unregister()

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()

