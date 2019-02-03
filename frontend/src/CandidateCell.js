import React, {Component} from "react";
import {createSelectable} from "react-selectable-fast";

export class CandidateCell extends Component {

    shouldComponentUpdate(props) {
        return props.selected !== this.props.selected || props.selecting !== this.props.selecting;
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