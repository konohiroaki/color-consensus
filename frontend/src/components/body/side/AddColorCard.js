import React, {Component} from "react";
import axios from "axios";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPlus} from "@fortawesome/free-solid-svg-icons";
import {actions as searchBar} from "../../../modules/searchBar";
import {connect} from "react-redux";

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

    render() {
        console.log("rendering [add color card]");
        return <div>
            <Card/>
            <AddColorModal lang={this.state.lang} langSetter={input => this.setState({lang: input})}
                           name={this.state.name} nameSetter={input => this.setState({name: input})}
                           code={this.state.code} codeSetter={input => this.setState({code: input})}
                           handleClick={this.handleClick}/>
        </div>;
    }

    handleClick() {
        axios.post(`${process.env.WEBAPI_HOST}/api/v1/colors`, this.state)
            .then(() => this.props.fetchColors());
        this.setState({lang: "", name: "", code: ""});
    }
}

const Card = () => (
    <a className="card btn bg-dark border border-secondary m-2" data-toggle="modal" data-target="#color-add-modal">
        <div className="p-3">
            <FontAwesomeIcon icon={faPlus}/>
        </div>
    </a>
);

const AddColorModal = (props) => (
    <div className="modal fade" tabIndex="-1" id="color-add-modal">
        <div className="modal-dialog modal-dialog-centered">
            <div className="modal-content bg-dark">
                <ModalHeader/>
                <ModalBody props={props}/>
                <ModalFooter handleClick={props.handleClick}/>
            </div>
        </div>
    </div>
);

const ModalHeader = () => (
    <div className="modal-header">
        <span className="modal-title">Add Color</span>
    </div>
);

const ModalBody = ({props}) => (
    <div className="modal-body">
        {/* TODO: should be drop down */}
        <LanguageInput lang={props.lang} langSetter={props.langSetter}/>
        <ColorNameInput name={props.name} nameSetter={props.nameSetter}/>
        <ColorCodeInput code={props.code} codeSetter={props.codeSetter}/>
    </div>
);

const LanguageInput = props => (
    <div>
        <label className="mb-0">Language:</label>
        <input type="text" className="form-control" placeholder="en"
               value={props.lang} onChange={e => props.langSetter(e.target.value)}/>
    </div>
);

const ColorNameInput = props => (
    <div>
        <label className="mb-0">Color Name:</label>
        <input type="text" className="form-control" placeholder="red"
               value={props.name} onChange={e => props.nameSetter(e.target.value)}/>
    </div>
);

const ColorCodeInput = props => (
    <div>
        <label className="mb-0">Base Color Code:</label>
        <input type="text" className="form-control" placeholder="#ff0000"
               value={props.code} onChange={e => props.codeSetter(e.target.value)}/>
    </div>
);

const ModalFooter = ({handleClick}) => (
    <div className="modal-footer">
        <button type="button" className="btn btn-secondary" data-dismiss="modal">Cancel</button>
        {/* TODO: dismiss the modal only when submit is success */}
        <button type="button" className="btn btn-primary" data-dismiss="modal" onClick={handleClick}>
            Add Color
        </button>
    </div>
);

const mapDispatchToProps = dispatch => ({
    fetchColors: () => dispatch(searchBar.fetchColors()),
});

export default connect(null, mapDispatchToProps)(AddColorCard);
