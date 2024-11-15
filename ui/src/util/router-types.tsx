import {LoaderFunctionArgs, redirect, RouteObject} from "react-router-dom";
import {PageProps} from "../pages/page-props";
import React from "react";
import {socket} from "./socket";
import {fetchAuthStatus} from "./service";
import {LoaderFunction} from "@remix-run/router/utils";
import {PaginatedListView} from "../components/list-view/impl/paginated-list-view";
import {getApiUrl} from "./functions";
import {LiveListView} from "../components/list-view/impl/live-list-view";

const defaultLoaderFunction = (checkAuthStatus: boolean, apiEndpoint: string) => async ({params, request}: LoaderFunctionArgs) => {
    const query = request.url.split("?")[1];
    if (checkAuthStatus) {
        const authStatus = await fetchAuthStatus();
        if (!authStatus["has_token"]) {
            return redirect("/auth");
        }
    }

    const response = await fetch(getApiUrl(apiEndpoint, params, query));
    return await response.json();
}

export const fetchListView = (
    viewName: string,
    apiEndpoint: string = viewName,
    checkAuthStatus: boolean = false,
    paginated: boolean = true,
    loaderFunction: LoaderFunction = defaultLoaderFunction(checkAuthStatus, apiEndpoint),
): RouteObject => {
    return {
        path: viewName,
        element: <PaginatedListView socket={socket} apiEndpoint={paginated ? apiEndpoint : undefined} />,
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
    apiEndpoint: string,
    checkAuthStatus: boolean = false,
    refreshInterval: number = 2500,
    loaderFunction: LoaderFunction = defaultLoaderFunction(checkAuthStatus, apiEndpoint),
): RouteObject => {

    return {
        path: viewName,
        element: <LiveListView socket={socket} apiEndpoint={apiEndpoint} refreshInterval={refreshInterval}/>,
        loader: loaderFunction,
    }
}