import {useLocation, useNavigate} from "react-router-dom";
import {ClickWheelResponse} from "../../../hooks/use-clickwheel";
import {useCallback, useEffect} from "react";
import {ItemLoader} from "./item-loader";
import {useItemActionToasts} from "./item-action-toasts";
import {usePlayerDevice} from "react-spotify-web-playback-sdk";

type ListViewCallbacksProps = {
    itemLoader: ItemLoader;
    cw: Omit<ClickWheelResponse, "selectedIndex">;
    selectedIndex: number;
    itemsHash: string;
    onSelectedIndexChange?: (index: number) => void;
}

export const useListViewCallbacks = ({
                                         selectedIndex,
                                         onSelectedIndexChange,
                                         itemLoader,
                                         itemsHash,
                                         cw: {setOnMenuButton, setOnSelectButtonLongPress, setOnSelectButton,}
                                     }: ListViewCallbacksProps) => {
    const [errorToast, successToast] = useItemActionToasts();
    const navigate = useNavigate();
    const {key, pathname} = useLocation();
    const device = usePlayerDevice();

    const onSelectButton = useCallback((currentIndex: number) => {
        const currentItem = itemLoader(currentIndex);

        if (currentItem.path) {
            navigate(currentItem.path);
        } else if (currentItem.actionType === "GET" && currentItem.requestUrl) {
            fetch(currentItem.requestUrl, {
                method: "GET"
            }).then(r => r.status !== 200 ? errorToast() : successToast(currentItem.toastMessage));
        } else if (currentItem.actionType === "POST" && currentItem.requestUrl) {
            fetch(currentItem.requestUrl, {
                headers: {
                    "Content-Type": "application/json"
                },
                method: "POST",
                body: JSON.stringify({
                    "device_id": device?.device_id
                })
            }).then(r => r.status !== 200 ? errorToast() : successToast(currentItem.toastMessage));
        }
    }, [itemsHash, itemLoader, key]);

    const onMenuButton = useCallback((_: number) => {
        localStorage.removeItem(key + "-selectedIndex");
        navigate(-1);
    }, [itemsHash, itemLoader, key]);

    const onSelectButtonLongPress = useCallback((currentIndex: number) => {
        const currentItem = itemLoader(currentIndex);
        if (currentItem.actions) {
            navigate("/actions", {
                state: {
                    returnToIndex: currentIndex,
                    previousPath: pathname,
                    trackTitle: currentItem.title,
                    actions: currentItem.actions
                }
            });
        }
    }, [itemsHash, itemLoader, key]);

    useEffect(() => {
        setOnSelectButton(() => onSelectButton);
        setOnMenuButton(() => onMenuButton);
        setOnSelectButtonLongPress(() => onSelectButtonLongPress);
    }, [key, itemsHash]);

    useEffect(() => {
        const listViewItem = document.querySelectorAll("div.listViewItemButton")[selectedIndex] as HTMLElement;
        document.getElementById("list")?.scrollTo({left: 0, top: listViewItem.offsetTop - 81, behavior: "smooth"});
        if (onSelectedIndexChange) {
            onSelectedIndexChange(selectedIndex);
        }
    }, [selectedIndex]);
}