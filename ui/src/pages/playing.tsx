import {PageProps} from "./page-props";
import {Container} from "@chakra-ui/react";
import {PlayerView} from "../components/player-view";
import {usePlayerSetup} from "../hooks/use-player-setup";
import {useClickwheel} from "../hooks/use-clickwheel";
import {useNavigate} from "react-router-dom";

export const Playing = ({socket}: PageProps) => {
    const playbackState = usePlayerSetup();
    const navigate = useNavigate();

    useClickwheel({
        maxClickWheelValue: 0, socket, onMenuButton: (_: number) => {
            navigate(-1);
        }
    })

    return (
        <Container width="full" mx="none" px="none">
            {playbackState && <PlayerView playbackState={playbackState}/>}
        </Container>
    )
}