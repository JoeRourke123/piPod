import * as React from "react"
import {Gear, Joystick, MicrophoneStage, MusicNotes} from "@phosphor-icons/react";
import {PageProps} from "../state/PageProps";
import {ListViewProps} from "../util/ListViewTypes";
import {ListView} from "../components/ListView";
import {useLoaderData, useNavigate} from "react-router-dom";
import {useEffect} from "react";

export const Home = (props: PageProps) => {
    const isAuth: any = useLoaderData();
    const navigate = useNavigate();

    useEffect(() => {
        if (!isAuth["has_token"]) {
            navigate("/auth", {
                replace: true,
                state: { authUrl: isAuth["auth_url"] }
            });
        }
    }, []);

    const items: ListViewProps = {
        title: "iPod",
        showStatus: true,
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
