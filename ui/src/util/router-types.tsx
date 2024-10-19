import {LoaderFunctionArgs, redirect, RouteObject} from "react-router-dom";
import {CategoryListProps} from "./ListViewTypes";
import {PageProps} from "../state/PageProps";
import React from "react";
import {socket} from "./socket";
import {fetchAuthStatus} from "./service";
import {LoaderFunction} from "@remix-run/router/utils";
import {Simulate} from "react-dom/test-utils";
import {PaginatedListView} from "../components/list-view/implementations/paginated-list-view";
import {getApiUrl} from "./functions";
import {LiveListView} from "../components/list-view/implementations/live-list-view";


const defaultLoaderFunction = (checkAuthStatus: boolean, apiEndpoint: string) => async ({params}: LoaderFunctionArgs) => {
    if (checkAuthStatus) {
        const authStatus = await fetchAuthStatus();
        if (!authStatus["has_token"]) {
            return redirect("/auth");
        }
    }

    const response = await fetch(getApiUrl(apiEndpoint, params));
    return await response.json();
}

export const fetchListView = (
    viewName: string,
    apiEndpoint: string = viewName,
    checkAuthStatus: boolean = false,
    paginated: boolean = true,
    loaderFunction: LoaderFunction = defaultLoaderFunction(checkAuthStatus, apiEndpoint),
    props?: CategoryListProps): RouteObject => {
    return {
        path: viewName,
        element: <PaginatedListView socket={socket} apiEndpoint={paginated ? apiEndpoint : undefined} />,
        loader: loaderFunction,
    }
}

export const simpleView = (
    viewName: string,
    viewObject: (_: WebSocket) => React.JSX.Element,
    props?: PageProps): RouteObject => {
    return {
        path: viewName,
        element: viewObject(socket),
    }
}

export const liveListView = (
    viewName: string,
    apiEndpoint: string,
    checkAuthStatus: boolean = false,
    refreshInterval: number = 1000,
    loaderFunction: LoaderFunction = defaultLoaderFunction(checkAuthStatus, apiEndpoint),
): RouteObject => {

    return {
        path: viewName,
        element: <LiveListView socket={socket} apiEndpoint={apiEndpoint} refreshInterval={refreshInterval}/>,
        loader: loaderFunction,
    }
}