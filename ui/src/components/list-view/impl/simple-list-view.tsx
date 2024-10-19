import {PageProps} from "../../../pages/page-props";
import {ListViewItemDetails} from "../list-view-types";
import {BaseListView} from "../base-list-view";
import {useItemsHash} from "../utils/items-hash";

type SimpleListViewProps = PageProps & {
    title: string,
    items: ListViewItemDetails[],
    showStatus: boolean
};

export const SimpleListView = ({socket, title, items, showStatus}: SimpleListViewProps) => {
    const [itemsHash, _] = useItemsHash(items);
    const itemLoader = (currentOffset: number) => items[currentOffset];

    return <BaseListView socket={socket} title={title} showStatus={showStatus} itemsHash={itemsHash}
                         itemCount={items.length} itemLoader={itemLoader}/>
}