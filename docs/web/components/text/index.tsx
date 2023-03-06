import React from "react";
import styles from './style.module.css'
import TextTransition, { presets } from "react-text-transition";

const TEXTS = [
    "Cosmos",
    "Berachain",
    "Anything",
];

const Text = () => {
    const [index, setIndex] = React.useState(0);

    React.useEffect(() => {
        const intervalId = setInterval(() =>
            setIndex(index => index + 1),
            2000 // every 3 seconds
        );
        return () => clearTimeout(intervalId);
    }, []);

    return (
        
        <h1 className={styles.subheader}>
            The New Standard of EVM on
            <TextTransition springConfig={presets.wobbly} inline style={{marginLeft: '4px'}}> 
             {TEXTS[index % TEXTS.length]}
            </TextTransition>
        </h1>
        
    );
};

export default Text;