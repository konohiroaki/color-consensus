import React, {Component} from "react";
import ResultCell from "./ResultCell";
import axios from "axios";

class ResultList extends Component {

    constructor(props) {
        super(props);
        // +2 to avoid array out of bound error.
        this.boardSize = this.props.candidateSize + 2;
        this.state = {
            border: Array(this.boardSize).fill(Array(this.boardSize).fill(
                {top: false, right: false, bottom: false, left: false}
            ))
        };
        this.updateBorderState = this.updateBorderState.bind(this);
    }

    render() {
        if (this.props.items.length === 0) {
            console.log("candidate list is empty");
            return <div/>;
        }

        let list = [];
        for (let i = 0; i < this.props.candidateSize; i++) {
            let row = [];
            for (let j = 0; j < this.props.candidateSize; j++) {
                const key = i * this.props.candidateSize + j;
                const ii = i + 1, jj = j + 1;
                row.push(<ResultCell key={key} color={this.props.items[key]}
                                     border={this.state.border[ii][jj]}/>);
            }
            list.push(<div key={i}>{row}</div>);
        }
        return (
            <div className="text-center" style={{lineHeight: "0", padding: "10px"}}>
                {list}
            </div>
        );
    }

    componentDidMount() {
        this.updateBorderState();
    }

    updateBorderState() {
        axios.get(`http://localhost:5000/api/v1/colors/detail/${this.props.target.lang}/${this.props.target.name}`)
            .then(({data}) => {
                console.log(data.colors);
                // TODO: update this.state.border according to the consensus.
            });
    }
}

export default ResultList;
