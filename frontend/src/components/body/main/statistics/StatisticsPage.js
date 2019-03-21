import React, {Component} from "react";
import StatisticsHeader from "./StatisticsHeader";
import ColorBoard from "./ColorBoard";
import axios from "axios";

class StatisticsPage extends Component {

    constructor(props) {
        super(props);
        this.state = {
            voteCount: 0
        };

        this.boardSideLength = 31;
        this.colorCodeList = [];
    }

    render() {
        console.log("rendering statistics page", this.props.target);
        if (this.props.target === undefined) {
            return null;
        }

        return <div>
            <StatisticsHeader target={this.props.target} voteCount={this.state.voteCount} history={this.props.history}/>
            <StatisticsPageButtons history={this.props.history}/>
            <ColorBoard target={this.state.target} colorCodes={this.colorCodeList} boardSideLength={this.boardSideLength}
                        setVoteCount={(count) => this.setState({voteCount: count})}/>
        </div>;
    }

    componentDidMount() {
        this.updateColorCodeList();
    }

    componentDidUpdate() {
        this.updateColorCodeList();
    }

    updateColorCodeList() {
        if (this.props.target !== this.state.target) {
            const baseCode = this.props.target.code.substring(1);
            const size = Math.pow(this.boardSideLength, 2);
            const url = `${process.env.WEBAPI_HOST}/api/v1/colors/candidates/${baseCode}?size=${size}`;
            axios.get(url).then(({data}) => {
                this.colorCodeList = data;
                this.setState({target: this.props.target});
            });
        }
    }
}

const StatisticsPageButtons = props => (
    <div className="row">
        <button className="ml-auto btn btn-secondary m-3" onClick={() => props.history.push("/")}>
            Back to voting
        </button>
    </div>
);

export default StatisticsPage;
