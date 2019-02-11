import React, {Component} from "react";
import {SelectableCandidateCell} from "./CandidateCell";
import update, {extend} from "immutability-helper";

class CandidateList extends Component {

    constructor(props) {
        super(props);
        this.state = {
            border: new Array(this.props.candidateSize + 2).fill(
                new Array(this.props.candidateSize + 2).fill(
                    {top: false, right: false, bottom: false, left: false}
                ))
        };
        this.selected = new Array(this.props.candidateSize + 2).fill(
            new Array(this.props.candidateSize + 2).fill(
                false
            ));
        this.setCellState = this.setCellState.bind(this);
    }

    setCellState({i, j}, selected) {
        if (this.selected[i][j] !== selected) {
            const bar = update(this.state.border, {
                [i]: {
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
            this.state.border = bar;
            this.forceUpdate();
        }
    }

    // TODO: place cells more nicely.
    render() {
        console.log("rendering candidate list");
        if (this.props.items.length === 0) {
            console.log("candidate list is empty");
            return <div/>;
        }
        let list = [];
        for (let i = 1; i < this.props.candidateSize + 1; i++) {
            let row = [];
            for (let j = 1; j < this.props.candidateSize + 1; j++) {
                const key = i * this.props.candidateSize + j;
                row.push(<SelectableCandidateCell key={key} color={this.props.items[key]}
                                                  border={this.state.border[i][j]}
                                                  setCellState={this.setCellState.bind(null, {i: i, j: j})}/>);
            }
            list.push(<div key={i}>{row}</div>);
        }
        return (
            <div className="text-center" style={{lineHeight: "0", padding: "10px"}}>
                {list}
            </div>
        );
    }
}

export default CandidateList;