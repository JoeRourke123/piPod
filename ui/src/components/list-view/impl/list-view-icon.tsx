import * as icons from "@phosphor-icons/react";

export const ListViewIcon = ({name, colour = "blackAlpha", fontSize = 24}: {
    name: string,
    colour?: string,
    fontSize?: number
}) => {
    // @ts-ignore
    const Icon: typeof icons.Icon | undefined = icons[name];

    return <>
        {
            Icon && <Icon color={colour} fontSize={fontSize} fontWeight={800}/>
        }
    </>
}