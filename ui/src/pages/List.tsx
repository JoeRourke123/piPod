import * as React from "react"
import {Gear, Joystick, MicrophoneStage, MusicNotes} from "@phosphor-icons/react";
import {PageProps} from "../state/PageProps";
import {ListViewProps} from "../util/ListViewTypes";
import {ListView} from "../components/ListView";
import {useLoaderData} from "react-router-dom";

export const List = (props: PageProps) => {
    const items: any = useLoaderData();

    return ListView({ ...items, ...props })
}
