import React, {Component} from "react";
import axios from "axios";
import {SelectableGroup, DeselectAll} from "react-selectable-fast";
import CandidateList from "./CandidateList";
import {BrowserRouter as Router, Link, Route} from "react-router-dom";

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
        this.handleSelectionFinish = this.handleSelectionFinish.bind(this);
        this.submit = this.submit.bind(this);
    }

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
            console.log("main content got candidate list from server", this.candidateSize, data);
            this.candidates = data;
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
            <Router>
                <div className="container-fluid pt-3" style={Object.assign({overflowY: "auto"}, this.props.style)}>
                    <div className="row">
                        <div className="mr-auto ml-5">
                            <p>Language: {this.props.target.lang}</p>
                            <p>Color Name: {this.props.target.name}</p>
                        </div>

                        <div className="ml-auto">
                            <Route exact path="/" render={() => (
                                <div>
                                    <Link to={"/statistics"}>
                                        <button className="btn btn-secondary m-3">Skip to statistics</button>
                                    </Link>
                                    <button className="btn btn-primary m-3" onClick={this.submit}>Submit</button>
                                </div>
                            )}>
                            </Route>
                            <Route exact path="/statistics" render={() => (
                                <Link to={"/"}>
                                    <button className="btn btn-secondary m-3">Back to voting</button>
                                </Link>
                            )}>
                            </Route>
                        </div>
                    </div>

                    {/* TODO: using too much Route with render={} seems danger. switch to component={} */}
                    {/* TODO: add content here. */}
                    <Route path="/statistics" render={() => (
                        <div>Hey!</div>
                    )}/>
                    <Route exact path="/" render={() => (
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
                    )}>
                    </Route>
                </div>
            </Router>
        );
    }
}

export default MainContent;