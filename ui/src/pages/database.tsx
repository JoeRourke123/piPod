import {PageProps} from "./page-props";
import {Box, Container, Flex, Heading, Spacer, VStack} from "@chakra-ui/react";
import {useCallback, useState} from "react";
import {Editor} from "../components/database/editor";
import {CollectionSelector} from "../components/database/collection-selector";
import {IdSelector} from "../components/database/id-selector";
import {OffsetSelector} from "../components/database/offset-selector";
import {DeleteButton} from "../components/database/delete-button";
import {RefreshButton} from "../components/database/refresh-button";

export const Database = (props: PageProps) => {
    const [currentCollection, setCurrentCollection] = useState("");
    const [currentDocument, setCurrentDocument] = useState("");
    const [offset, setOffset] = useState(0);
    const [refreshHash, setRefreshHash] = useState(0);

    const triggerRefresh = useCallback(() => {
        setRefreshHash(Math.floor(Math.random() * 1000));
    }, [refreshHash]);

    return (
        <VStack p="6">
            <Container w="full">
                <Heading my="6">PiPod Database</Heading>
                <Box my="6" borderWidth='1px' borderRadius='lg' p="4">
                    <Flex gap="4" alignItems="end">
                        <CollectionSelector setCurrentCollection={setCurrentCollection} currentCollection={currentCollection} />
                        <IdSelector documentId={currentDocument} setDocumentId={setCurrentDocument} />
                        <OffsetSelector offset={offset} setOffset={setOffset} />
                        <RefreshButton triggerRefresh={triggerRefresh} />
                        <Spacer />
                        <DeleteButton collection={currentCollection} triggerRefresh={triggerRefresh} />
                    </Flex>
                </Box>
                <Box my="6" borderWidth='1px' borderRadius='lg' p="4" overflowY="scroll" width="full" height="600px">
                    <Editor currentCollection={currentCollection} currentDocumentId={currentDocument} offset={offset} refreshHash={refreshHash} />
                </Box>
            </Container>
        </VStack>
    )
}