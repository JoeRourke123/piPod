import {AdditionalListViewInfo, ListViewItemDetails} from "../list-view-types";
import {Box, Flex} from "@chakra-ui/react";
import * as React from "react";
import {useRenderItems} from "../utils/item-builder";
import {ItemLoader} from "../utils/item-loader";

export type ListViewBodyProps = {
    selectedIndex: number;
    itemLoader: ItemLoader;
    additionalInfo: AdditionalListViewInfo[];
    itemsHash: string;
}

export const ListViewBody = ({ selectedIndex, itemsHash, itemLoader, additionalInfo }: ListViewBodyProps) => {
    const renderedItems = useRenderItems({
        itemLoader,
        itemsHash,
        selectedIndex,
        additionalInfo
    });

    return <Box className="listView" id="list">
        <Flex align="start" gap={2} pt={3} pb={4} flexFlow="column" children={renderedItems} />
    </Box>;
}