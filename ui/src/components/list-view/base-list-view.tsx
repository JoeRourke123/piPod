import {PageProps} from "../../pages/page-props";
import {Box} from "@chakra-ui/react";
import * as React from "react";
import {useClickwheel} from "../../hooks/use-clickwheel";
import {useListViewCallbacks} from "./utils/list-view-callbacks";
import {ItemLoader} from "./utils/item-loader";
import {ListViewBody} from "./elems/list-view-body";
import {ListViewHeader} from "./elems/list-view-header";
import {AnimatedLayout} from "../animated-layout";
import {AdditionalListViewInfo} from "./list-view-types";

type BaseListViewProps = PageProps & {
    title: string;
    additionalInfo: AdditionalListViewInfo[];
    showStatus: boolean;
    icon?: string;
    itemsHash: string;
    itemCount: number;
    itemLoader: ItemLoader;

    onSelectedIndexChange?: (index: number) => void;
};

export const BaseListView = ({
                                 title,
                                 icon,
                                 itemLoader,
                                 itemCount,
                                 showStatus,
                                 itemsHash,
                                 additionalInfo,
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
        itemCount,
        hasAdditionalInfo: additionalInfo.length > 0,
    });

    return <>
        <AnimatedLayout>
            <Box fontSize="l" p={3} px={5}>
                <ListViewHeader icon={icon} title={title} showStatus={showStatus}/>
                <ListViewBody
                    itemLoader={itemLoader}
                    itemsHash={itemsHash}
                    additionalInfo={additionalInfo}
                    selectedIndex={selectedIndex}/>
            </Box>
        </AnimatedLayout>
    </>;
}