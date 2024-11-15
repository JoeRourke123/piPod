import React, {useMemo} from "react";
import {ListViewItemProps} from "./list-view-item";
import {Box, Collapse} from "@chakra-ui/react";
import {useLocation} from "react-router-dom";

export const ListViewItemCard = ({ item, itemIndex, currentIndex, isSelected, children }: ListViewItemProps & { isSelected: boolean, children: React.JSX.Element }) => {
    const { key } = useLocation();
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
                backgroundImage: item.backgroundImage,
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

    const innerItemStyles = useMemo(() => {
        if (isSelected) {
            return {
                borderRadius: "lg",
                backdropFilter: "auto",
                backdropBlur: "2px",
                backdropBrightness: item.backgroundImage ? "0.8" : "1",
            };
        } else {
            return {};
        }
    }, [isSelected, key, item]);

    return <>
        <Box key={item.path} as="div" className="listViewItemButton" width="full" justifyContent="start"
             bgPosition="center"
             bgSize="cover"
             borderRadius="lg"
             backgroundImage={itemStyles.backgroundImage}
             backgroundColor={itemStyles.backgroundColor}
             boxShadow={itemStyles.boxShadow}
             textColor={itemStyles.textColor}
        >
            <Box width="full"
                 boxSize="full"
                 px="10px"
                 py="8px"
                 {...innerItemStyles}
            >
                <Collapse animateOpacity startingHeight={30} endingHeight={item.subtitle ? 52 : 30} style={{width: "100%"}} in={isSelected}>
                    { children }
                </Collapse>
            </Box>
        </Box>
        </>;
}