import React, {Component} from "react";
import ColorCell from "./ColorCell";
import axios from "axios";
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
        this.coordForColor = {};
        this.hasSetBorder = false;

        this.updateSelectedState = this.updateSelectedState.bind(this);
    }

    render() {
        console.log("rendering color board for statistics");
        if (this.props.colors.length === 0) {
            console.log("colors array was empty");
            return null;
        }

        let list = [];
        for (let i = 0; i < this.props.candidateSize; i++) {
            let row = [];
            for (let j = 0; j < this.props.candidateSize; j++) {
                const key = i * this.props.candidateSize + j;
                const ii = i + 1, jj = j + 1;
                row.push(<ColorCell key={key} color={this.props.colors[key]}
                                    border={this.state.border[ii][jj]}/>);
                this.coordForColor = update(this.coordForColor, {[this.props.colors[key]]: {$set: {ii: ii, jj: jj}}});
            }
            list.push(<div key={i}>{row}</div>);
        }
        return <div className="text-center" style={{lineHeight: "0", padding: "10px"}}>
            {list}
        </div>;
    }

    componentDidMount() {
        if (this.props.colors.length !== 0 && !this.hasSetBorder) {
            this.updateSelectedState();
        }
    }

    componentDidUpdate() {
        if (this.props.colors.length !== 0 && !this.hasSetBorder) {
            this.updateSelectedState();
        }
    }

    // FIXME: fix ugly code.
    // when colors array comes, modify the this.state.border to show the statistics.
    updateSelectedState() {
        axios.get(`${process.env.WEBAPI_HOST}/api/v1/colors/detail/${this.props.target.lang}/${this.props.target.name}`)
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

export default ColorBoard;
