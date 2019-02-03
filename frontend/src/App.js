import React, {Component} from "react";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import axios from "axios";
import {SelectableGroup, createSelectable, SelectAll, DeselectAll} from "react-selectable-fast";

class App extends Component {

    render() {
        return (
            <div className="bg-dark text-light" style={{display: "flex", flexDirection: "column", height: "100%"}}>
                <Header style={{flex: "0 0 80px"}}/>
                <div style={{flex: "1 1 auto", display: "flex", flexDirection: "row"}}>
                    <MainContent/>
                    <SideBar className="border-left border-secondary"/>
                </div>
            </div>
        );
    }
}

class Header extends Component {
    render() {
        return (
            <nav className="navbar navbar-dark bg-dark border-bottom border-secondary" style={this.props.style}>
                <a className="navbar-brand" href="/">Color Consensus</a>
            </nav>
        );
    }
}

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

    afterSelect(selectedTargets) {
        console.log(selectedTargets.length);
    }

    render() {
        return (
            <div className="container-fluid">
                <div className="container-fluid bg-light" style={{overflow: "auto"}}>
                    <SelectableGroup
                        clickClassName="tick"
                        enableDeselect
                        allowClickWithoutSelected={true}
                        duringSelection={console.log("during selection")}
                        onSelectionClear={console.log("selection clear")}
                        onSelectionFinish={console.log("selection finish")}
                    >
                        <List items={this.state.candidates}/>
                    </SelectableGroup>
                </div>
            </div>
        );
    }
}

function CandidateCell(props) {
    if (props.selected || props.selecting) {
        return (
            <div style={{
                display: "inline-block", padding: "1px", margin: "0 -1px -1px 0",
                borderStyle: "dashed", borderWidth: "1px",
                borderTopColor: "#008888", borderRightColor: "#008888", borderBottomColor: "#008888", borderLeftColor: "#008888"
            }}>
                <div ref={props.selectableRef} style={{width: "15px", height: "15px", backgroundColor: props.color}}/>
            </div>
        );
    } else {
        return (
            <div style={{display: "inline-block", padding: "1px", margin: "0 -1px -1px 0", border: "1px dashed transparent"}}>
                <div ref={props.selectableRef} style={{width: "15px", height: "15px", backgroundColor: props.color}}/>
            </div>
        );
    }
}

const SelectableCandidateCell = createSelectable(CandidateCell);

class List extends Component {

    shouldComponentUpdate(nextProps) {
        return nextProps.items !== this.props.items;
    }

    render() {
        if (this.props.items.length === 0) {
            return <div/>;
        }
        let list = [];
        console.log(this.props.items);
        for (let i = 0; i < 51; i++) {
            let row = [];
            for (let j = 0; j < 51; j++) {
                row.push(<SelectableCandidateCell key={i * 51 + j} color={this.props.items[i][j]}/>);
            }
            list.push(<div key={i}>{row}</div>);
        }
        return (
            <div>
                <SelectAll className="selectable-button">
                    <button>Select all</button>
                </SelectAll>
                <DeselectAll className="selectable-button">
                    <button>Clear selection</button>
                </DeselectAll>
                <div style={{lineHeight: "0", padding: "10px"}}>
                    {list}
                </div>
            </div>
        );
    }
}

class SideBar extends Component {
    constructor(props) {
        super(props);
        this.state = {
            colorList: []
        };
        this.updateColorList = this.updateColorList.bind(this);
    }

    componentDidMount() {
        // TODO: remove domain when releasing.
        axios.get("http://localhost:5000/api/v1/colors/keys").then(this.updateColorList);
    }

    updateColorList({data}) {
        this.setState({colorList: data});
    }

    render() {
        // FIXME: make the search box work.
        // TODO: implement add modal.
        let colorList = [];
        let langSet = new Set();
        for (let v of this.state.colorList) {
            colorList.push(<ColorCard lang={v.lang} name={v.name} code={v.base_code} key={v.lang + ":" + v.name}/>);
            langSet.add(v.lang);
        }
        let langList = [];
        for (let v of langSet) {
            langList.push(<div className="dropdown-item" key={v}>{v}</div>);
        }

        return (
            <div className={this.props.className}>
                <div className="input-group">
                    <button className="btn btn-outline-secondary dropdown-toggle" type="button" data-toggle="dropdown">Language</button>
                    <div className="dropdown-menu">
                        {langList}
                    </div>
                    <input type="text" className="form-control"/>
                </div>
                <div style={{overflowY: "auto", height: "100%"}}>
                    {colorList}
                    <AddColorCard/>
                </div>
            </div>
        );
    }
}

class ColorCard extends Component {
    constructor(props) {
        super(props);
        this.render.bind(this);
    }

    handleClick(lang, name, code) {
        console.log(lang, name, code);
        // TODO: somehow affect to main content.
        // draw(lang, name, code);
    }

    render() {
        return (
            <a className="card btn bg-dark border border-secondary m-2" onClick={this.handleClick.bind(this, this.props.lang, this.props.name, this.props.code)}>
                <div className="row">
                    <div className="col-3 border-right border-secondary p-3">{this.props.lang}</div>
                    <div className="col-9 p-3">{this.props.name}</div>
                </div>
            </a>
        );
    }
}

import {faPlus} from "@fortawesome/free-solid-svg-icons";

class AddColorCard extends Component {

    static handleClick(e) {
        // TODO: show modal for adding color.
    }

    render() {
        return (
            <a className="card btn bg-dark border border-secondary m-2" onClick={this.handleClick}>
                <div className="p-3">
                    <FontAwesomeIcon icon={faPlus}/>
                </div>
            </a>
        );
    }
}

export default App;
