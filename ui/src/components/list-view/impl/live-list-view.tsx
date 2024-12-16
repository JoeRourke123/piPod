import {PaginatedListViewProps} from "./paginated-list-view";
import React, {useCallback, useEffect, useState} from "react";
import {BaseListView} from "../base-list-view";
import {unmarshallView, useListViewLoader} from "../utils/view-loader";
import {useLocation} from "react-router-dom";
import {useApiUrl} from "../utils/api-url";
import {useItemsHash} from "../utils/items-hash";
import {ListViewItemDetails} from "../list-view-types";
import {Simulate} from "react-dom/test-utils";
import abort = Simulate.abort;

type LiveListViewProps = Omit<PaginatedListViewProps, 'paginated'> & {
    refreshInterval: number;
}

export const LiveListView = (props: LiveListViewProps): React.JSX.Element => {
    let { title, showStatus, items: initialItems , icon, additionalInfo } = useListViewLoader();
    const {key} = useLocation()

    const apiUrl = useApiUrl();
    const [items, setItems] = useState(initialItems);
    const [itemCount, setItemCount] = useState<number>(items.length);
    const [itemsHash, updateItemsHash] = useItemsHash(items);

    const itemLoader = useCallback((currentOffset: number): ListViewItemDetails => {
        return items[currentOffset];
    }, [key, items, itemsHash, itemCount]);

    useEffect(() => {
        setItems([]);
        updateItemsHash([]);
        setItemCount(0);
        setItems(initialItems);
        updateItemsHash(initialItems);
        setItemCount(initialItems.length);
        const abortController = new AbortController();
        const interval = setInterval(() => {
            if (apiUrl) {
                const currentLocation = window.location.href;
                fetch(apiUrl, { signal: abortController.signal }).then(response => response.json()).then(json => {
                    if (window.location.href !== currentLocation) {
                        return;
                    }
                    const updatedView = unmarshallView(json);
                    setItems(updatedView.items);
                    updateItemsHash(updatedView.items);
                    setItemCount(updatedView.items.length);
                });
            }
        }, props.refreshInterval);

        return () => {
            abortController.abort("moved page");
            clearInterval(interval);
        };
    }, [key]);

    return <BaseListView
        title={title}
        showStatus={showStatus}
        itemCount={itemCount}
        itemsHash={itemsHash}
        itemLoader={itemLoader}
        socket={props.socket}
        additionalInfo={additionalInfo}
        icon={icon}
    />;
}