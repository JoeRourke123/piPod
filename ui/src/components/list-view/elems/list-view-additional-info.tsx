import {AdditionalListViewInfo} from "../list-view-types";
import {Box, Flex, Text, Wrap} from "@chakra-ui/react";
import {ListViewIcon} from "../impl/list-view-icon";

export const AdditionalInfoList = ({ additionalInfo }: { additionalInfo: AdditionalListViewInfo[] }) => {
    if (additionalInfo.length) {
        return <Wrap flexDirection="row" px="4px" pb="12px" spacingX={5} spacingY={1}>
            {additionalInfo.map((info, index) => <ListViewAdditionalInfo key={index} {...info} />)}
        </Wrap>;
    } else {
        return <></>;
    }
}

const ListViewAdditionalInfo = (additionalInfo: AdditionalListViewInfo) => {
    return <Flex flexDirection="row" alignItems="center" justifyContent="start">
        { additionalInfo.icon && <Box pr={1}><ListViewIcon name={additionalInfo.icon} fontSize={18} colour="blackAlpha" /></Box> }
        <Text fontSize="sm" color="blackAlpha" fontWeight={additionalInfo.bold ? 600 : 300}>{ additionalInfo.text}</Text>
    </Flex>;
}