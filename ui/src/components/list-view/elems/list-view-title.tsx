import {Box, Text} from "@chakra-ui/react";
import * as React from "react";
import Marquee from "react-fast-marquee";
import {useMemo} from "react";

export const ListViewTitle = ({title}: { title: string }) => {
    const shouldScroll = useMemo(() => title !== undefined && title.length >= 26, [title]);

    return <Marquee delay={3} play={shouldScroll} loop={2}>
        <Text fontSize="lg" fontWeight="700">{title}</Text>{ shouldScroll ? <Box width={8}><span> </span></Box> : <></>}
    </Marquee>;
}