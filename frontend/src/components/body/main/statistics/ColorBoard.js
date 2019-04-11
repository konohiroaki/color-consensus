import React, {Component} from "react";
import ColorCell from "./ColorCell";
import axios from "axios";
import {isSameColor} from "../../../common/Utility";
import {connect} from "react-redux";

class ColorBoard extends Component {

    constructor(props) {
        super(props);
        this.baseColor = {};
    }

    render() {
        if (this.props.colorCodeList.length === 0 || this.props.cellBorder.length === 0) {
            return null;
        }
        console.log("rendering [statistics color board]");

        return <div className="text-center" style={{lineHeight: "0", padding: "10px"}}>
            {
                this.getCellList().split(this.props.boardSideLength)
                    .map((v, k) => <div key={k}>{v}</div>)
            }
        </div>;
    }

    getCellList() {
        return this.props.colorCodeList.map((v, k) => {
            const ii = Math.floor(k / this.props.boardSideLength) + 1;
            const jj = k % this.props.boardSideLength + 1;
            return <ColorCell key={k} colorCode={this.props.colorCodeList[k]} coord={{ii: ii, jj: jj}}
                              border={this.props.cellBorder[ii][jj]}/>;
        });
    }
}

const mapStateToProps = state => ({
    boardSideLength: state.board.sideLength,
    votes: state.statistics.votes,
    cellBorder: state.statistics.cellBorder,
});

export default connect(mapStateToProps)(ColorBoard);
