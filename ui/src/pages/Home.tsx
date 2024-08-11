import * as React from "react"
import {Box, Button, Flex, Heading, Spacer, Text, VStack,} from "@chakra-ui/react"
import {BatteryMedium, Gear, Headphones, Joystick, MicrophoneStage, MusicNotes, WifiHigh} from "@phosphor-icons/react";
import {useEffect, useState} from "react";
import {ListViewItem} from "../components/ListViewItem";
import {PageProps} from "../state/PageProps";
import {useClickwheel} from "../util/useClickwheel";

export const Home = (props: PageProps) => {
    const maxItems = 4;
    const [selectedIndex, setSelectedIndex] = useState(0);

    const _ = useClickwheel(maxItems, props.socket);

    useEffect(() => {
        const handleKeyDown = (e: KeyboardEvent) => {
            const key = e.key;

            if (key === "ArrowUp") {
                setSelectedIndex((index) => index === 0 ? maxItems - 1 : (index - 1) % maxItems);
            } else if (key === "ArrowDown") {
                setSelectedIndex((index) => (index + 1) % maxItems);
            }
        };

        document.addEventListener("keyup", handleKeyDown, false);

        return () => {
            document.removeEventListener('keyup', handleKeyDown);
        };
    }, []);

    useEffect(() => {
        if (selectedIndex + 1 >= maxItems) {
            document.getElementById("list")?.scrollBy(0, 40);
        }
    }, [selectedIndex]);

    return (
        <Box fontSize="l" p={3} px={5}>
            <Flex align="center" px={3}>
                <Heading>
                    iPod
                </Heading>
                <Spacer/>
                <Flex gap={3}>
                    <Headphones scale={12}/>
                    <WifiHigh scale={12}/>
                    <BatteryMedium scale={12}/>
                </Flex>
            </Flex>
            <Box className="listView" id="list">
                <Flex align="start" gap={3} py={4} flexFlow="column">
                    <ListViewItem currentIndex={selectedIndex} itemIndex={0} title={"Music"} icon={(c: string) => <MusicNotes color={c} scale={12} />} />
                    <ListViewItem currentIndex={selectedIndex} itemIndex={1} title={"Podcasts"} icon={(c: string) => <MicrophoneStage color={c} scale={12} />} />
                    <ListViewItem currentIndex={selectedIndex} itemIndex={2} title={"Games"} icon={(c: string) => <Joystick color={c} scale={12} />} />
                    <ListViewItem currentIndex={selectedIndex} itemIndex={3} title={"Settings"} icon={(c: string) => <Gear color={c} scale={12} />} />
                </Flex>
            </Box>
        </Box>
    );
}
