import {useLoaderData} from "react-router-dom";
import {ListViewItemDetails, ListViewProps} from "../list-view-types";

export const useListViewLoader = (): ListViewProps => {
    const loaderData: any = useLoaderData();

    return unmarshallView(loaderData);
}

export const unmarshallView = (responseData: any): ListViewProps => {
    return {
        title: responseData["title"] || "PiPod",
        icon: responseData["icon"],
        showStatus: responseData["show_status"] || responseData["showStatus"],
        items: unmarshallItems(responseData["items"]) || [],
        additionalInfo: unmarshallAdditionalInfo(responseData["additional_info"] || responseData["additionalInfo"]),
    }
}

const unmarshallAdditionalInfo = (additionalInfo: any[]) => {
    if (additionalInfo) {
        return additionalInfo.map((info: any) => {
            return {
                text: info["text"],
                icon: info["icon"],
                bold: info["bold"],
            }
        });
    }

    return [];
};

const unmarshallItems = (items: any[]): ListViewItemDetails[] | undefined => {
    if (!items || items.length === 0) {
        return undefined;
    }

    return items.map((item: any): ListViewItemDetails => {
        return {
            title: item["title"],
            subtitle: item["subtitle"],
            path: item["path"],
            requestUrl: item["request_url"] || item["requestUrl"],
            actionType: item["action_type"] || item["actionType"],
            actions: unmarshallItems(item["actions"]),
            toastMessage: item["toast_message"] || item["toastMessage"],
            icon: item["icon"],
            disabled: item["disabled"],
            backgroundImage: item["background_image"] || item["backgroundImage"],
        }
    });
}