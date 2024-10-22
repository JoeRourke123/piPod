import { motion } from "framer-motion";
import React, {ReactNode} from "react";

type AnimatedLayoutProps = {
    children: React.JSX.Element;
}

// I want a fade in bottom-up - fade out top-down animation
// so these are my variants
const variants = {
    hidden: { opacity: 0, x: 200, y: 0 },
    enter: { opacity: 1, x: 0, y: 0 },
    exit: { opacity: 0, x: -200, y: 0 }
}

export const AnimatedLayout = ({ children }: AnimatedLayoutProps): React.JSX.Element => {
    // return (
        // <motion.div
        //     initial="hidden"
        //     animate="enter"
        //     exit="exit"
        //     variants={variants}
        //     transition={{duration: 0.75, type: "linear"}}
        //     className="relative"
        // >
        // {children}
        // </motion.div>
    // );
    return <> { children } </>;
};