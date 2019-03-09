import React, {Component} from "react";
import StatisticsContentHeader from "./StatisticsContentHeader";
import ResultList from "./ResultList";
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

        return (
            <div>
                <StatisticsContentHeader target={this.props.target} history={this.props.history}/>
                <ResultList target={this.props.target} items={this.candidates} candidateSize={this.candidateSize}/>
            </div>
        );
    }

    componentDidMount() {
        this.updateCandidateList();
    }

    componentDidUpdate() {
        this.updateCandidateList();
    }

    updateCandidateList() {
        if (this.props.target !== this.state.target) {
            const target = this.props.target;
            axios.get(`http://localhost:5000/api/v1/colors/candidates/${target.code.substring(1)}?size=${Math.pow(this.candidateSize, 2)}`)
                .then(({data}) => {
                    console.log("main content got candidate list from server");
                    this.candidates = data;
                    this.setState({target: target});
                });
        }
    }
}

export default StatisticsPage;
