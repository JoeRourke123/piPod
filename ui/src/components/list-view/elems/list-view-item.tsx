import {Box, Collapse, Fade, Flex, Spacer, Text} from "@chakra-ui/react";
import {ListViewItemDetails} from "../list-view-types";
import React, {useMemo} from "react";
import {ChevronRightIcon} from "@chakra-ui/icons";
import Marquee from "react-fast-marquee";

export type ListViewItemProps = {
    currentIndex: number
    itemIndex: number
    item: ListViewItemDetails
};

export const ListViewItem = ({currentIndex, itemIndex, item}: ListViewItemProps) => {
    const isSelected = useMemo(() => {
        return currentIndex === itemIndex;
    }, [currentIndex, itemIndex]);

    const backgroundColor = isSelected ? "cyan.400" : "white";
    const textColor = isSelected ? "white" : "blackAlpha";
    const boxShadow = isSelected ? "md" : "none";

    return <>
        <Box as="div" className="listViewItemButton" width="full" justifyContent="start"
             bg={backgroundColor}
             color={textColor}
             boxShadow={boxShadow}
             borderRadius="lg"
             px="10px"
             py="8px"
        >
            <Collapse startingHeight={30} style={{width: "100%"}} in={isSelected}>
                <ListViewItemContent title={item.title} isSelected={isSelected} subtitle={item.subtitle}/>
            </Collapse>
        </Box>
    </>;
}

const ListViewItemContent = ({title, isSelected, subtitle}: { title: string, isSelected: boolean, subtitle?: string }) => {
    if (isSelected) {
        return <>
            <Flex flexDirection="row" justifyContent="space-between" width="fll" alignItems="center">
                <Box width="calc(100% - 24px)">
                    <Marquee delay={1} play={title.length >= 26} loop={1}>
                        <Text className="listViewItem" fontSize={18} pl={1} fontWeight="600">{title}</Text>{ isSelected ? <Box width={12}><span> </span></Box> : <></>}
                    </Marquee>
                    { subtitle &&
                        <Box pl={1}>
                            <Text>{ subtitle }</Text>
                        </Box>}
                </Box>
                <Spacer />
                <Box width="24px" height="24px">
                    <ChevronRightIcon color="white"/>
                </Box>
            </Flex>
        </>;
    } else {
        return <>
            <Flex flexDirection="row" justifyContent="space-between" width="full" alignItems="center">
                <Box width="calc(100% - 24px)">
                    <Text className="listViewItem" fontSize={18} pl={1}>
                        {title}
                    </Text>
                    { subtitle &&
                        <Box pl={1}>
                            <Text>{ subtitle }</Text>
                        </Box>}
                </Box>
            </Flex>
        </>;
    }
}