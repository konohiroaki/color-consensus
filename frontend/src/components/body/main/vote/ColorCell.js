import React, {Component} from "react";
import {createSelectable} from "react-selectable-fast";
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
        // when selecting -> selected, no need to update
        if (this.props.selecting && nextProps.selected) {
            return false;
        }
        // when selected, selecting state is different, need to update
        return this.props.selected !== nextProps.selected || this.props.selecting !== nextProps.selecting;
    }

    render() {
        console.log("rendering voting color cell");
        return <div style={{
            display: "inline-block", padding: "1px",
            borderWidth: "1px", borderStyle: "solid",
            borderTopColor: getBorderColor(this.props.border.top),
            borderRightColor: getBorderColor(this.props.border.right),
            borderBottomColor: getBorderColor(this.props.border.bottom),
            borderLeftColor: getBorderColor(this.props.border.left),
            userSelect: "none", userDrag: "none"
        }}>
            <div ref={this.props.selectableRef}
                 style={{width: this.props.cellSize, height: this.props.cellSize, backgroundColor: this.props.colorCode}}/>
        </div>;
    }

    componentDidUpdate() {
        this.props.setCellState(this.props.selected || this.props.selecting);
    }
}

const getBorderColor = hasBorder => hasBorder ? "#fff" : "transparent";

const mapStateToProps = state => ({
    cellSize: state.board.cellSize,
});

const SelectableColorCell = createSelectable(connect(mapStateToProps)(ColorCell));

export {SelectableColorCell};
