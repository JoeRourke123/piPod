import {Box, Button} from "@chakra-ui/react";
import {Trash} from "@phosphor-icons/react";

type DeleteButtonProps = {
    collection: string;
    triggerRefresh: () => void;
}

export const DeleteButton = ({collection, triggerRefresh}: DeleteButtonProps) => {
    const handleClick = () => {
        fetch(`http://localhost:9091/db/collections/${collection}`, {
            method: "DELETE"
        }).then(() => {
            triggerRefresh();
        });
    }

    return (
        <Box>
            <Button colorPalette="red" onClick={handleClick}>
                <Trash />
            </Button>
        </Box>
    )
}