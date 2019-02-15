import React, {Component} from "react";
import $ from "jquery";
import axios from "axios";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPlus} from "@fortawesome/free-solid-svg-icons";

class AddColorCard extends Component {

    constructor(props) {
        super(props);
        this.state = {
            lang: "",
            name: "",
            code: ""
        };
        this.handleClick = this.handleClick.bind(this);
    }

    shouldComponentUpdate() {
        return false;
    }

    handleClick() {
        axios.post("http://localhost:5000/api/v1/colors", this.state)
            .then(() => {
                this.props.updateColorList();
            });
        // TODO: should get rid of jquery. need to make it controlled component?
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
                                {/* TODO: should be drop down */}
                                Language:
                                <input type="text" className="form-control" id="add-color-lang" placeholder="en"
                                       onChange={e => this.setState({lang: e.target.value})}/>
                                Color Name:
                                <input type="text" className="form-control" id="add-color-name" placeholder="red"
                                       onChange={e => this.setState({name: e.target.value})}/>
                                Base Color Code:
                                <input type="text" className="form-control" id="add-color-code" placeholder="#ff0000"
                                       onChange={e => this.setState({code: e.target.value})}/>
                            </div>
                            <div className="modal-footer">
                                <button type="button" className="btn btn-secondary" data-dismiss="modal">Cancel</button>
                                {/* TODO: dismiss the modal only when submit is success */}
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