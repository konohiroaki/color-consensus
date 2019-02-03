import React, {Component, createRef} from "react";
import axios from "axios";
import {SelectableGroup} from "react-selectable-fast";
import {SelectableCandidateCell, Counter} from "./CandidateCell";

class MainContent extends Component {

    // TODO: move selected status to this level.
    constructor(props) {
        super(props);
        this.counterRef = createRef();
        this.state = {
            candidates: [],
        };
        this.updateCandidates = this.updateCandidates.bind(this);
        this.handleSelecting = this.handleSelecting.bind(this);
        this.handleSelectionFinish = this.handleSelectionFinish.bind(this);
    }

    handleSelecting(selectingItems) {
        this.counterRef.current.handleSelectin(selectingItems);
    };

    handleSelectionFinish(selectedItems) {
        this.counterRef.current.handleSelectionFinis(selectedItems);
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
        axios.get("http://localhost:5000/api/v1/colors/candidates/ff0000").then(this.updateCandidates);
    }

    render() {
        return (
            <div className="container-fluid" style={{overflow: "auto"}}>
                <Counter ref={this.counterRef}/>
                <SelectableGroup
                    className="selectable"
                    clickClassName="tick"
                    enableDeselect
                    allowClickWithoutSelected={true}
                    duringSelection={this.handleSelecting}
                    onSelectionFinish={this.handleSelectionFinish}>
                    <List items={this.state.candidates}/>
                </SelectableGroup>
            </div>
        );
    }
}

class List extends Component {

    render() {
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
            <div style={{lineHeight: "0", padding: "10px"}}>
                {list}
            </div>
        );
    }
}

export default MainContent;