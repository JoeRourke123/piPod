import {ListViewItemDetails} from "../../../util/ListViewTypes";
import {Box, Flex} from "@chakra-ui/react";
import * as React from "react";
import {useRenderItems} from "../utils/item-builder";
import {ItemLoader} from "../utils/item-loader";

export type ListViewBodyProps = {
    selectedIndex: number;
    itemLoader: ItemLoader;
    itemsHash: string;
}

export const ListViewBody = ({ selectedIndex, itemsHash, itemLoader }: ListViewBodyProps) => {
    const renderedItems = useRenderItems({itemLoader, itemsHash, selectedIndex});

    return <Box className="listView" id="list">
        <Flex align="start" gap={1} py={4} flexFlow="column" children={renderedItems} />
    </Box>;
}