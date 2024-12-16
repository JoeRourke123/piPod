import {JsonEditor, OnChangeFunction, UpdateFunction} from "json-edit-react";
import {useEffect, useState} from "react";
import {useToast} from "@chakra-ui/toast";

export type DbEditorProps = {
    currentCollection: string;
    currentDocumentId: string;
    offset: number;
    refreshHash: number;
};

export const Editor = ({currentCollection, currentDocumentId, offset, refreshHash}: DbEditorProps) => {
    const [documentData, setDocumentData] = useState({
        "msg": "No documents have been found (yet)."
    });

    const toast = useToast();

    useEffect(() => {
        if (currentCollection) {
            fetch(`http://localhost:9091/db/collections/${currentCollection}?offset=${offset}&id=${currentDocumentId}`).then((response) => {
                return response.json();
            }).then((json) => {
                setDocumentData(json);
            });
        }
    }, [currentDocumentId, currentCollection, offset, refreshHash]);

    const onDocUpdated: UpdateFunction = (data) => {
        const docData = data.newData;
        fetch(`http://localhost:9091/db/collections/${currentCollection}?id=${currentDocumentId}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(docData)
        }).then((response) => {
            toast({
                title: "Changes Saved!",
                description: `${currentDocumentId} has been updated.`,
                status: "success",
                duration: 2500,
            });
        });
    }

    return (
        <JsonEditor data={documentData} onUpdate={onDocUpdated} minWidth="100%" defaultValue={"No documents have been found."} />
    )
}