import {PageProps} from "../state/PageProps";
import {useLoaderData, useLocation, useParams} from "react-router-dom";
import {ListView} from "../components/ListView";
import {Heading} from "@chakra-ui/react";
import {ListViewItemProps} from "../util/ListViewTypes";
import {useEffect, useState} from "react";

export const CategoryList = (props: PageProps) => {
    const data: any = useLoaderData();
    const { key } = useLocation();
    const [basePageLoaderUrl, setBasePageLoaderUrl] = useState<string | undefined>(undefined);
    const params = useParams();

    useEffect(() => {
        if (props.basePageLoaderUrl) {
            let fullApiEndpoint = props.basePageLoaderUrl;

            if (params["type"]) {
                fullApiEndpoint += params["type"];
            } if (params["id"]) {
                fullApiEndpoint += "/" + params["id"];
            }

            setBasePageLoaderUrl(fullApiEndpoint);
        }
    }, [key]);

    const pageLoader = async (currentOffset: number): Promise<ListViewItemProps[]> => {
        if (basePageLoaderUrl) {
            const apiUrl = `http://localhost:9091${basePageLoaderUrl}?next=${currentOffset}`;
            const response = await fetch(apiUrl);

            const json = await response.json()
            return json["items"];
        } else {
            return [];
        }
    }

    return (
        <ListView
            customTitle={<Heading size="md" as="h3">{ data["title"] }</Heading>}
            showStatus={data["show_status"]}
            pageLoader={pageLoader}
            basePageLoaderUrl={basePageLoaderUrl}
            items={data["items"]}
            socket={props.socket} />
    )
}