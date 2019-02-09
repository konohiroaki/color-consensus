import React, {Component} from "react";
import {SelectableCandidateCell} from "./CandidateCell";

class CandidateList extends Component {

    shouldComponentUpdate(props) {
        return props.items !== this.props.items;
    }

    render() {
        console.log("rendering candidate list");
        if (this.props.items.length === 0) {
            return <div/>;
        }
        let list = [];
        for (let i = 0; i < 51; i++) {
            let row = [];
            for (let j = 0; j < 51; j++) {
                row.push(<SelectableCandidateCell key={i * 51 + j} color={this.props.items[i][j]}/>);
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