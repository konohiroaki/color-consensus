import React, {Component} from "react";
import StatisticsHeader from "./StatisticsHeader";
import ColorBoard from "./ColorBoard";
import {connect} from "react-redux";

class StatisticsPage extends Component {

    constructor(props) {
        super(props);
        this.state = {
            voteCount: 0,
        };
    }

    render() {
        if (this.props.baseColor === null) {
            return null;
        }
        console.log("rendering statistics page", this.props.baseColor.code, this.props.colorCodeList[0]);

        return <div>
            <StatisticsHeader target={this.props.baseColor} voteCount={this.state.voteCount} history={this.props.history}/>
            <StatisticsPageButtons history={this.props.history}/>
            {this.props.baseColor.code === this.props.colorCodeList[0] &&
             <ColorBoard target={this.props.baseColor} colorCodeList={this.props.colorCodeList}
                         boardSideLength={this.props.boardSideLength} setVoteCount={(count) => this.setState({voteCount: count})}/>
            }
        </div>;
    }
}

const StatisticsPageButtons = props => (
    <div className="row">
        <button className="ml-auto btn btn-secondary m-3" onClick={() => props.history.push("/")}>
            Back to voting
        </button>
    </div>
);

const mapStateToProps = state => ({
    boardSideLength: state.board.sideLength,
    baseColor: state.board.baseColor,
    colorCodeList: state.board.colorCodeList,
});

export default connect(mapStateToProps)(StatisticsPage);
