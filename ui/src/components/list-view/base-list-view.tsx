import {PageProps} from "../../pages/page-props";
import {Box, Flex} from "@chakra-ui/react";
import * as React from "react";
import {useClickwheel} from "../../util/useClickwheel";
import {useRenderItems} from "./utils/item-builder";
import {useListViewCallbacks} from "./utils/list-view-callbacks";
import {ItemLoader} from "./utils/item-loader";
import {ListViewTitle} from "./sub-components/list-view-title";
import {ListViewStatusBar} from "./sub-components/list-view-status-bar";
import {ListViewBody} from "./sub-components/list-view-body";
import { ListViewHeader } from "./sub-components/list-view-header";

type BaseListViewProps = PageProps & {
    title: string;
    showStatus: boolean;
    itemsHash: string;
    itemCount: number;
    itemLoader: ItemLoader;

    onSelectedIndexChange?: (index: number) => void;
};

export const BaseListView = ({
                                 title,
                                 itemLoader,
                                 itemCount,
                                 showStatus,
                                 itemsHash,
                                 socket,
                                 onSelectedIndexChange,
                             }: BaseListViewProps) => {
    const {selectedIndex, ...cw} = useClickwheel({
        socket,
        maxClickWheelValue: itemCount,
    });

    useListViewCallbacks({
        cw,
        itemLoader,
        selectedIndex,
        onSelectedIndexChange,
        itemsHash,
    });

    return <>
        <Box fontSize="l" p={3} px={5}>
            <ListViewHeader title={title} showStatus={showStatus} />
            <ListViewBody
                itemLoader={itemLoader}
                itemsHash={itemsHash}
                selectedIndex={selectedIndex}/>
        </Box>
    </>;
}