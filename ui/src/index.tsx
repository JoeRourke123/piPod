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


const container = document.getElementById("root")
if (!container) throw new Error('Failed to find the root element');
const root = ReactDOM.createRoot(container);

const socket = io("pipod.local:9090", {
    addTrailingSlash: false,
    path: "",
    transportOptions: {
        websocket: {
            path: ""
        }
    }
});

const router = createBrowserRouter([
    {
        path: "/",
        element: <Home socket={socket} />,
    },
    {
        path: "/covers",
        element: <Covers socket={socket} />,
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

