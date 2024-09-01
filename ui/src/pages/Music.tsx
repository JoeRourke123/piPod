import {PageProps} from "../state/PageProps";
import {ListView} from "../components/ListView";
import {ListViewProps} from "../util/ListViewTypes";
import {Guitar, MusicNote, Playlist, UserFocus, VinylRecord} from "@phosphor-icons/react";

export const Music = (props: PageProps) => {
    const items: ListViewProps = {
        title: "Music",
        showStatus: true,
        items: [
            {
                title: "Albums",
                icon: (c: string) => <VinylRecord color={c} scale={12} />,
                path: "/list/albums"
            },
            {
                title: "Playlists",
                icon: (c: string) => <Playlist color={c} scale={12} />,
                path: "/list/playlists"
            },
            {
                title: "Artists",
                icon: (c: string) => <Guitar color={c} scale={12} />,
                path: "/list/artists"
            },
            {
                title: "Songs",
                icon: (c: string) => <MusicNote color={c} scale={12} />,
                path: "list/songs"
            },
            {
                title: "Made For You",
                icon: (c: string) => <UserFocus color={c} scale={12} />,
                path: "list/madeforyou"
            }
        ]
    }

    return ListView({...items, ...props});
}