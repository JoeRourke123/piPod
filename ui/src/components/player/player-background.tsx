import {useContext} from "react";
import {CurrentTrackContext} from "./current-track-context";
import {Box} from "@chakra-ui/react";

export const PlayerBackground = () => {
    const [currentTrack, _] = useContext(CurrentTrackContext);

    return <Box id="backgroundImage" backgroundImage={`url(${currentTrack?.imageUrl})`} backgroundSize="cover" backgroundPosition="center" position="fixed" top="0" left="0" width="100%" height="100%" zIndex={-1} />
}