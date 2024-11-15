import {ListViewItemDetails} from "../list-view-types";
import {PageProps} from "../../../pages/page-props";
import React, {useCallback, useEffect, useState} from "react";
import {useLocation, useSearchParams} from "react-router-dom";
import {BaseListView} from "../base-list-view";
import {unmarshallView, useListViewLoader} from "../utils/view-loader";
import {useItemsHash} from "../utils/items-hash";
import {useApiUrl} from "../utils/api-url";
import {PipodCache} from "../../../util/pipod-cache";

export type PaginatedListViewProps = PageProps & {
    apiEndpoint?: string
    useCache?: boolean
}

export const PaginatedListView = (props: PaginatedListViewProps): React.JSX.Element => {
    let { title, showStatus, items: initialItems, icon, additionalInfo } = useListViewLoader();
    const {key} = useLocation();
    const [query] = useSearchParams();
    const apiUrl = useApiUrl(props.apiEndpoint);
    const [items, setItems] = useState(initialItems);
    const [itemCount, setItemCount] = useState<number>(items.length);
    const [itemsHash, updateItemsHash] = useItemsHash(items);
    const [loadingStatus, setLoadingStatus] = useState<0 | 1 | 2>(0);

    useEffect(() => {
        const cachedView = PipodCache.view.get(key);
        if (cachedView) {
            setItems(cachedView.items);
        } else {
            setItems(initialItems);
        }

        setItemCount(initialItems.length);
        updateItemsHash(initialItems);
    }, [key]);

    const itemLoader = useCallback((currentOffset: number): ListViewItemDetails => {
        return items[currentOffset];
    }, [key, items, itemsHash])

    const onSelectedIndexChange = useCallback((index: number) => {
        if (apiUrl && (itemCount - index) <= 10) {
            if (loadingStatus === 0) {
                let apiUrlWithOffset = `${apiUrl}next=${items.length}`;
                if (query.get("filter")) {
                    apiUrlWithOffset += `&filter=${query.get("filter")}`;
                } if (query.get("sort")) {
                    apiUrlWithOffset += `&sort=${query.get("sort")}`;
                }
                setLoadingStatus(1);
                fetch(apiUrlWithOffset).then(response => response.json()).then(json => {
                    const updatedItems = [...items, ...unmarshallView(json).items];
                    setItems(updatedItems);
                    PipodCache.view.set(key, { items: updatedItems, title, icon, additionalInfo, showStatus });
                    setItemCount(updatedItems.length);
                    updateItemsHash(updatedItems);
                    setLoadingStatus(updatedItems.length === items.length ? 2 : 0);
                });
            }
        }

    }, [key, itemsHash, loadingStatus]);

    return <BaseListView
        additionalInfo={additionalInfo}
        title={title}
        showStatus={showStatus}
        itemCount={itemCount}
        itemsHash={itemsHash}
        itemLoader={itemLoader}
        onSelectedIndexChange={onSelectedIndexChange}
        socket={props.socket}
        icon={icon}
    />;
}