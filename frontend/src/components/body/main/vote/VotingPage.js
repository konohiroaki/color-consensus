import React, {Component} from "react";
import {DeselectAll, SelectableGroup} from "react-selectable-fast";
import axios from "axios";
import ColorBoard from "./ColorBoard";
import update from "immutability-helper";

class VotingPage extends Component {

    constructor(props) {
        super(props);
        this.state = {};

        this.candidateSize = 31;
        this.candidates = [];
        this.selected = [];

        this.updateCandidateList = this.updateCandidateList.bind(this);
        this.handleSelectionFinish = this.handleSelectionFinish.bind(this);
        this.handleSubmitClick = this.handleSubmitClick.bind(this);
        this.submit = this.submit.bind(this);
    }

    render() {
        if (this.props.target === undefined) {
            return null;
        }

        return <div>
            <VotingHeader target={this.props.target}/>
            <VotingPageButtons history={this.props.history} handleSubmitClick={this.handleSubmitClick}/>
            <SelectableGroup enableDeselect allowClickWithoutSelected
                             onSelectionFinish={this.handleSelectionFinish}>
                <DeselectAllButton/>
                <ColorBoard colors={this.candidates} candidateSize={this.candidateSize}/>
            </SelectableGroup>
        </div>;
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
            const code = this.props.target.code.substring(1); // remove "#"
            const size = Math.pow(this.candidateSize, 2);
            const url = `${process.env.WEBAPI_HOST}/api/v1/colors/candidates/${code}?size=${size}`;
            axios.get(url).then(({data}) => {
                this.candidates = data;
                // FIXME: doesn't deselect on color change.
                this.selected = [];
                this.setState({target: this.props.target});
            });
        }
    }

    handleSelectionFinish(selectedItems) {
        this.selected = selectedItems.map(item => item.props.color);
    };

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
        const url = `${process.env.WEBAPI_HOST}/api/v1/votes`;
        const body = {"lang": lang, "name": name, "colors": this.selected};
        axios.post(url, body).then(() => this.props.history.push("/statistics"));
    }
}

const VotingHeader = (props) => (
    <div className="card bg-dark border border-secondary">
        <div className="card-body">
            <div className="row ml-0 mr-0">
                <div className="col-3 card bg-dark border border-secondary p-2 text-center">
                    <div className="row">
                        <span className="col-4 border-right border-secondary p-3">{props.target.lang}</span>
                        <span className="col-8 p-3">{props.target.name}</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
);

const VotingPageButtons = (props) => (
    <div className="row">
        <div className="ml-auto">
            <div>
                <button className="btn btn-secondary m-3" onClick={() => props.history.push("/statistics")}>
                    Skip to statistics
                </button>
                <button className="btn btn-primary m-3" onClick={props.handleSubmitClick}>
                    Submit
                </button>
            </div>
        </div>
    </div>
);

const DeselectAllButton = () => (
    <div className="row">
        <div className="ml-auto">
            <DeselectAll className="btn btn-secondary m-3">Clear</DeselectAll>
        </div>
    </div>
);

export default VotingPage;
