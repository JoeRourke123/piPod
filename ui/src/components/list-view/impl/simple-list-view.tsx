import {PageProps} from "../../../pages/page-props";
import {AdditionalListViewInfo, ListViewItemDetails} from "../list-view-types";
import {BaseListView} from "../base-list-view";
import {useItemsHash} from "../utils/items-hash";

type SimpleListViewProps = PageProps & {
    title: string;
    items: ListViewItemDetails[];
    showStatus: boolean;
    icon?: string;
    additionalInfo: AdditionalListViewInfo[];
};

export const SimpleListView = ({socket, title, items, showStatus, icon, additionalInfo}: SimpleListViewProps) => {
    const [itemsHash, _] = useItemsHash(items);
    const itemLoader = (currentOffset: number) => items[currentOffset];

    return <BaseListView socket={socket} title={title} showStatus={showStatus} itemsHash={itemsHash} icon={icon}
                         itemCount={items.length} itemLoader={itemLoader} additionalInfo={additionalInfo}/>
}