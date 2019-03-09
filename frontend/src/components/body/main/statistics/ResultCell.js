import React, {Component} from "react";

class ResultCell extends Component {

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
        console.log("rendering result cell");
        return (
            <div style={{
                display: "inline-block", padding: "1px",
                borderWidth: "1px", borderStyle: "solid",
                borderTopColor: ResultCell.getBorderColor(this.props.border.top),
                borderRightColor: ResultCell.getBorderColor(this.props.border.right),
                borderBottomColor: ResultCell.getBorderColor(this.props.border.bottom),
                borderLeftColor: ResultCell.getBorderColor(this.props.border.left),
                userSelect: "none", userDrag: "none"
            }}>
                <div style={{width: this.cellSize, height: this.cellSize, backgroundColor: this.props.color}}/>
            </div>
        );
    }

    static getBorderColor(hasBorder) {
        return hasBorder ? "#fff" : "transparent";
    }
}

export default ResultCell;
