import React, {Component, createRef} from "react";
import axios from "axios";
import {SelectableGroup, DeselectAll} from "react-selectable-fast";
import {SelectableCandidateCell} from "./CandidateCell";

class MainContent extends Component {

    constructor(props) {
        super(props);
        this.state = {
            target: {
                lang: "",
                name: "",
                code: ""
            },
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
        // TODO: add to selected in this.state
        let selected = [];
        for (const v of selectedItems) {
            selected.push(v.props.color);
        }
        this.setState({selected: selected});
    };

    //TODO: should draw when triggered.
    draw(lang, name, code) {
        // console.log(lang + ":" + name + ":" + code);
        axios.get("http://localhost:5000/api/v1/colors/candidates/" + code.substring(1)).then(this.updateCandidates);
    }

    updateCandidates({data}) {
        let list = [];
        for (let i = 0; i < 51; i++) {
            let row = [];
            for (let j = 0; j < 51; j++) {
                row.push(data[i * 51 + j]);
            }
            list.push(row);
        }
        this.setState({candidates: list});
    }

    componentDidMount() {
        //TODO: get color from sidebar?
        // const target = {lang: "en", name: "red", code: "#ff0000"};
        // axios.get("http://localhost:5000/api/v1/colors/candidates/" + target.code.substring(1)).then(this.updateCandidates);
        // this.setState({target: target});
    }

    submit() {
        const {lang, name} = this.state.target;

        axios.post("http://localhost:5000/api/v1/votes/" + lang + "/" + name, this.state.selected)
            .then(() => console.log("submitted data"));
    }

    render() {
        console.log("rendering main content");
        return (
            <div className="container-fluid pt-3" style={{overflow: "auto"}}>
                {/* TODO: skip and see statistics button*/}
                <div className="row">
                    <div className="mr-auto ml-5">
                        <p>Language: {this.state.target.lang}</p>
                        <p>Color Name: {this.state.target.name}</p>
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
                    <List items={this.state.candidates}/>
                </SelectableGroup>
            </div>
        );
    }
}

class List extends Component {

    shouldComponentUpdate(props) {
        return props.items !== this.props.items;
    }

    render() {
        console.log("rendering list");
        if (this.props.items.length === 0) {
            return <div/>;
        }
        let list = [];
        for (let i = 0; i < 51; i++) {
            let row = [];
            for (let j = 0; j < 51; j++) {
                row.push(<SelectableCandidateCell key={i * 51 + j} color={this.props.items[i][j]}/>);
            }
            list.push(<div key={i}>{row}</div>);
        }
        return (
            <div className="text-center" style={{lineHeight: "0", padding: "10px"}}>
                {list}
            </div>
        );
    }
}

export default MainContent;