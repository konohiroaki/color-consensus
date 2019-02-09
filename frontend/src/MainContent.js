import React, {Component} from "react";
import axios from "axios";
import {SelectableGroup, DeselectAll} from "react-selectable-fast";
import CandidateList from "./CandidateList";

class MainContent extends Component {

    constructor(props) {
        super(props);
        this.state = {
            target: {},
            candidates: [],
            selected: []
        };
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
        this.setState({selected: selected});
    };

    draw(target) {
        console.log("drawing");
        return axios.get("http://localhost:5000/api/v1/colors/candidates/" + target.code.substring(1))
            .then(this.updateCandidates.bind(null, target));
    }

    updateCandidates(target, {data}) {
        let list = [];
        for (let i = 0; i < 51; i++) {
            let row = [];
            for (let j = 0; j < 51; j++) {
                row.push(data[i * 51 + j]);
            }
            list.push(row);
        }
        console.log("will draw", target, this.state.target);
        this.setState({target: target, candidates: list});
    }

    submit() {
        const {lang, name} = this.state.target;

        axios.post("http://localhost:5000/api/v1/votes/" + lang + "/" + name, this.state.selected)
            .then(() => console.log("submitted data"));
    }

    // memo: this.props.target -> new target color
    //       this.state.target -> current target color
    render() {
        console.log("rendering main content");
        if (Object.entries(this.props.target).length === 0) {
            return <div/>;
        }
        console.log("is target and target same?", this.props.target, this.state.target);
        if (this.props.target !== this.state.target) {
            this.draw(this.props.target);
        }

        return (
            <div className="container-fluid pt-3" style={Object.assign({overflowY: "auto"}, this.props.style)}>
                {/* TODO: skip and see statistics button*/}
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
                    <CandidateList items={this.state.candidates}/>
                </SelectableGroup>
            </div>
        );
    }
}

export default MainContent;