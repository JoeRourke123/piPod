import {ListViewItemDetails} from "../list-view-types";
import React, {useMemo} from "react";
import {Box, Flex, Text} from "@chakra-ui/react";
import Marquee from "react-fast-marquee";
import {ListViewIcon} from "../impl/list-view-icon";
import {CaretRight} from "@phosphor-icons/react";

export const ListViewItemContent = ({title, isSelected, subtitle, icon, disabled}: ListViewItemDetails & { isSelected: boolean }) => {
    const iconColour = isSelected ? "white" : "blackAlpha";

    const titleContent = useMemo(() => {
        if (title.length >= 26) {
            return <Box width="full" height="full">
                <Marquee delay={1} loop={1} style={{width: "100%", height: "28px"}}>
                    <Text className="listViewItem" fontSize={18} pl={1}
                          fontWeight="600">{title}</Text>
                    <Box width={12}><span> </span></Box>
                </Marquee>
            </Box>;
        } else {
            return <Box width="full" height="full">
                <Text className="listViewItem" fontSize={18} pl={1}
                      fontWeight="600">{title}</Text>
            </Box>
        }
    }, [isSelected, title]);

    if (isSelected) {
        return <>
            <Flex flexDirection="row" justifyContent="space-between" width="full" alignItems="center">
                <Flex flexDirection="row" alignItems="center" width="full">
                    {icon && <Box pr={2} pl={1}><ListViewIcon name={icon} fontSize={16} colour={iconColour}/></Box>}
                    <Box>
                        {titleContent}
                        {subtitle &&
                            <Box pl={1}>
                                <Text>{subtitle}</Text>
                            </Box>}
                    </Box>
                </Flex>
                <CaretRight color="white" fontSize="24px"/>
            </Flex>
        </>;
    } else {
        return <>
            <Flex flexDirection="row" justifyContent="start" width="full" alignItems="center">
                {!disabled && icon && <Box pr={2} pl={1}><ListViewIcon name={icon} fontSize={16} colour={iconColour}/></Box>}
                <Box>
                    <Text className="listViewItem" fontSize={18} pl={1}>
                        {title}
                    </Text>
                    {subtitle &&
                        <Box pl={1}>
                            <Text>{subtitle}</Text>
                        </Box>}
                </Box>
            </Flex>
        </>;
    }
}