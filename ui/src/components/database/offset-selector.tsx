import {Box, NumberInput, Text} from "@chakra-ui/react";

type OffsetSelectorProps = {
    offset: number;
    setOffset: (offset: number) => void;
}

export const OffsetSelector = ({offset, setOffset}: OffsetSelectorProps) => {
    return (
        <Box>
            <Text fontSize="sm" pb="1" pl="0.5" fontWeight="bold">Offset</Text>
            <NumberInput.Root value={offset} onValueChange={(e: any) => setOffset(e.value)}>
                <NumberInput.Input />
                <NumberInput.Control>
                    <NumberInput.IncrementTrigger />
                    <NumberInput.DecrementTrigger />
                </NumberInput.Control>
            </NumberInput.Root>
        </Box>
    );
}