import React, {ForwardRefExoticComponent, useCallback, useEffect} from "react";
import {Button, Spacer, Text} from "@chakra-ui/react";
import {Gear, IconProps, Joystick} from "@phosphor-icons/react";
import {ChevronRightIcon, Icon} from "@chakra-ui/icons";
import {ListViewItemDetails} from "../../../util/ListViewTypes";

export type ListViewItemProps = {
    currentIndex: number
    itemIndex: number
    item: ListViewItemDetails
    icon?: any
};

export const ListViewItem = ({currentIndex, itemIndex, item, icon}: ListViewItemProps) => {
    if (currentIndex === itemIndex) {
        return (
            <Button className="listViewItemButton" width="100%" justifyContent="start" leftIcon={icon && icon("white")} colorScheme='cyan' variant='solid' rightIcon={<ChevronRightIcon color="white" />}>
                <Text className="listViewItem" color="white" pl={1}>{ item.title }</Text><Spacer/>
            </Button>
        );
    } else {
        return (
            <Button className="listViewItemButton" width="100%" justifyContent="start" leftIcon={icon && icon("black")} colorScheme='blackAlpha' variant='ghost'>
                <Text className="listViewItem" pl={1}>{ item.title }</Text>
            </Button>
        );
    }
}