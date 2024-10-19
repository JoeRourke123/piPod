import {Flex, Spacer} from "@chakra-ui/react";
import {BatteryMedium, Headphones, WifiHigh} from "@phosphor-icons/react";
import * as React from "react";

export const ListViewStatusBar = ({showStatus}: { showStatus: boolean }) => {
    return <>
        <Spacer/>
        { showStatus && <Flex gap={3}>
            <Headphones scale={12}/>
            <WifiHigh scale={12}/>
            <BatteryMedium scale={12}/>
        </Flex> }
    </>;
}