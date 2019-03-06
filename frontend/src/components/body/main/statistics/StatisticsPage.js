import React, {Component} from "react";
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
        console.log("rendering statistics page");
        if (this.props.target === undefined) {
            return null;
        }

        return (
            <div>
                <div className="card bg-dark border border-secondary">
                    <div className="card-body">
                        <div className="row ml-0 mr-0">
                            <div className="col-3 card bg-dark border border-secondary p-2 text-center">
                                <div className="row">
                                    <span className="col-4 border-right border-secondary p-3">{this.props.target.lang}</span>
                                    <span className="col-8 p-3">{this.props.target.name}</span>
                                </div>
                            </div>
                            {/* TODO: complete select box impl */}
                            <div className="col-9 input-group">
                                <select className="custom-select" defaultValue="">
                                    <option value="">Nationality</option>
                                    <option value="3">Others</option>
                                </select>
                                <select className="custom-select" defaultValue="">
                                    <option value="">Age Group</option>
                                    <option value="3">Others</option>
                                </select>
                                <select className="custom-select" defaultValue="">
                                    <option value="">Gender</option>
                                    <option value="1">Male</option>
                                    <option value="2">Female</option>
                                    <option value="3">Others</option>
                                </select>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="row">
                    <div className="ml-auto">
                        <button className="btn btn-secondary m-3" onClick={() => this.props.history.push("/")}>
                            Back to voting
                        </button>
                    </div>
                </div>

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
