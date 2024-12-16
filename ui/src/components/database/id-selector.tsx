import {Box, Input, Text} from "@chakra-ui/react";
import {ChangeEvent, ChangeEventHandler} from "react";

type IdSelectorProps = {
    documentId: string;
    setDocumentId: (documentId: string) => void;
}

export const IdSelector = ({ documentId, setDocumentId }: IdSelectorProps) => {
    const handleChange: ChangeEventHandler<HTMLInputElement> = (event: ChangeEvent<HTMLInputElement>) => {
        setDocumentId(event.target.value);
    }

    return (
        <Box>
            <Text fontSize="sm" pb="1" pl="0.5" fontWeight="bold">Document ID</Text>
            <Input placeholder="Enter a document ID" onChange={handleChange} />
        </Box>
    );
}