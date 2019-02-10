import React, {Component} from "react";
import {SelectableCandidateCell} from "./CandidateCell";

class CandidateList extends Component {

    render() {
        console.log("rendering candidate list");
        if (this.props.items.length === 0) {
            console.log("candidate list is empty");
            return <div/>;
        }
        let list = [];
        for (let i = 0; i < this.props.candidateSize; i++) {
            let row = [];
            for (let j = 0; j < this.props.candidateSize; j++) {
                row.push(<SelectableCandidateCell key={i * this.props.candidateSize + j} color={this.props.items[i][j]}/>);
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