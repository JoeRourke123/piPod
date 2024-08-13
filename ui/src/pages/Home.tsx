import * as React from "react"
import {Gear, Joystick, MicrophoneStage, MusicNotes} from "@phosphor-icons/react";
import {PageProps} from "../state/PageProps";
import {ListViewProps} from "../util/ListViewTypes";
import {ListView} from "../components/ListView";

export const Home = (props: PageProps) => {
    const items: ListViewProps = {
        title: "iPod",
        items: [
            {
                title: "Music",
                icon: (c: string) => <MusicNotes scale={12} color={c}/>,
                path: "/music"
            },
            {
                title: "Podcasts",
                icon: (c: string) => <MicrophoneStage scale={12} color={c}/>,
                path: "/podcasts"
            },
            {
                title: "Games",
                icon: (c: string) => <Joystick scale={12} color={c}/>,
                path: "/games"
            },
            {
                title: "Settings",
                icon: (c: string) => <Gear scale={12} color={c}/>,
                path: "/settings"
            }
        ]
    }

    return ListView({ ...items, ...props })
}
