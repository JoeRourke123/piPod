import {useLoaderData} from "react-router-dom";
import {ListViewItemDetails, ListViewProps} from "../../../util/ListViewTypes";

export const useListViewLoader = (): ListViewProps => {
    const loaderData: any = useLoaderData();

    return unmarshallView(loaderData);
}

export const unmarshallView = (responseData: any): ListViewProps => {
    return {
        title: responseData["title"] || "PiPod",
        showStatus: responseData["show_status"],
        items: unmarshallItems(responseData["items"]) || [],
    }
}

const unmarshallItems = (items: any[]): ListViewItemDetails[] | undefined => {
    if (!items || items.length === 0) {
        return undefined;
    }

    return items.map((item: any) => {
        return {
            title: item["title"],
            path: item["path"],
            requestUrl: item["request_url"],
            actionType: item["action_type"],
            actions: unmarshallItems(item["actions"]),
            toastMessage: item["toast_message"]
        }
    });
}