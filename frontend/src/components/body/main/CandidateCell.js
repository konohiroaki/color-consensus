import React, {Component} from "react";
import {createSelectable} from "react-selectable-fast";

export class CandidateCell extends Component {

    constructor(props) {
        super(props);
        this.cellSize = "13px";
    }

    shouldComponentUpdate(nextProps) {
        // when color is changed, need to update
        if (this.props.color !== nextProps.color) {
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
        const color = this.props.selected || this.props.selecting ? "#fff" : "transparent";
        return (
            <div style={{
                display: "inline-block", padding: "1px", margin: "0 -1px -1px 0",
                borderWidth: "1px", borderStyle: "solid",
                borderTopColor: color, borderRightColor: color,
                borderBottomColor: color, borderLeftColor: color,
                userSelect: "none", userDrag: "none"
            }}>
                <div ref={this.props.selectableRef}
                     style={{width: this.cellSize, height: this.cellSize, backgroundColor: this.props.color}}/>
            </div>
        );
    }
}

const SelectableCandidateCell = createSelectable(CandidateCell);

export {SelectableCandidateCell};