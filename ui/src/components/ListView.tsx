import {ListViewItemProps, ListViewProps} from "../util/ListViewTypes";
import {PageProps} from "../state/PageProps";
import * as React from "react";
import {useCallback, useEffect, useState} from "react";
import {useClickwheel, useSelectButton} from "../util/useClickwheel";
import {ClickWheelData} from "../util/clickwheelListeners";
import {Box, Flex, Heading, Spacer, usePrevious} from "@chakra-ui/react";
import {BatteryMedium, Headphones, WifiHigh} from "@phosphor-icons/react";
import {ListViewItem} from "./ListViewItem";
import {useBeforeUnload, useLocation, useNavigate} from "react-router-dom";
import {handleKeyUp} from "../util/handleKeyUp";

export const ListView = (props: ListViewProps & PageProps) => {
    const [listViewItems, setListViewItems] = useState<ListViewItemProps[]>(props.items);
    const [selectedIndex, setSelectedIndex] = useState(0);
    const previousIndex = usePrevious(selectedIndex);
    const [isPressed, setIsPressed] = useState(false);
    const [startPosition, setStartPosition] = useState(0);
    const [currentPosition, setCurrentPosition] = useState(0);
    const [isLoadingMore, setIsLoadingMore] = useState(false);
    const [canLoadMore, setCanLoadMore] = useState(true);
    const previousPosition = usePrevious(currentPosition);
    const navigate = useNavigate();
    const { key } = useLocation();

    const viewItemComponents = listViewItems.map((item: ListViewItemProps, index: number) => {
        return <ListViewItem key={index} currentIndex={selectedIndex} itemIndex={index} title={item.title}
                             icon={item.icon || props.fallbackIcon}/>
    })

    useClickwheel(props.socket, (cw: ClickWheelData) => {
        if (cw.isClickWheelPressed !== isPressed) {
            setIsPressed(cw.isClickWheelPressed);
        }
        setCurrentPosition(cw.clickWheelPosition);
    });

    useSelectButton(props.socket, () => {
        const currentlySelected = listViewItems[selectedIndex];
        navigate(currentlySelected.path || "./");
    });

    useEffect(() => {
        if (isPressed) {
            setStartPosition(currentPosition);
        }
    }, [isPressed]);

    useEffect( () => {
        if (currentPosition <= 46 && currentPosition >= 2) {
            const isClockwise = currentPosition > previousPosition;

            if (isClockwise && selectedIndex + 1 >= listViewItems.length) {
                return
            } else if (!isClockwise && selectedIndex - 1 < 0) {
                return
            } else if (Math.abs(currentPosition - startPosition) >= 5) {
                setStartPosition(currentPosition);
                setSelectedIndex(selectedIndex + (isClockwise ? 1 : -1));
            }
        }
    }, [currentPosition]);

    useEffect(() => {
        document.getElementById("list")?.scrollTo({left: 0, top: 0});
        setListViewItems(props.items);
        document.onkeyup = handleKeyUp(setSelectedIndex, listViewItems, navigate);
    }, [key]);

    useBeforeUnload(() => {
        document.onkeyup = null;
    })

    useEffect(() => {
        document.onkeyup = handleKeyUp(setSelectedIndex, listViewItems, navigate);
        if (previousIndex > selectedIndex) {
            document.getElementById("list")?.scrollBy({left: 0, top: -44, behavior: "smooth"});
        } else if (selectedIndex + 1 >= 4 && previousIndex < selectedIndex) {
            document.getElementById("list")?.scrollBy({left: 0, top: 44, behavior: "smooth"});
        }

        if ((listViewItems.length - selectedIndex) <= 25 && props.basePageLoaderUrl) {
            if (props.pageLoader && !isLoadingMore && canLoadMore) {
                props.pageLoader(listViewItems.length).then(async (nextItems: ListViewItemProps[]) => {
                    const items = [...listViewItems, ...nextItems];
                    setCanLoadMore(items.length !== listViewItems.length);
                    setListViewItems(items);
                    setIsLoadingMore(false);
                });
                setIsLoadingMore(true);
            }
        }
    }, [selectedIndex]);

    return (
        <Box fontSize="l" p={3} px={5}>
            <Flex align="center" className="titleBar" px={3}>
                {props.customTitle !== undefined && props.customTitle}
                {props.title !== undefined && <Heading>
                    {props.title}
                </Heading>}
                {props.showStatus && <>
                    <Spacer/>
                    <Flex gap={3}>
                        <Headphones scale={12}/>
                        <WifiHigh scale={12}/>
                        <BatteryMedium scale={12}/>
                    </Flex>
                </>}
            </Flex>
            <Box className="listView" id="list">
                <Flex align="start" gap={1} py={4} flexFlow="column">
                    {viewItemComponents}
                </Flex>
            </Box>
        </Box>
    );
}