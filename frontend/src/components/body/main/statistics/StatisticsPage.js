import React, {Component} from "react";
import StatisticsHeader from "./StatisticsHeader";
import ColorBoard from "./ColorBoard";
import axios from "axios";

class StatisticsPage extends Component {

    constructor(props) {
        super(props);
        this.state = {};

        this.candidateSize = 31;
        this.candidates = [];
    }

    render() {
        console.log("rendering statistics page", this.props.target);
        if (this.props.target === undefined) {
            return null;
        }

        return <div>
            <StatisticsHeader target={this.props.target} history={this.props.history}/>
            <VotingPageButton history={this.props.history}/>
            <ColorBoard target={this.props.target} items={this.candidates} candidateSize={this.candidateSize}/>
        </div>;
    }

    componentDidMount() {
        this.updateCandidateList();
    }

    componentDidUpdate() {
        this.updateCandidateList();
    }

    updateCandidateList() {
        if (this.props.target !== this.state.target) {
            const url = `${process.env.WEBAPI_HOST}/api/v1/colors/candidates/${this.props.target.code.substring(1)}?size=${Math.pow(this.candidateSize, 2)}`;
            axios.get(url).then(({data}) => {
                console.log("main content got candidate list from server");
                this.candidates = data;
                this.setState({target: this.props.target});
            });
        }
    }
}

const VotingPageButton = (props) => (
    <div className="row">
        <div className="ml-auto">
            <button className="btn btn-secondary m-3" onClick={() => props.history.push("/")}>
                Back to voting
            </button>
        </div>
    </div>
);

export default StatisticsPage;
