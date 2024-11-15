import {PageProps} from "../page-props";
import {Layer, Rect, Stage, Circle} from "react-konva";

export const Brickbreaker = (props: PageProps) => {
    const width = window.innerWidth;
    const height = window.innerHeight;
    const initialPaddleX = width / 2 - 30;
    const paddleY = height - 20;

    const initialBallX = width / 2;
    const initialBallY = height - 25;

    return <>
        <Stage width={width} height={height}>
            <Layer>
                <Rect
                    x={initialPaddleX}
                    y={paddleY}
                    width={60}
                    height={5}
                    fill="blue"
                />
                <Circle
                    radius={5}
                    fill="red"
                    x={initialBallX}
                    y={initialBallY}
                />
            </Layer>
        </Stage>
    </>;
}