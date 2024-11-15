import {ListViewTitle} from "./list-view-title";
import {ListViewStatusBar} from "./list-view-status-bar";
import {Box, Flex} from "@chakra-ui/react";
import * as React from "react";
import {ListViewIcon} from "../impl/list-view-icon";

type ListViewHeaderProps = {
    title: string;
    showStatus: boolean;
    icon?: string;
}

export const ListViewHeader = ({
                                   title,
                                   icon,
                                   showStatus
                               }: ListViewHeaderProps) => <Flex align="center" className="titleBar" pr={3} pl={1}>
    { icon && <Box mr={2}><ListViewIcon fontSize={18} name={icon} /></Box> }
    <ListViewTitle title={title}/>
    <ListViewStatusBar showStatus={showStatus}/>
</Flex>;