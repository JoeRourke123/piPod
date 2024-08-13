import {PageProps} from "../state/PageProps";
import {useLoaderData, useParams} from "react-router-dom";
import {ListViewProps} from "../util/ListViewTypes";
import {ListView} from "../components/ListView";

export const CategoryList = (props: PageProps) => {
    // @ts-ignore
    const data: ListViewProps = useLoaderData();

    return (
        <ListView title={data.title} items={data.items} socket={props.socket} />
    )
}