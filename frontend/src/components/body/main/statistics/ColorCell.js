import React, {Component} from "react";

class ColorCell extends Component {

    constructor(props) {
        super(props);
        this.cellSize = "15px";
    }

    shouldComponentUpdate(nextProps) {
        // when color is changed, need to update
        if (this.props.color !== nextProps.color) {
            return true;
        }
        // when border state changed, need to update
        if (this.props.border.top !== nextProps.border.top
            || this.props.border.right !== nextProps.border.right
            || this.props.border.bottom !== nextProps.border.bottom
            || this.props.border.left !== nextProps.border.left) {
            return true;
        }
        return false;
    }

    render() {
        console.log("rendering statistics color cell");
        return <div style={{
            display: "inline-block", padding: "1px",
            borderWidth: "1px", borderStyle: "solid",
            borderTopColor: getBorderColor(this.props.border.top),
            borderRightColor: getBorderColor(this.props.border.right),
            borderBottomColor: getBorderColor(this.props.border.bottom),
            borderLeftColor: getBorderColor(this.props.border.left),
            userSelect: "none", userDrag: "none"
        }}>
            <div style={{width: this.cellSize, height: this.cellSize, backgroundColor: this.props.color}}/>
        </div>;
    }
}

const getBorderColor = category => {
    if (category === 0) {
        return "transparent";
    } else if (category === 1) {
        return "#999";
    } else if (category === 2) {
        return "#ccc";
    } else {
        return "#fff";
    }
};

export default ColorCell;
