import React, {Component} from "react";
import {createSelectable} from "react-selectable-fast";

export class CandidateCell extends Component {

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
        // when selecting -> selected, no need to update
        if (this.props.selecting && nextProps.selected) {
            return false;
        }
        // when selected, selecting state is different, need to update
        return this.props.selected !== nextProps.selected || this.props.selecting !== nextProps.selecting;
    }

    render() {
        console.log("rendering candidate cell");
        return (
            <div style={{
                display: "inline-block", padding: "1px", margin: "0 -1px -1px 0",
                borderWidth: "1px", borderStyle: "solid",
                borderTopColor: CandidateCell.getBorderColor(this.props.border.top),
                borderRightColor: CandidateCell.getBorderColor(this.props.border.right),
                borderBottomColor: CandidateCell.getBorderColor(this.props.border.bottom),
                borderLeftColor: CandidateCell.getBorderColor(this.props.border.left),
                userSelect: "none", userDrag: "none"
            }}>
                <div ref={this.props.selectableRef}
                     style={{width: this.cellSize, height: this.cellSize, backgroundColor: this.props.color}}/>
            </div>
        );
    }

    static getBorderColor(hasBorder) {
        return hasBorder ? "#fff" : "transparent";
    }

    componentDidUpdate() {
        this.props.setCellState(this.props.selected || this.props.selecting);
    }
}

const SelectableCandidateCell = createSelectable(CandidateCell);

export {SelectableCandidateCell};