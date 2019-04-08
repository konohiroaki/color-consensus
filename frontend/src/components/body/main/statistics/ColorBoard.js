import React, {Component} from "react";
import ColorCell from "./ColorCell";
import axios from "axios";
import {isSameColor} from "../../../common/Utility";
import {connect} from "react-redux";

class ColorBoard extends Component {

    constructor(props) {
        super(props);
        // +2 to avoid array out of bound error.
        const boardSideLength = this.props.boardSideLength + 2;
        this.state = {
            border: Array(boardSideLength).fill(0)
                .map(() => Array(boardSideLength).fill({top: 0, right: 0, bottom: 0, left: 0}))
        };
        this.ratio = Array(boardSideLength).fill(0)
            .map(() => Array(boardSideLength).fill(0));
        this.coordForColor = {};
        this.baseColor = {};

        this.updateSelectedState = this.updateSelectedState.bind(this);
        this.updateBorderState = this.updateBorderState.bind(this);
    }

    render() {
        if (this.props.colorCodeList.length === 0) {
            return null;
        }
        console.log("rendering [statistics color board]");

        const list = this.getCellList();
        this.setCoordForColor(list);

        return <div className="text-center" style={{lineHeight: "0", padding: "10px"}}>
            {
                list.split(this.props.boardSideLength)
                    .map((v, k) => <div key={k}>{v}</div>)
            }
        </div>;
    }

    componentDidMount() {
        if (this.props.colorCodeList.length !== 0 && !isSameColor(this.props.baseColor, this.baseColor)) {
            this.updateSelectedState();
        }
    }

    getCellList() {
        return this.props.colorCodeList.map((v, k) => {
            const ii = Math.floor(k / this.props.boardSideLength) + 1;
            const jj = k % this.props.boardSideLength + 1;
            return <ColorCell key={k} colorCode={this.props.colorCodeList[k]} coord={{ii: ii, jj: jj}}
                              border={this.state.border[ii][jj]}/>;
        });
    }

    // coordForColor => {#ff0000: {ii: 1, jj: 1}, #f00000: {ii: 1, jj: 2}, ...}
    setCoordForColor(list) {
        this.coordForColor = list.reduce((acc, v) => {
            acc[v.props.colorCode] = {ii: v.props.coord.ii, jj: v.props.coord.jj};
            return acc;
        }, {});
    }

    updateSelectedState() {
        const url = `${process.env.WEBAPI_HOST}/api/v1/votes-stats?lang=${this.props.baseColor.lang}&name=${this.props.baseColor.name}`;
        axios.get(url).then(({data}) => {
            // data => {vote:10, colors:{#ff0000:5, #ff1000:3, ...}
            this.baseColor = this.props.baseColor;
            this.props.setVoteCount(data.count);
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
        // deep copy technique
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

const mapStateToProps = state => ({
    boardSideLength: state.board.sideLength,
});

export default connect(mapStateToProps)(ColorBoard);
