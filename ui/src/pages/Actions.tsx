import {PageProps} from "../state/PageProps";
import React from "react";
import {useLocation} from "react-router-dom";
import {ListViewItemDetails} from "../util/ListViewTypes";
import {SimpleListView} from "../components/list-view/implementations/simple-list-view";

export const Actions = (props: PageProps): React.JSX.Element => {
    const {state} = useLocation();
    const {actions, title}: { title: string, actions: ListViewItemDetails[] } = state;

    return <SimpleListView socket={props.socket} title={title} items={actions} showStatus />
}