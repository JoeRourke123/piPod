import {Box, Button} from "@chakra-ui/react";
import {ArrowsClockwise, Trash} from "@phosphor-icons/react";

type RefreshButtonProps = {
    triggerRefresh: () => void;
}

export const RefreshButton = ({triggerRefresh}: RefreshButtonProps) => {
    return (
        <Box>
            <Button colorPalette="blue" onClick={triggerRefresh}>
                <ArrowsClockwise />
            </Button>
        </Box>
    )
}