import React, {Component} from "react";
import {createSelectable} from "react-selectable-fast";

export class CandidateCell extends Component {

    shouldComponentUpdate(nextProps, nextState) {
        // when color is changed, need to update
        if (this.props.color !== nextProps.color) {
            return true;
        }
        // when selecting -> selected, no need to update
        if (this.props.selecting && nextProps.selected) {
            return false;
        }
        // when selected, selecting state is different, need to update
        return nextProps.selected !== this.props.selected || nextProps.selecting !== this.props.selecting;
    }

    render() {
        console.log("rendering candidate cell");
        const cellSize = "13px";
        let color = this.props.selected || this.props.selecting ? "#fff" : "transparent";
        return (
            <div style={{
                display: "inline-block", padding: "1px", margin: "0 -1px -1px 0",
                borderWidth: "1px", borderStyle: "solid", borderColor: color,
                userSelect: "none", userDrag: "none"
            }}>
                <div ref={this.props.selectableRef} style={{width: cellSize, height: cellSize, backgroundColor: this.props.color}}/>
            </div>
        );
    }
}

const SelectableCandidateCell = createSelectable(CandidateCell);

export {SelectableCandidateCell};