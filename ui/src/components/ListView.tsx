import {ListViewItemProps, ListViewProps} from "../util/ListViewTypes";
import {PageProps} from "../state/PageProps";
import * as React from "react";
import {useEffect, useState} from "react";
import {useClickwheel, useSelectButton} from "../util/useClickwheel";
import {ClickWheelData} from "../util/clickwheelListeners";
import {Box, Flex, Heading, Spacer, usePrevious} from "@chakra-ui/react";
import {BatteryMedium, Headphones, WifiHigh} from "@phosphor-icons/react";
import {ListViewItem} from "./ListViewItem";
import {useNavigate} from "react-router-dom";

export const ListView = (props: ListViewProps & PageProps) => {
    const maxItems = props.items.length;
    const [selectedIndex, setSelectedIndex] = useState(0);
    const previousIndex = usePrevious(selectedIndex);
    const [isPressed, setIsPressed] = useState(false);
    const [startPosition, setStartPosition] = useState(0);
    const [currentPosition, setCurrentPosition] = useState(0);
    const previousPosition = usePrevious(currentPosition);
    const navigate = useNavigate();

    const viewItemComponents = props.items.map((item: ListViewItemProps, index: number) => {
        return <ListViewItem currentIndex={selectedIndex} itemIndex={index} title={item.title}
                             icon={item.icon || props.fallbackIcon}/>
    })

    useClickwheel(props.socket, (cw: ClickWheelData) => {
        if (cw.isClickWheelPressed !== isPressed) {
            setIsPressed(cw.isClickWheelPressed);
        }
        setCurrentPosition(cw.clickWheelPosition);
    });

    useSelectButton(props.socket, () => {
        const currentlySelected = props.items[selectedIndex];
        navigate(currentlySelected.path || "./");
    });

    useEffect(() => {
        if (isPressed) {
            setStartPosition(currentPosition);
        }
    }, [isPressed]);

    useEffect(() => {
        const isClockwise = currentPosition > previousPosition;

        if (Math.abs(currentPosition - startPosition) >= 5) {
            setStartPosition(currentPosition);
            setSelectedIndex(selectedIndex + (isClockwise ? 1 : -1));
        }
    }, [currentPosition]);

    useEffect(() => {
        const handleKeyDown = (e: KeyboardEvent) => {
            const key = e.key;

            if (key === "ArrowUp") {
                setSelectedIndex((index) => Math.max(0, index - 1));
            } else if (key === "ArrowDown") {
                setSelectedIndex((index) => Math.min((index + 1), maxItems - 1));
            }
        };

        document.addEventListener("keyup", handleKeyDown, false);

        return () => {
            document.removeEventListener('keyup', handleKeyDown);
        };
    }, []);

    useEffect(() => {
        console.log(selectedIndex, maxItems);
        if (previousIndex > selectedIndex) {
            document.getElementById("list")?.scrollBy({left: 0, top: -45, behavior: "smooth"});
        } else if (selectedIndex + 1 >= 4 && previousIndex < selectedIndex) {
            document.getElementById("list")?.scrollBy({left: 0, top: 45, behavior: "smooth"});
        }
    }, [selectedIndex]);

    return (
        <Box fontSize="l" p={3} px={5}>
            <Flex align="center" className="titleBar" px={3}>
                <Heading>
                    {props.title}
                </Heading>
                <Spacer/>
                <Flex gap={3}>
                    <Headphones scale={12}/>
                    <WifiHigh scale={12}/>
                    <BatteryMedium scale={12}/>
                </Flex>
            </Flex>
            <Box className="listView" id="list">
                <Flex align="start" gap={1} py={4} flexFlow="column">
                    {viewItemComponents}
                </Flex>
            </Box>
        </Box>
    );
}