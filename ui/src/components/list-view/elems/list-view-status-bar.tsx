import {Flex, Spacer} from "@chakra-ui/react";
import {BatteryMedium, Headphones, WifiHigh} from "@phosphor-icons/react";
import * as React from "react";
import {useContext} from "react";
import {StatusContext} from "../../status-context-provider";

export const ListViewStatusBar = ({showStatus}: { showStatus: boolean }) => {
    const status = useContext(StatusContext);

    return <>
        <Spacer/>
        { showStatus && <Flex gap={3}>
            <Headphones scale={12}/>
            { status.isInternetEnabled && <WifiHigh scale={12}/>}
            <BatteryMedium scale={12}/>
        </Flex> }
    </>;
}