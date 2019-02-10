import React, {Component} from "react";
import axios from "axios";
import {SelectableGroup, DeselectAll} from "react-selectable-fast";
import CandidateList from "./CandidateList";

class MainContent extends Component {

    constructor(props) {
        super(props);
        this.state = {
            target: {},
        };
        this.candidateSize = 31;
        this.candidates = [];
        this.selected = [];
        this.updateCandidates = this.updateCandidates.bind(this);
        this.handleSelecting = this.handleSelecting.bind(this);
        this.handleSelectionFinish = this.handleSelectionFinish.bind(this);
        this.submit = this.submit.bind(this);
    }

    handleSelecting(selectingItems) {
        // TODO: remove border between selected and selected
    };

    handleSelectionFinish(selectedItems) {
        let selected = [];
        for (const v of selectedItems) {
            selected.push(v.props.color);
        }
        this.selected = selected;
    };

    updateCandidates(target) {
        return axios.get("http://localhost:5000/api/v1/colors/candidates/" + target.code.substring(1)
                         + "?size=" + Math.pow(this.candidateSize, 2)).then(({data}) => {
            console.log("main content got candidate list from server", data, this.candidateSize);
            let list = [];
            for (let i = 0; i < this.candidateSize; i++) {
                let row = [];
                for (let j = 0; j < this.candidateSize; j++) {
                    row.push(data[i * this.candidateSize + j]);
                }
                list.push(row);
            }

            console.log(list);
            this.candidates = list;
            // FIXME: doesn't deselect on color change.
            this.selected = [];
            this.setState({target: target});
        });
    }

    submit() {
        const {lang, name} = this.state.target;

        axios.post("http://localhost:5000/api/v1/votes/" + lang + "/" + name, this.selected)
            .then(() => console.log("submitted data"));
    }

    // memo: this.props.target -> new target color
    //       this.state.target -> current target color
    render() {
        console.log("rendering main content");
        if (Object.entries(this.props.target).length === 0) {
            return <div/>;
        }
        if (this.props.target !== this.state.target) {
            this.updateCandidates(this.props.target);
        }

        return (
            <div className="container-fluid pt-3" style={Object.assign({overflowY: "auto"}, this.props.style)}>
                {/* TODO: add skip and see statistics button*/}
                <div className="row">
                    <div className="mr-auto ml-5">
                        <p>Language: {this.props.target.lang}</p>
                        <p>Color Name: {this.props.target.name}</p>
                    </div>
                    <div className="ml-auto">
                        {/* FIXME: doesn't work when inside <SelectableGroup> */}
                        <button className="btn btn-primary m-3" onClick={this.submit}>Submit</button>
                    </div>
                </div>

                <SelectableGroup
                    className="selectable"
                    clickClassName="tick"
                    enableDeselect
                    allowClickWithoutSelected={true}
                    duringSelection={this.handleSelecting}
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
}

export default MainContent;