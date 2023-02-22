import React from "react";
import { motion } from "framer-motion";
const transition = { duration: 2, ease: "easeInOut" };

const Wave = () => {
    return (
        // background-image: linear-gradient(135deg, #567189, #6ad56c, #fad6a5);

        <svg xmlns="http://www.w3.org/2000/svg" width="100vw" height="100vh" viewBox="0 0 100 100" preserveAspectRatio="none" >
            <linearGradient id="gradient" x1="0" y1="0" x2="1" y2="0">
            <stop offset="0%" stop-color="#567189"/>

            <stop offset="25%" stop-color="#6ad56c"/>
            <stop offset="100%" stop-color="#fad6a5"/>
            </linearGradient>
          <motion.path
          id="arch"
            d="M0 50 Q 50 20, 100 50"
            stroke="url(#gradient)"
                        fill="transparent"
            strokeWidth={6}
            strokeLinecap="round"
            initial={{ pathLength: 0 }}
            animate={{ pathLength: 1 }}
            transition={transition}
          />
        </svg>
      );
};

export default Wave; 