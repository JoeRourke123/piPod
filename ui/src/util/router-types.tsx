import {LoaderFunctionArgs, redirect, RouteObject} from "react-router-dom";
import {PageProps} from "../pages/page-props";
import React from "react";
import {socket} from "./socket";
import {fetchAuthStatus} from "./service";
import {LoaderFunction} from "@remix-run/router/utils";
import {PaginatedListView} from "../components/list-view/impl/paginated-list-view";
import {getApiUrl} from "./functions";
import {LiveListView} from "../components/list-view/impl/live-list-view";

const defaultLoaderFunction = (checkAuthStatus: boolean) => async ({request}: LoaderFunctionArgs) => {
    const url = new URL(request.url);
    if (checkAuthStatus) {
        const authStatus = await fetchAuthStatus();
        if (!authStatus["hasToken"]) {
            return redirect("/auth");
        }
    }

    const response = await fetch(getApiUrl(url.pathname, Object.fromEntries(url.searchParams)));
    return await response.json();
}

export const fetchListView = (
    viewName: string,
    checkAuthStatus: boolean = false,
    paginated: boolean = true,
    loaderFunction: LoaderFunction = defaultLoaderFunction(checkAuthStatus),
): RouteObject => {
    return {
        path: viewName,
        element: <PaginatedListView socket={socket} paginated={paginated} />,
        loader: loaderFunction,
    }
}

export const simpleView = (
    viewName: string,
    viewObject: (_: WebSocket) => React.JSX.Element,
): RouteObject => {
    return {
        path: viewName,
        element: viewObject(socket),
    }
}

export const liveListView = (
    viewName: string,
    checkAuthStatus: boolean = false,
    refreshInterval: number = 2500,
    loaderFunction: LoaderFunction = defaultLoaderFunction(checkAuthStatus),
): RouteObject => {

    return {
        path: viewName,
        element: <LiveListView socket={socket} refreshInterval={refreshInterval}/>,
        loader: loaderFunction,
    }
}