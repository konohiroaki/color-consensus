import React, {Component} from "react";
import {DeselectAll, SelectableGroup} from "react-selectable-fast";
import axios from "axios";
import CandidateList from "./CandidateList";

class VotingPage extends Component {

    constructor(props) {
        super(props);
        this.state = {};

        this.candidateSize = 31;
        this.candidates = [];
        this.selected = [];

        this.updateCandidateList = this.updateCandidateList.bind(this);
        this.handleSubmitClick = this.handleSubmitClick.bind(this);
        this.submit = this.submit.bind(this);
        this.handleSelectionFinish = this.handleSelectionFinish.bind(this);
    }

    render() {
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
                        </div>
                    </div>
                </div>
                <div className="row">
                    <div className="ml-auto">
                        <div>
                            <button className="btn btn-secondary m-3" onClick={() => this.props.history.push("/statistics")}>
                                Skip to statistics
                            </button>
                            <button className="btn btn-primary m-3" onClick={this.handleSubmitClick}>
                                Submit
                            </button>
                        </div>
                    </div>
                </div>

                <SelectableGroup
                    className="selectable"
                    clickClassName="tick"
                    enableDeselect
                    allowClickWithoutSelected={true}
                    onSelectionFinish={this.handleSelectionFinish}>
                    <div className="row">
                        <div className="ml-auto">
                            <DeselectAll className="btn btn-secondary m-3">Clear</DeselectAll>
                        </div>
                    </div>
                    <CandidateList items={this.candidates} candidateSize={this.candidateSize}/>
                </SelectableGroup>
            </div>
        );
    }

    componentDidMount() {
        this.updateCandidateList();
    }

    componentDidUpdate() {
        this.updateCandidateList();
    }

    // this.props.target -> new target color
    // this.state.target -> current target color
    updateCandidateList() {
        if (this.props.target !== this.state.target) {
            const target = this.props.target;
            axios.get(`http://localhost:5000/api/v1/colors/candidates/${target.code.substring(1)}?size=${Math.pow(this.candidateSize, 2)}`)
                .then(({data}) => {
                    console.log("main content got candidate list from server");
                    this.candidates = data;
                    // FIXME: doesn't deselect on color change.
                    this.selected = [];
                    this.setState({target: target});
                });
        }
    }

    handleSubmitClick() {
        const userId = this.props.userId;
        if (userId === undefined || userId === null) {
            this.props.loginModalRef.openLoginModal(this.submit);
        } else {
            this.submit();
        }
    }

    submit() {
        const {lang, name} = this.state.target;
        axios.post(`http://localhost:5000/api/v1/votes`, {
            "lang": lang,
            "name": name,
            "colors": this.selected
        })
            .then(() => {
                console.log("submitted data");
                this.props.history.push("/statistics");
            });
    }

    handleSelectionFinish(selectedItems) {
        let selected = [];
        for (const v of selectedItems) {
            selected.push(v.props.color);
        }
        this.selected = selected;
    };
}

export default VotingPage;
