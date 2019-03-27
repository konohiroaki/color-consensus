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
        if (this.props.displayedColor === null) {
            return null;
        }
        console.log("rendering statistics page", this.props.displayedColor.code, this.props.displayedColorList[0]);

        return <div>
            <StatisticsHeader target={this.props.displayedColor} voteCount={this.state.voteCount} history={this.props.history}/>
            <StatisticsPageButtons history={this.props.history}/>
            {this.props.displayedColor.code === this.props.displayedColorList[0] &&
             <ColorBoard target={this.props.displayedColor} colorCodes={this.props.displayedColorList}
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
    displayedColor: state.colors.displayedColor,
    displayedColorList: state.colors.displayedColorList,
    boardSideLength: state.colors.boardSideLength,
});

export default connect(mapStateToProps)(StatisticsPage);
