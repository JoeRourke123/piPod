import {Heading} from "@chakra-ui/react";
import * as React from "react";

export const ListViewTitle = ({title}: { title: string }) => {
    return <Heading size="md" as="h3">
        {title}
    </Heading>;
}