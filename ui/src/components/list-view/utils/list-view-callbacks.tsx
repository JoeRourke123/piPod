import {useLocation, useNavigate} from "react-router-dom";
import {ClickWheelResponse} from "../../../hooks/use-clickwheel";
import {useCallback, useContext, useEffect} from "react";
import {ItemLoader} from "./item-loader";
import {useItemActionToasts} from "./item-action-toasts";
import {usePlayerDevice} from "react-spotify-web-playback-sdk";
import {PipodCache} from "../../../util/pipod-cache";
import {CurrentTrackContext} from "../../player/current-track-context";

type ListViewCallbacksProps = {
    itemLoader: ItemLoader;
    itemCount: number;
    cw: Omit<ClickWheelResponse, "selectedIndex">;
    selectedIndex: number;
    itemsHash: string;
    hasAdditionalInfo: boolean;
    onSelectedIndexChange?: (index: number) => void;
}

export const useListViewCallbacks = ({
                                         selectedIndex,
                                         onSelectedIndexChange,
                                         itemLoader,
                                         hasAdditionalInfo,
                                         itemsHash,
                                         itemCount,
                                         cw: {setOnMenuButton, setOnSelectButtonLongPress, setOnSelectButton,}
                                     }: ListViewCallbacksProps) => {
    const [errorToast, successToast] = useItemActionToasts();
    const navigate = useNavigate();
    const {key, pathname} = useLocation();
    const device = usePlayerDevice();
    const [currentTrack, _] = useContext(CurrentTrackContext);

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
                    "deviceId": device?.device_id,
                    "albumId": currentTrack?.albumId,
                })
            }).then(r => r.status !== 200 ? errorToast() : successToast(currentItem.toastMessage));
        }
    }, [itemsHash, itemLoader, key]);

    const onMenuButton = useCallback((_: number) => {
        PipodCache.clear(key);
        navigate(-1);
    }, [itemsHash, itemLoader, key]);

    const onSelectButtonLongPress = useCallback((currentIndex: number) => {
        const currentItem = itemLoader(currentIndex);
        if (currentItem.actions) {
            navigate("/actions", {
                state: {
                    returnToIndex: currentIndex,
                    previousPath: pathname,
                    title: currentItem.title,
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
        const scrollOffset = hasAdditionalInfo ? 55 : 81;
        const listViewItem = document.querySelectorAll("div.listViewItemButton")[selectedIndex] as HTMLElement;
        const list = document.getElementById("list");
        if (list && listViewItem) {
            list.scrollTo({left: 0, top: listViewItem.offsetTop - scrollOffset, behavior: "smooth"});
            if (selectedIndex === 0) {
                list.scrollTo({left: 0, top: 0, behavior: "smooth"});
            } else if (selectedIndex === itemCount - 1) {
                // Jank setTimeout here to wait for the card to expand
                setTimeout(() => list.scrollTo({left: 0, top: list.scrollHeight, behavior: "smooth"}), 80);
            }
        }
        if (onSelectedIndexChange) {
            onSelectedIndexChange(selectedIndex);
        }
    }, [selectedIndex]);
}