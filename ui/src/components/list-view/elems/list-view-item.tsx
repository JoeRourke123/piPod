import {ListViewItemDetails} from "../list-view-types";
import React, {useMemo} from "react";
import {useLocation} from "react-router-dom";
import {ListViewItemContent} from "./list-view-item-content";
import {ListViewItemCard} from "./list-view-item-card";

export type ListViewItemProps = {
    currentIndex: number
    itemIndex: number
    item: ListViewItemDetails
};

export const ListViewItem = ({currentIndex, itemIndex, item}: ListViewItemProps) => {
    const {key} = useLocation();

    const isSelected = useMemo(() => {
        return currentIndex === itemIndex;
    }, [currentIndex, itemIndex, key]);

    return <>
        <ListViewItemCard isSelected={isSelected} item={item} itemIndex={itemIndex} currentIndex={currentIndex}>
            <ListViewItemContent icon={item.icon} title={item.title} disabled={item.disabled} isSelected={isSelected}
                                 subtitle={item.subtitle}/>
        </ListViewItemCard>
    </>;
}
