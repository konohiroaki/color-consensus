import React, {Component} from "react";
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

    handleClick() {
        axios.post("http://localhost:5000/api/v1/colors", this.state)
            .then(() => {
                this.props.updateColorList();
            });
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
                <AddColorModal lang={this.state.lang}
                               langSetter={e => this.setState({lang: e.target.value})}
                               name={this.state.name}
                               nameSetter={e => this.setState({name: e.target.value})}
                               code={this.state.code}
                               codeSetter={e => this.setState({code: e.target.value})}
                               handleClick={this.handleClick}/>
            </div>
        );
    }
}

class AddColorModal extends Component {
    render() {
        return (
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
                                   value={this.props.lang} onChange={this.props.langSetter}/>
                            Color Name:
                            <input type="text" className="form-control" id="add-color-name" placeholder="red"
                                   value={this.props.name} onChange={this.props.nameSetter}/>
                            Base Color Code:
                            <input type="text" className="form-control" placeholder="#ff0000"
                                   value={this.props.code} onChange={this.props.codeSetter}/>
                        </div>
                        <div className="modal-footer">
                            <button type="button" className="btn btn-secondary" data-dismiss="modal">Cancel</button>
                            {/* TODO: dismiss the modal only when submit is success */}
                            <button type="button" className="btn btn-primary" data-dismiss="modal" onClick={this.props.handleClick}>
                                Add Color
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default AddColorCard;
