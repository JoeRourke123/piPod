import {ListViewItem} from "../sub-components/list-view-item";
import * as React from "react";
import {ListViewItemDetails} from "../../../util/ListViewTypes";
import {useCallback, useMemo} from "react";
import {ItemLoader} from "./item-loader";

type ItemBuilderOptions = {
    itemLoader: ItemLoader,
    selectedIndex: number,
    itemsHash: string
}

export const useRenderItems = (
    { itemLoader, selectedIndex, itemsHash}: ItemBuilderOptions,
): React.JSX.Element[] => {
    return useMemo(() => itemBuilder(itemLoader, selectedIndex), [itemsHash, selectedIndex]);
}

const itemBuilder = (itemLoader: ItemLoader, selectedIndex: number): React.JSX.Element[] => {
    const items: React.JSX.Element[] = [];
    let itemOffset = 0;
    let item = itemLoader(itemOffset);

    while(item) {
        items.push(<ListViewItem currentIndex={selectedIndex} itemIndex={itemOffset} item={item} />);
        itemOffset++;
        item = itemLoader(itemOffset);
    }

    return items;
};