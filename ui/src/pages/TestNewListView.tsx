import {PageProps} from "../state/PageProps";
import {BaseListView} from "../components/list-view/base-list-view";
import {ListViewItemDetails} from "../util/ListViewTypes";

export const NewListView = (props: PageProps) => {
    const itemHash = "test";
    const items: ListViewItemDetails[] = [
        {
            title: "Item 1",
            path: "/item1",
        },
        {
            title: "Item 2",
            path: "/item2",
        },
        {
            title: "Item 3",
            path: "/item3",
        },
        {
            title: "Item 4",
            path: "/item3",
        },
        {
            title: "Item 5",
            path: "/item3",
        },
        {
            title: "Item 6",
            path: "/item3",
        },
    ];

    const itemLoader = (index: number) => {
        return items[index];
    }


    return <BaseListView
        socket={props.socket}
        title={"PiPod"}
        showStatus={true}
        itemsHash={itemHash}
        itemCount={items.length}
        itemLoader={itemLoader} />
}