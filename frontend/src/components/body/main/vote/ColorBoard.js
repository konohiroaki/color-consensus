import React, {Component} from "react";
import {SelectableColorCell} from "./ColorCell";
import update from "immutability-helper";

class ColorBoard extends Component {

    constructor(props) {
        super(props);
        // +2 to avoid array out of bound error.
        this.boardSize = this.props.candidateSize + 2;
        this.state = {
            border: Array(this.boardSize).fill(Array(this.boardSize).fill(
                {top: false, right: false, bottom: false, left: false}
            ))
        };
        this.selected = Array(this.boardSize).fill(Array(this.boardSize).fill(false));
        this.setCellState = this.setCellState.bind(this);
    }

    // TODO: place cells more nicely.
    render() {
        console.log("rendering color board for voting");
        if (this.props.colors.length === 0) {
            console.log("colors array was empty");
            return null;
        }

        return <div className="text-center" style={{lineHeight: "0", padding: "10px"}}>
            {this.getCellList()}
        </div>;
    };

    getCellList() {
        const list = this.props.colors.map((v, k) => {
            const i = Math.floor(k / this.props.candidateSize) + 1;
            const j = k % this.props.candidateSize + 1;
            return <SelectableColorCell key={k} color={this.props.colors[k]} border={this.state.border[i][j]}
                                        setCellState={this.setCellState.bind(null, {i: i, j: j})}/>;
        });
        return list.split(this.props.candidateSize).map((v, k) => <div key={k}>{v}</div>);
    }

    setCellState({i, j}, selected) {
        if (this.selected[i][j] !== selected) {
            const border = update(this.state.border, {
                [i]: {
                    // [i][j-1] (left cell)'s right border is active when left cell is selected and current cell is NOT selected.
                    [j - 1]: {right: {$set: this.selected[i][j - 1] && !selected}},
                    [j]: {
                        top: {$set: selected && !this.selected[i - 1][j]},
                        right: {$set: selected && !this.selected[i][j + 1]},
                        bottom: {$set: selected && !this.selected[i + 1][j]},
                        left: {$set: selected && !this.selected[i][j - 1]}
                    },
                    [j + 1]: {left: {$set: this.selected[i][j + 1] && !selected}}
                },
                [i - 1]: {[j]: {bottom: {$set: this.selected[i - 1][j] && !selected}}},
                [i + 1]: {[j]: {top: {$set: this.selected[i + 1][j] && !selected}}}
            });
            this.selected = update(this.selected, {[i]: {[j]: {$set: selected}}});

            // TODO: alternative way for this anti pattern. (but this works fine..)
            this.state.border = border;
            this.forceUpdate();
        }
    }
}

Array.prototype.split = function (n) {
    let array = this;
    let result = [];

    for (let i = 0; i < array.length; i += n) {
        result.push(array.slice(i, i + n));
    }
    return result;
};

export default ColorBoard;
