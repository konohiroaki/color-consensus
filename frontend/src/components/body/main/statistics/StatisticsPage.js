import React, {Component} from "react";
import StatisticsHeader from "./StatisticsHeader";
import ColorBoard from "./ColorBoard";
import {connect} from "react-redux";
import {actions as statistics} from "../../../../modules/statistics";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faArrowLeft} from "@fortawesome/free-solid-svg-icons";

class StatisticsPage extends Component {

    render() {
        if (this.props.baseColor === null) {
            return null;
        }
        console.log("rendering [statistics page]",
            "baseColor.code:", this.props.baseColor.code,
            "codeList[0]:", this.props.colorCodeList[0]);

        return <div>
            <StatisticsPageButtons className="border-bottom border-secondary pb-3"
                                   history={this.props.history}/>
            <StatisticsHeader baseColor={this.props.baseColor}/>
            {this.props.baseColor.code === this.props.colorCodeList[0] &&
             <ColorBoard baseColor={this.props.baseColor} colorCodeList={this.props.colorCodeList}/>}
        </div>;
    }

    componentDidMount() {
        if (this.props.baseColor !== null && this.props.baseColor.code === this.props.colorCodeList[0]) {
            this.props.setVotes(this.props.baseColor);
        }
    }

    componentDidUpdate() {
        this.componentDidMount();
    }
}

const StatisticsPageButtons = ({className, history}) => (
    <div className={className + " row"}>
        <button className="mr-auto btn btn-secondary m-3" onClick={() => history.push("/")}>
            <span><FontAwesomeIcon icon={faArrowLeft}/> Go to voting</span>
        </button>
    </div>
);

const mapStateToProps = state => ({
    boardSideLength: state.board.sideLength,
    baseColor: state.board.baseColor,
    colorCodeList: state.board.colorCodeList,
});

const mapDispatchToProps = dispatch => ({
    setVotes: color => dispatch(statistics.setVotes(color)),
});

export default connect(mapStateToProps, mapDispatchToProps)(StatisticsPage);
