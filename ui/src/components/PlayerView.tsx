import {
    Box,
    Flex,
    Heading,
    IconButton,
    Image,
    Slider,
    SliderFilledTrack,
    SliderThumb,
    SliderTrack,
    VStack
} from "@chakra-ui/react";
import {DotsThree} from "@phosphor-icons/react";

export const PlayerView = ({playbackState}: { playbackState: Spotify.PlaybackState }) => {
    const image = <Image height={120} rounded="lg" boxShadow="dark-lg" srcSet={playbackState.track_window.current_track.album.images[2].url} />
    const title = <VStack width="full" alignItems="start" gap="0" pl="4">
        <Heading h="6" size="sm" mb="10px">{ playbackState.track_window.current_track.name }</Heading>
        <Box>
            <b style={{fontSize: "14px"}}>{ playbackState.track_window.current_track.artists[0].name}</b>
            <p style={{fontSize: "14px"}}>{playbackState.track_window.current_track.album.name}</p>
        </Box>
    </VStack>

    return <VStack height="100vh" px="2" pt="5">
        <Flex height="full" width="100%" flexDirection="row" alignItems={"center"} justifyContent={"space-around"}>
            {image}
            {title}
        </Flex>
        <Box height="80px" width="100%">
            <Flex width="full">
                    <Slider focusThumbOnChange={false} colorScheme="cyan" aria-label='slider-ex-1' value={playbackState.position} defaultValue={1} min={0} max={playbackState.duration}>
                        <SliderTrack>
                            <SliderFilledTrack />
                        </SliderTrack>
                        <SliderThumb boxSize="8px" />
                    </Slider>
                <Box pl="4">
                    <IconButton
                        colorScheme="cyan"
                        variant="ghost"
                        size="sm"
                        aria-label="more"
                        icon={<DotsThree scale={6} color="white" weight="bold" />}
                    />
                </Box>
            </Flex>
        </Box>
    </VStack>
}