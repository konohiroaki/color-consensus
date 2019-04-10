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
                .map(() => Array(boardSideLength).fill({top: false, right: false, bottom: false, left: false}))
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

    // FIXME: on first mount, it doesn't have vote data yet, so it doesn't show any border.
    // FIXME: on filter update, it doesn't apply because updateSelectedState isn't executed on update
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
        this.baseColor = this.props.baseColor;
        this.setRatio(this.props.votes);
        this.updateBorderState();
    }

    setRatio(votes) {
        this.ratio = this.ratio.map((e) => e.map(() => 0));
        votes
            .filter(vote => {
                console.log(this.props.nationalityFilter, vote.voter.nationality);
                return this.props.nationalityFilter === "" || this.props.nationalityFilter === vote.voter.nationality;
            })
            .flatMap(vote => vote.colors)
            .forEach(color => {
                console.log(color);
                const coord = this.coordForColor[color];
                this.ratio[coord.ii][coord.jj] += 1 / votes.length;
            });
    }

    updateBorderState() {
        // deep copy technique
        let border = JSON.parse(JSON.stringify(this.state.border));
        const ratio = this.ratio;
        const percentile = (100 - 60) / 100; // TODO: use user input for value subtracting from 100.
        for (let ii = 1; ii < border.length - 1; ii++) {
            for (let jj = 1; jj < border.length - 1; jj++) {
                border[ii][jj] = {
                    // TODO: make the condition simpler if possible
                    top: ratio[ii][jj] !== 0 && ratio[ii - 1][jj] > percentile && ratio[ii][jj] <= percentile
                         || ratio[ii][jj] !== 0 && ratio[ii - 1][jj] <= percentile && ratio[ii][jj] > percentile,
                    right: ratio[ii][jj] !== 0 && ratio[ii][jj + 1] > percentile && ratio[ii][jj] <= percentile
                           || ratio[ii][jj] !== 0 && ratio[ii][jj + 1] <= percentile && ratio[ii][jj] > percentile,
                    bottom: ratio[ii][jj] !== 0 && ratio[ii + 1][jj] > percentile && ratio[ii][jj] <= percentile
                            || ratio[ii][jj] !== 0 && ratio[ii + 1][jj] <= percentile && ratio[ii][jj] > percentile,
                    left: ratio[ii][jj] !== 0 && ratio[ii][jj - 1] > percentile && ratio[ii][jj] <= percentile
                          || ratio[ii][jj] !== 0 && ratio[ii][jj - 1] <= percentile && ratio[ii][jj] > percentile,
                };
            }
        }
        this.setState({border: border});
    }
}

const mapStateToProps = state => ({
    boardSideLength: state.board.sideLength,
    votes: state.statistics.votes,
    nationalityFilter: state.statistics.nationalityFilter
});

export default connect(mapStateToProps)(ColorBoard);
