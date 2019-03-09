import React, {Component} from "react";
import ResultCell from "./ResultCell";
import axios from "axios";
import update from "immutability-helper";

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
        this.coordForColor = {};
        this.hasSetBorder = false;

        this.updateSelectedState = this.updateSelectedState.bind(this);
    }

    render() {
        console.log("rendering result list page");
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
                this.coordForColor = update(this.coordForColor, {[this.props.items[key]]: {$set: {ii: ii, jj: jj}}});
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
        if (this.props.items.length !== 0 && !this.hasSetBorder) {
            this.updateSelectedState();
        }
    }

    componentDidUpdate() {
        if (this.props.items.length !== 0 && !this.hasSetBorder) {
            this.updateSelectedState();
        }
    }

    // FIXME: fix ugly code.
    // when item list comes, modify the this.state.border to show the statistics.
    updateSelectedState() {
        axios.get(`http://localhost:5000/api/v1/colors/detail/${this.props.target.lang}/${this.props.target.name}`)
            .then(({data}) => {
                console.log(data.colors);
                let border = this.state.border;
                for (let color in data.colors) {
                    const coord = this.coordForColor[color];
                    border = update(border, {[coord.ii]: {[coord.jj]: {$set: {top: true, right: true, bottom: true, left: true}}}});
                }
                this.hasSetBorder = true;
                // FIXME: set proper border.
                // TODO: instead of setting #fff or transparent, categorize the colors to several percentiles and border with several colors for each category.
                this.setState({border: border});
            });
    }
}

export default ResultList;
