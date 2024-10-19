import {ListViewTitle} from "./list-view-title";
import {ListViewStatusBar} from "./list-view-status-bar";
import {Flex} from "@chakra-ui/react";
import * as React from "react";

type ListViewHeaderProps = {
    title: string;
    showStatus: boolean;
}

export const ListViewHeader = ({
                            title,
                            showStatus
}: ListViewHeaderProps) => <Flex align="center" className="titleBar" px={3}>
    <ListViewTitle title={title}/>
    <ListViewStatusBar showStatus={showStatus}/>
</Flex>;