import {ListViewItem} from "../elems/list-view-item";
import * as React from "react";
import {ItemLoader} from "./item-loader";
import {useMemo} from "react";
import {AdditionalListViewInfo} from "../list-view-types";
import {AdditionalInfoList} from "../elems/list-view-additional-info";

type ItemBuilderOptions = {
    itemLoader: ItemLoader;
    selectedIndex: number;
    itemsHash: string;
    additionalInfo: AdditionalListViewInfo[];
}

export const useRenderItems = (
    { itemLoader, selectedIndex, itemsHash, additionalInfo}: ItemBuilderOptions,
): React.JSX.Element[] => {
    return useMemo(() => itemBuilder(itemLoader, selectedIndex, additionalInfo), [itemsHash, selectedIndex]);
}

const itemBuilder = (itemLoader: ItemLoader, selectedIndex: number, additionalInfo: AdditionalListViewInfo[]): React.JSX.Element[] => {
    const items: React.JSX.Element[] = [];

    if (additionalInfo) {
        items.push(<AdditionalInfoList additionalInfo={additionalInfo} />);
    }

    let itemOffset = 0;
    let item = itemLoader(itemOffset);

    while(item) {
        items.push(<ListViewItem key={item.path} currentIndex={selectedIndex} itemIndex={itemOffset} item={item} />);
        itemOffset++;
        item = itemLoader(itemOffset);
    }

    return items;
};