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
        console.log("rendering sidebar");
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
        console.log("rendering color card");
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

    shouldComponentUpdate() {
        return false;
    }

    constructor(props) {
        super(props);
        this.state = {
            lang: "",
            name: ""
        };
        this.handleLangChange = this.handleLangChange.bind(this);
        this.handleNameChange = this.handleNameChange.bind(this);
        this.handleCodeChange = this.handleCodeChange.bind(this);
        this.handleClick = this.handleClick.bind(this);
    }

    handleLangChange(e) {
        this.setState({lang: e.target.value});
    }

    handleNameChange(e) {
        this.setState({name: e.target.value});
    }

    handleCodeChange(e) {
        this.setState({code: e.target.value});
    }

    handleClick() {
        // TODO: post lang and name to add it in db.
        const {lang, name} = this.state;
        console.log(lang, name);
        this.setState({lang: "", name: ""});
    }

    // FIXME: modal disappears when clicked. it shouldn't disappear except cancel or submit button or outside space.
    render() {
        console.log("rendering add color card");
        return (
            <a className="card btn bg-dark border border-secondary m-2" data-toggle="modal" data-target="#color-add-modal">
                <div className="p-3">
                    <FontAwesomeIcon icon={faPlus}/>
                </div>
                <div className="modal" tabIndex="-1" id="color-add-modal">
                    <div className="modal-dialog modal-dialog-centered">
                        <div className="modal-content bg-dark">
                            <div className="modal-header">
                                <span className="modal-title">Add Color</span>
                            </div>
                            <div className="modal-body">
                                <form>
                                    <label htmlFor="add-color-lang" className="col-form-label text-left">Language:</label>
                                    <input type="text" className="form-control" id="add-color-lang" value={this.state.lang} onChange={this.handleLangChange} placeholder="en"/>
                                    <label htmlFor="add-color-name" className="col-form-label text-left">Color Name:</label>
                                    <input type="text" className="form-control" id="add-color-name" value={this.state.name} onChange={this.handleNameChange} placeholder="red"/>
                                    <label htmlFor="add-color-code" className="col-form-label text-left">Base Color Code:</label>
                                    <input type="text" className="form-control" id="add-color-code" value={this.state.code} onChange={this.handleCodeChange} placeholder="#ff0000"/>
                                </form>
                            </div>
                            <div className="modal-footer">
                                <button type="button" className="btn btn-secondary" data-dismiss="modal">Cancel</button>
                                <button type="button" className="btn btn-primary" onClick={this.handleClick}>Add Color</button>
                            </div>
                        </div>
                    </div>
                </div>
            </a>
        );
    }
}

export default SideBar;
