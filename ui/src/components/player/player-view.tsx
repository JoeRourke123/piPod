import {
    Box, Button,
    Flex, IconButton,
    Image,
    Text,
    VStack
} from "@chakra-ui/react";
import React, {useContext, useEffect, useMemo, useState} from "react";
import {Airplay, SpeakerHigh} from "@phosphor-icons/react";
import {PlayerContext} from "./player-provider";
import {convertMsToTime} from "../../util/functions";
import {CurrentTrackContext} from "./current-track-context";
import { Slider } from "@chakra-ui/slider";

export const PlayerView = () => {
    const [currentTrack, _] = useContext(CurrentTrackContext);
    const playerContext = useContext(PlayerContext);

    const durationTime = useMemo(() => convertMsToTime(currentTrack?.duration || 0), [currentTrack]);
    const [positionTime, setPositionTime] = useState("0:00");

    useEffect(() => {
        playerContext.onPositionUpdated = (position: number) => setPositionTime(convertMsToTime(position));
    }, []);

    if (!currentTrack) {
        return <></>;
    }


    const image = <Image height={75} rounded="lg" id="thumbnail" srcSet={currentTrack.imageUrl} />
    const title = <VStack width="full" alignItems="start" pl="4" gap="0" color="white">
        <Text fontWeight="bold">{ currentTrack.trackName }</Text>
        <Box p="none" m="none">
            <Text>{ currentTrack.artist }</Text>
        </Box>
    </VStack>

    return <VStack height="100vh" gap="2" px="6" pt="6" width="full">
        <Flex width="100%" flexDirection="row" alignItems={"center"} justifyContent={"space-around"}>
            {image}
            {title}
        </Flex>
        <Box pt="4" width="full">
            <Flex width="full" flexDirection="column" gap="1">
                <Box width="full" height="25px">
                    <Slider size="md" defaultValue={40} width="full" zIndex={999} min={0} max={currentTrack.duration} value={playerContext.getPosition()} />
                </Box>
                <Flex width="full" flexDirection="row" justifyContent="space-between">
                    <Text fontSize="sm" color="whiteAlpha.800">{ durationTime }</Text>
                    <Text fontSize="sm" color="whiteAlpha.800">-{ positionTime }</Text>
                </Flex>
            </Flex>
        </Box>
        <Flex height="full" width="full" alignItems="start" flexDirection="row" justifyContent="space-between">
            <IconButton
                size="sm"
                variant="solid"
                aria-label="Volume"
            ><SpeakerHigh /></IconButton>
            <Button size="sm" rounded={20}>
                <Airplay /> Joe's AirPods
            </Button>
        </Flex>
    </VStack>
}