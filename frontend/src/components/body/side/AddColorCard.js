import React, {Component} from "react";
import $ from "jquery";
import axios from "axios";
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
            name: "",
            code: ""
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
        const {lang, name, code} = this.state;
        console.log(lang, name, code);
        axios.post("http://localhost:5000/api/v1/colors", {
            lang: lang, name: name, code: code
        }).then(() => {
            this.props.updateColorList();
        });
        $("#add-color-lang").val("");
        $("#add-color-name").val("");
        $("#add-color-code").val("");
        this.setState({lang: "", name: "", code: ""});
    }

    render() {
        console.log("rendering add color card");
        return (
            <div>
                <a className="card btn bg-dark border border-secondary m-2" data-toggle="modal" data-target="#color-add-modal">
                    <div className="p-3">
                        <FontAwesomeIcon icon={faPlus}/>
                    </div>
                </a>
                <div className="modal fade" tabIndex="-1" id="color-add-modal">
                    <div className="modal-dialog modal-dialog-centered">
                        <div className="modal-content bg-dark">
                            <div className="modal-header">
                                <span className="modal-title">Add Color</span>
                            </div>
                            <div className="modal-body">
                                <form>
                                    {/* TODO: should be drop down */}
                                    <label htmlFor="add-color-lang" className="col-form-label text-left">Language:</label>
                                    <input type="text" className="form-control" id="add-color-lang" onChange={this.handleLangChange} placeholder="en"/>
                                    <label htmlFor="add-color-name" className="col-form-label text-left">Color Name:</label>
                                    <input type="text" className="form-control" id="add-color-name" onChange={this.handleNameChange} placeholder="red"/>
                                    <label htmlFor="add-color-code" className="col-form-label text-left">Base Color Code:</label>
                                    <input type="text" className="form-control" id="add-color-code" onChange={this.handleCodeChange} placeholder="#ff0000"/>
                                </form>
                            </div>
                            <div className="modal-footer">
                                <button type="button" className="btn btn-secondary" data-dismiss="modal">Cancel</button>
                                <button type="button" className="btn btn-primary" data-dismiss="modal" onClick={this.handleClick}>
                                    Add Color
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default AddColorCard;