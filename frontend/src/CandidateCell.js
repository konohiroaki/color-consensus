import React, {Component} from "react";
import {createSelectable} from "react-selectable-fast";

export class CandidateCell extends Component {

    render() {
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

class Counter extends Component {
    constructor() {
        super();
        this.state = {
            selectedItems: [],
            selectingItems: [],
        };
    }

    handleSelectin(selectingItems) {
        this.setState({selectingItems});
    };

    handleSelectionFinis(selectedItems) {
        this.setState({
            selectedItems: selectedItems,
            selectingItems: [],
        });
    };

    render() {
        const {selectedItems, selectingItems} = this.state;

        let selectedColors = [];
        let selectingColors = [];
        for (const v of selectedItems) {
            selectedColors.push(v.props.color);
        }
        console.log(selectingItems);
        for (const v of selectingItems) {
            selectingColors.push(v.props.color);
        }
        return (
            <div>
                <p>Selected:{" "}<span className="counter">{selectedColors}</span></p>
                <p>Selecting:{" "}<span className="counter">{selectingColors}</span></p>
            </div>
        );
    }
}

const SelectableCandidateCell = createSelectable(CandidateCell);

export {SelectableCandidateCell};