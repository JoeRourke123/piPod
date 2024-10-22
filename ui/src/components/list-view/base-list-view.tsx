import {PageProps} from "../../pages/page-props";
import {Box, Flex} from "@chakra-ui/react";
import * as React from "react";
import {useClickwheel} from "../../hooks/use-clickwheel";
import {useRenderItems} from "./utils/item-builder";
import {useListViewCallbacks} from "./utils/list-view-callbacks";
import {ItemLoader} from "./utils/item-loader";
import {ListViewTitle} from "./elems/list-view-title";
import {ListViewStatusBar} from "./elems/list-view-status-bar";
import {ListViewBody} from "./elems/list-view-body";
import { ListViewHeader } from "./elems/list-view-header";
import {AnimatedLayout} from "../animated-layout";

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
        <AnimatedLayout>
            <Box fontSize="l" p={3} px={5}>
                <ListViewHeader title={title} showStatus={showStatus} />
                <ListViewBody
                    itemLoader={itemLoader}
                    itemsHash={itemsHash}
                    selectedIndex={selectedIndex}/>
            </Box>
        </AnimatedLayout>
    </>;
}