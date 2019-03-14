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

        const list = this.getCellList();
        this.setCoordForColor(list);

        return <div className="text-center" style={{lineHeight: "0", padding: "10px"}}>
            {
                list
                    .split(this.props.candidateSize)
                    .map((v, k) => <div key={k}>{v}</div>)
            }
        </div>;
    }

    getCellList() {
        return this.props.colors.map((v, k) => {
            const ii = Math.floor(k / this.props.candidateSize) + 1;
            const jj = k % this.props.candidateSize + 1;
            return <ColorCell key={k} color={this.props.colors[k]} coord={{ii: ii, jj: jj}}
                              border={this.state.border[ii][jj]}/>;
        });
    }

    setCoordForColor(list) {
        this.coordForColor = list.reduce((acc, v) => {
            acc[v.props.color] = {ii: v.props.coord.ii, jj: v.props.coord.jj};
            return acc;
        }, {});
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
