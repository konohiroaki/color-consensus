import {Component} from "react";
import axios from "axios";
import $ from "jquery";
import React from "react";

class MainContent extends Component {
    constructor(props) {
        super(props);
        this.state = {
            candidates: [],
        };
        this.updateCandidates = this.updateCandidates.bind(this);
    }

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

    handleMouseDown(e) {
        const offsets = $("#color-picker").offset();
        console.log("mouse down", e.pageX - offsets.left, e.pageY - offsets.top);
    }

    handleMouseUp(e) {
        const offsets = $("#color-picker").offset();
        console.log("mouse up", e.pageX - offsets.left, e.pageY - offsets.top);
    }

    render() {
        return (
            <div id="color-picker" className="container-fluid" style={{overflow: "auto"}}
                 onMouseDown={this.handleMouseDown} onMouseUp={this.handleMouseUp}>
                <List items={this.state.candidates}/>
                {/*<SelectableGroup enableDeselect={true}>*/}
                {/*<List items={this.state.candidates}/>*/}
                {/*</SelectableGroup>*/}
            </div>
        );
    }
}

class List extends Component {

    shouldComponentUpdate(nextProps) {
        return nextProps.items !== this.props.items;
    }

    render() {
        if (this.props.items.length === 0) {
            return <div/>;
        }
        let list = [];
        for (let i = 0; i < 51; i++) {
            let row = [];
            for (let j = 0; j < 51; j++) {
                row.push(<CandidateCell key={i * 51 + j} color={this.props.items[i][j]}/>);
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

class CandidateCell extends Component {
    render() {
        return (
            <div style={{
                display: "inline-block", padding: "1px", margin: "0 -1px -1px 0", border: "1px solid transparent",
                userSelect: "none", userDrag: "none"
            }}>
                <div style={{width: "15px", height: "15px", backgroundColor: this.props.color}}/>
            </div>
        );
    }
}

export default MainContent;