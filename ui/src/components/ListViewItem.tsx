import React, {ForwardRefExoticComponent} from "react";
import {Button, Spacer, Text} from "@chakra-ui/react";
import {Gear, IconProps, Joystick} from "@phosphor-icons/react";
import {ChevronRightIcon, Icon} from "@chakra-ui/icons";

export type ListViewItemProps = {
    currentIndex: number
    itemIndex: number
    title: string
    icon: any
};

export const ListViewItem = ({currentIndex, itemIndex, title, icon}: ListViewItemProps) => {
    if (currentIndex === itemIndex) {
        return (
            <Button width="100%" justifyContent="start" leftIcon={icon("white")} colorScheme='cyan' variant='solid' rightIcon={<ChevronRightIcon color="white" />}>
                <Text color="white" pl={4}>{ title }</Text><Spacer/>
            </Button>
        );
    } else {
        return (
            <Button width="100%" justifyContent="start" leftIcon={icon("black")} colorScheme='blackAlpha' variant='ghost'>
                <Text pl={4}>{ title }</Text>
            </Button>
        );
    }
}