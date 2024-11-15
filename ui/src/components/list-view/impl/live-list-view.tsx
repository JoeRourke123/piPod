import {PaginatedListViewProps} from "./paginated-list-view";
import React, {useCallback, useEffect, useState} from "react";
import {BaseListView} from "../base-list-view";
import {unmarshallView, useListViewLoader} from "../utils/view-loader";
import {useLocation} from "react-router-dom";
import {useApiUrl} from "../utils/api-url";
import {useItemsHash} from "../utils/items-hash";
import {ListViewItemDetails} from "../list-view-types";

type LiveListViewProps = PaginatedListViewProps & {
    refreshInterval: number;
}

export const LiveListView = (props: LiveListViewProps): React.JSX.Element => {
    let { title, showStatus, items: initialItems , icon, additionalInfo } = useListViewLoader();
    const {key} = useLocation()

    const apiUrl = useApiUrl(props.apiEndpoint);
    const [items, setItems] = useState(initialItems);
    const [itemCount, setItemCount] = useState<number>(items.length);
    const [itemsHash, updateItemsHash] = useItemsHash(items);

    useEffect(() => {
        setItems([]);
        setItemCount(0);
        updateItemsHash([]);
        setItemCount(initialItems.length);
        updateItemsHash(initialItems);
        setItems(initialItems);
    }, [key]);

    const itemLoader = useCallback((currentOffset: number): ListViewItemDetails => {
        return items[currentOffset];
    }, [key, items, itemsHash])

    useEffect(() => {
        const interval = setInterval(() => {
            if (apiUrl) {
                const currentLocation = window.location.href;
                fetch(apiUrl).then(response => response.json()).then(json => {
                    if (window.location.href !== currentLocation) {
                        return;
                    }
                    const updatedView = unmarshallView(json);
                    setItemCount(updatedView.items.length);
                    updateItemsHash(updatedView.items);
                    setItems(updatedView.items);
                });
            }
        }, props.refreshInterval);

        return () => clearInterval(interval);
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