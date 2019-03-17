import React, {Component} from "react";
import ColorCell from "./ColorCell";
import axios from "axios";

class ColorBoard extends Component {

    constructor(props) {
        super(props);
        // +2 to avoid array out of bound error.
        const boardSize = this.props.candidateSize + 2;
        this.state = {
            border: Array(boardSize).fill(0)
                .map(() => Array(boardSize).fill({top: 0, right: 0, bottom: 0, left: 0}))
        };
        this.ratio = Array(boardSize).fill(0)
            .map(() => Array(boardSize).fill(0));
        this.coordForColor = {};
        this.target = {};

        this.updateSelectedState = this.updateSelectedState.bind(this);
        this.updateBorderState = this.updateBorderState.bind(this);
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

    componentDidMount() {
        if (this.props.colors.length !== 0
            && (this.target.lang !== this.props.target.lang || this.target.name !== this.props.target.name)) {
            this.updateSelectedState();
        }
    }

    componentDidUpdate() {
        if (this.props.colors.length !== 0
            && (this.target.lang !== this.props.target.lang || this.target.name !== this.props.target.name)) {
            this.updateSelectedState();
        }
    }

    getCellList() {
        return this.props.colors.map((v, k) => {
            const ii = Math.floor(k / this.props.candidateSize) + 1;
            const jj = k % this.props.candidateSize + 1;
            return <ColorCell key={k} color={this.props.colors[k]} coord={{ii: ii, jj: jj}}
                              border={this.state.border[ii][jj]}/>;
        });
    }

    // result will look like {#ff0000: {ii: 1, jj: 1}, #f00000: {ii: 1, jj: 2}, ...}
    setCoordForColor(list) {
        this.coordForColor = list.reduce((acc, v) => {
            acc[v.props.color] = {ii: v.props.coord.ii, jj: v.props.coord.jj};
            return acc;
        }, {});
    }

    updateSelectedState() {
        const url = `${process.env.WEBAPI_HOST}/api/v1/colors/detail/${this.props.target.lang}/${this.props.target.name}`;
        axios.get(url).then(({data}) => {
            this.target = this.props.target;
            this.setRatio(data.vote, data.colors);
            this.updateBorderState();
        });
    }

    setRatio(vote, colors) {
        this.ratio = this.ratio.map((e) => e.map(() => 0));
        for (let color in colors) {
            const coord = this.coordForColor[color];
            this.ratio[coord.ii][coord.jj] = getCategory(colors[color] / vote);
        }
    }

    updateBorderState() {
        let border = JSON.parse(JSON.stringify(this.state.border));
        const ratio = this.ratio;
        for (let ii = 1; ii < border.length - 1; ii++) {
            for (let jj = 1; jj < border.length - 1; jj++) {
                border[ii][jj] = {
                    top: ratio[ii - 1][jj] === ratio[ii][jj] ? 0 : ratio[ii][jj],
                    right: ratio[ii][jj + 1] === ratio[ii][jj] ? 0 : ratio[ii][jj],
                    bottom: ratio[ii + 1][jj] === ratio[ii][jj] ? 0 : ratio[ii][jj],
                    left: ratio[ii][jj - 1] === ratio[ii][jj] ? 0 : ratio[ii][jj]
                };
            }
        }
        this.setState({border: border});
    }
}

const getCategory = ratio => {
    if (ratio <= 0.10) {
        return 0;
    } else if (ratio <= 0.50) {
        return 1;
    } else if (ratio <= 0.75) {
        return 2;
    }
    return 3;
};

export default ColorBoard;
