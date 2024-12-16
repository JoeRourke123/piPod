import React, {useMemo} from "react";
import {ListViewItemProps} from "./list-view-item";
import {Box, Collapsible} from "@chakra-ui/react";
import {useLocation} from "react-router-dom";

export const ListViewItemCard = ({item, itemIndex, currentIndex, isSelected, children}: ListViewItemProps & {
    isSelected: boolean,
    children: React.JSX.Element
}) => {
    const {key} = useLocation();
    const itemStyles = useMemo(() => {
        if (isSelected) {
            if (item.disabled) {
                return {
                    backgroundImage: "none",
                    backgroundColor: "gray.400",
                    boxShadow: "none",
                    textColor: "white",
                };
            }

            return {
                backgroundImage: `url('${item.backgroundImage}')`,
                backgroundColor: item.backgroundImage ? "none" : "cyan.400",
                boxShadow: "md",
                textColor: "white",
            }
        } else {
            return {
                backgroundImage: "none",
                boxShadow: "none",
                backgroundColor: "white",
                textColor: "blackAlpha",
            };
        }
    }, [currentIndex, itemIndex, key, item]);

    const innerBackdropFilter = useMemo(() => {
        return isSelected && item.backgroundImage ? "brightness(0.75) blur(2px)" : "none";
    }, [currentIndex, itemIndex, key, item]);

    return <>
        <Box key={item.path} as="div" className="listViewItemButton" width="full" justifyContent="start"
             backgroundPosition="center"
             backgroundSize="cover"
             borderRadius="lg"
             backgroundImage={itemStyles.backgroundImage}
             backgroundColor={itemStyles.backgroundColor}
             boxShadow={itemStyles.boxShadow}
             color={itemStyles.textColor}
        >
            <Box width="full"
                 boxSize="full"
                 px="10px"
                 py="8px"
                 borderRadius="lg"
                 backdropFilter={innerBackdropFilter}
            >
                <Collapsible.Root animateopacity="true" startingheight={30} endingheight={item.subtitle ? 52 : 30}
                                  style={{width: "100%"}} open={isSelected}>
                    {children}
                </Collapsible.Root>
            </Box>
        </Box>
    </>;
}