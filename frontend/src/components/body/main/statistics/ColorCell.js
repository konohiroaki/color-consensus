import React, {Component} from "react";
import {connect} from "react-redux";

class ColorCell extends Component {

    shouldComponentUpdate(nextProps) {
        // when color is changed, need to update
        if (this.props.colorCode !== nextProps.colorCode) {
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
        console.log("rendering [statistics color cell]");
        return <div style={{
            display: "inline-block", padding: "1px",
            borderWidth: "1px", borderStyle: "solid",
            borderTopColor: getBorderColor(this.props.border.top),
            borderRightColor: getBorderColor(this.props.border.right),
            borderBottomColor: getBorderColor(this.props.border.bottom),
            borderLeftColor: getBorderColor(this.props.border.left),
            userSelect: "none", userDrag: "none"
        }}>
            <div style={{width: this.props.cellSize, height: this.props.cellSize, backgroundColor: this.props.colorCode}}/>
        </div>;
    }
}

const getBorderColor = flag => flag ? "#fff" : "transparent";

const mapStateToProps = state => ({
    cellSize: state.board.cellSize,
});

export default connect(mapStateToProps)(ColorCell);
