import {useAsync} from "react-use";
import React, {useMemo} from "react";
import {Box, NativeSelectField, NativeSelectRoot, Text} from "@chakra-ui/react";

type CollectionSelectorProps = {
    currentCollection: string;
    setCurrentCollection: (collection: string) => void;
}

export const CollectionSelector = ({currentCollection, setCurrentCollection}: CollectionSelectorProps) => {
    const state = useAsync(async (): Promise<string[]> => {
        const response = await fetch(`http://localhost:9091/db/collections`)
        return await response.json();
    }, [])

    const collections = useMemo(() => {
        return state.value || [];
    }, [state.value]);

    return (
        <Box>
            <Text fontSize="sm" pb="1" fontWeight="bold">Collection</Text>
            <NativeSelectRoot width="240px">
                <NativeSelectField
                    placeholder="Select a collection"
                    value={currentCollection}
                    onChange={(e) => setCurrentCollection(e.currentTarget.value)}
                >
                    {
                        collections.map((collection) => {
                            return (
                                <option value={collection}>{collection}</option>
                            )
                        })
                    }
                </NativeSelectField>
            </NativeSelectRoot>
        </Box>
    )
}