import React, {createContext, useCallback, useEffect, useMemo, useState} from "react";
import {PageProps} from "../pages/page-props";

export type StatusContextDetails = {
    isInternetEnabled: boolean;
}

export const StatusContext = createContext<StatusContextDetails>({ isInternetEnabled: false})

export const StatusContextProvider = (props: PageProps & { children: React.JSX.Element}) => {
    const [internetEnabled, setInternetEnabled] = useState<boolean>(false);

    const status = useMemo<StatusContextDetails>(() => {
        return {
            isInternetEnabled: internetEnabled
        }
    }, [internetEnabled]);

    const onMessageCallback = useCallback((event: MessageEvent) => {
        if (event.data) {
            const messageData = JSON.parse(event.data);

            if (messageData["isInternetEnabled"] !== undefined) {
                setInternetEnabled(messageData["isInternetEnabled"]);
            }
        }
    }, []);

    useEffect(() => {
        props.socket.addEventListener("message", onMessageCallback);

        return () => {
            props.socket.removeEventListener("message", onMessageCallback);
        }
    }, []);

    return <StatusContext.Provider value={status}>
        { props.children }
    </StatusContext.Provider>
}