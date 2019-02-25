import React, {Component} from "react";
import {DeselectAll, SelectableGroup} from "react-selectable-fast";
import {Link} from "react-router-dom";
import axios from "axios";
import CandidateList from "./CandidateList";
import $ from "jquery";

class VotingPage extends Component {

    constructor(props) {
        super(props);
        this.state = {};

        this.candidateSize = 31;
        this.candidates = [];
        // TODO: empty the list after submit.
        this.selected = [];

        this.updateCandidateList = this.updateCandidateList.bind(this);
        this.submit = this.submit.bind(this);
        this.handleSelectionFinish = this.handleSelectionFinish.bind(this);
    }

    render() {
        if (this.props.target === undefined) {
            return null;
        }

        return (
            <div>
                <div className="row">
                    <div className="mr-auto ml-5">
                        <p>Language: {this.props.target.lang}</p>
                        <p>Color Name: {this.props.target.name}</p>
                    </div>

                    <div className="ml-auto">
                        <div>
                            <button className="btn btn-secondary m-3" onClick={() => this.props.history.push("/statistics")}>
                                Skip to statistics
                            </button>
                            <button className="btn btn-primary m-3" onClick={this.submit}>
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
            axios.get("http://localhost:5000/api/v1/colors/candidates/" + target.code.substring(1)
                      + "?size=" + Math.pow(this.candidateSize, 2))
                .then(({data}) => {
                    console.log("main content got candidate list from server");
                    this.candidates = data;
                    // FIXME: doesn't deselect on color change.
                    this.selected = [];
                    this.setState({target: target});
                });
        }
    }

    submit() {
        const userId = this.props.userId;
        if (userId === undefined || userId === null) {
            // FIXME: very tightly coupled code.
            $("#signup-login-modal").modal();
            // FIXME: should move to stats page after signing in without pressing the button again.
            return;
        }
        const {lang, name} = this.state.target;

        axios.post("http://localhost:5000/api/v1/votes/" + lang + "/" + name, this.selected)
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