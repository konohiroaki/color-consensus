import React, {Component} from "react";
import axios from "axios";

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

import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
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

export default SideBar;
