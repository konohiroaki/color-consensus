import React, {Component} from "react";
import {SelectableColorCell} from "./ColorCell";

class ColorBoard extends Component {

    constructor(props) {
        super(props);
        // +2 to avoid array out of bound error.
        const boardSize = this.props.candidateSize + 2;
        this.state = {
            border: Array(boardSize).fill(0)
                .map(() => Array(boardSize).fill({top: false, right: false, bottom: false, left: false}))
        };
        this.selected = Array(boardSize).fill(0)
            .map(() => Array(boardSize).fill(false));
        this.setCellState = this.setCellState.bind(this);
    }

    // TODO: place cells more nicely.
    render() {
        console.log("rendering voting color board");
        if (this.props.colors.length === 0) {
            console.log("colors array was empty");
            return null;
        }

        return <div className="text-center" style={{lineHeight: "0", padding: "10px"}}>
            {
                this.getCellList()
                    .split(this.props.candidateSize)
                    .map((v, k) => <div key={k}>{v}</div>)
            }
        </div>;
    };

    getCellList() {
        return this.props.colors.map((v, k) => {
            const i = Math.floor(k / this.props.candidateSize) + 1;
            const j = k % this.props.candidateSize + 1;
            return <SelectableColorCell key={k} color={this.props.colors[k]} border={this.state.border[i][j]}
                                        setCellState={this.setCellState.bind(null, {i: i, j: j})}/>;
        });
    }

    setCellState({i, j}, selected) {
        if (this.selected[i][j] !== selected) {
            // deep copy
            let b = JSON.parse(JSON.stringify(this.state.border));
            b[i][j].top = selected && !this.selected[i - 1][j];
            b[i][j].right = selected && !this.selected[i][j + 1];
            b[i][j].bottom = selected && !this.selected[i + 1][j];
            b[i][j].left = selected && !this.selected[i][j - 1];
            b[i - 1][j].bottom = this.selected[i - 1][j] && !selected;
            b[i][j + 1].left = this.selected[i][j + 1] && !selected;
            b[i + 1][j].top = this.selected[i + 1][j] && !selected;
            b[i][j - 1].right = this.selected[i][j - 1] && !selected;

            this.selected[i][j] = selected;

            // TODO: alternative way for this anti pattern. (but this works fine..)
            this.state.border = b;
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
