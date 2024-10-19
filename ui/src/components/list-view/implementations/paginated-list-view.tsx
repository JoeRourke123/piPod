import {ListViewItemDetails} from "../../../util/ListViewTypes";
import {PageProps} from "../../../state/PageProps";
import React, {useCallback, useEffect, useState} from "react";
import {useLocation} from "react-router-dom";
import {BaseListView} from "../base-list-view";
import {useListViewLoader} from "../utils/view-loader";
import {useItemsHash} from "../utils/items-hash";
import {useApiUrl} from "../utils/api-url";

export type PaginatedListViewProps = PageProps & {
    apiEndpoint?: string
}

export const PaginatedListView = (props: PaginatedListViewProps): React.JSX.Element => {
    let { title, showStatus, items: initialItems } = useListViewLoader();
    const {key} = useLocation()

    const apiUrl = useApiUrl(props.apiEndpoint);
    const [items, setItems] = useState(initialItems);
    const [itemCount, setItemCount] = useState<number>(items.length);
    const [itemsHash, updateItemsHash] = useItemsHash(items);
    const [loadingStatus, setLoadingStatus] = useState<0 | 1 | 2>(0);

    useEffect(() => {
        setItems(initialItems);
        setItemCount(initialItems.length);
        updateItemsHash(initialItems);
    }, [key]);

    const itemLoader = useCallback((currentOffset: number): ListViewItemDetails => {
        return items[currentOffset];
    }, [key, items, itemsHash])

    const onSelectedIndexChange = useCallback((index: number) => {
        if (apiUrl && (itemCount - index) <= 25) {
            if (loadingStatus === 0) {
                const apiUrlWithOffset = `${apiUrl}?next=${items.length}`;
                setLoadingStatus(1);
                fetch(apiUrlWithOffset).then(response => response.json()).then(json => {
                    const updatedItems = [...items, ...json["items"]];
                    setItems(updatedItems);
                    setItemCount(updatedItems.length);
                    updateItemsHash(updatedItems);
                    setLoadingStatus(updatedItems.length === items.length ? 2 : 0);
                });
            }
        }

    }, [key, itemsHash, loadingStatus]);

    return <BaseListView
        title={title}
        showStatus={showStatus}
        itemCount={itemCount}
        itemsHash={itemsHash}
        itemLoader={itemLoader}
        onSelectedIndexChange={onSelectedIndexChange}
        socket={props.socket}/>;
}