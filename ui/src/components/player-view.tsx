import {
    Box, Button,
    Flex, IconButton,
    Image,
    Slider,
    SliderFilledTrack,
    SliderThumb,
    SliderTrack,
    Text,
    VStack
} from "@chakra-ui/react";
import {convertMsToTime} from "../util/functions";
import React from "react";
import {Airplay, SpeakerHigh} from "@phosphor-icons/react";

export const PlayerView = ({playbackState}: { playbackState: Spotify.PlaybackState }) => {
    const image = <Image height={75} rounded="lg" id="thumbnail" srcSet={playbackState.track_window.current_track.album.images[2].url} />
    const title = <VStack width="full" alignItems="start" pl="4" gap="0" color="white">
        <Text fontWeight="bold">{ playbackState.track_window.current_track.name }</Text>
        <Box p="none" m="none">
            <Text>{ playbackState.track_window.current_track.artists[0].name}</Text>
        </Box>
    </VStack>

    return <VStack height="100vh" gap="2" px="6" pt="6" width="full">
        <Image srcSet={playbackState.track_window.current_track.album.images[2].url} id="backgroundImage" />
        <Flex width="100%" flexDirection="row" alignItems={"center"} justifyContent={"space-around"}>
            {image}
            {title}
        </Flex>
        <Box pt="4" width="full">
            <Flex width="full" flexDirection="column" gap="1">
                <Slider focusThumbOnChange={false} colorScheme="white" aria-label='slider-ex-1' value={playbackState.position} defaultValue={1} min={0} max={playbackState.duration}>
                    <SliderTrack>
                        <SliderFilledTrack />
                    </SliderTrack>
                    <SliderThumb boxSize="8px" />
                </Slider>
                <Flex width="full" flexDirection="row" justifyContent="space-between">
                    <Text fontSize="sm" color="whiteAlpha.800">{ convertMsToTime(playbackState.duration) }</Text>
                    <Text fontSize="sm" color="whiteAlpha.800">-{ convertMsToTime(playbackState.duration - playbackState.position) }</Text>
                </Flex>
            </Flex>
        </Box>
        <Flex height="full" width="full" alignItems="start" flexDirection="row" justifyContent="space-between">
            <IconButton
                size="sm"
                isRound={true}
                variant="solid"
                aria-label="Volume"
                icon={<SpeakerHigh />}
            />
            <Button size="sm" rounded={20} leftIcon={<Airplay />}>
                Joe's AirPods
            </Button>
        </Flex>
    </VStack>
}