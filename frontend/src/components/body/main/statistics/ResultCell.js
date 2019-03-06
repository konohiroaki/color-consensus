import React, {Component} from "react";

class ResultCell extends Component {

    constructor(props) {
        super(props);
        this.cellSize = "15px";
    }

    render() {
        console.log("rendering candidate cell");
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
