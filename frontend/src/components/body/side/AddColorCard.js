import React, {Component} from "react";
import axios from "axios";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPlus} from "@fortawesome/free-solid-svg-icons";
import {actions as searchBar} from "../../../modules/searchBar";
import {connect} from "react-redux";
import $ from "jquery";
import {toast} from "react-toastify";
import {actions as colorCategory} from "../../../modules/colorCategory";

class AddColorCard extends Component {

    constructor(props) {
        super(props);
        this.state = {
            category: "",
            newCategory: "",
            name: "",
            code: ""
        };

        this.handleClick = this.handleClick.bind(this);
    }

    render() {
        console.log("rendering [add color card]");
        return <div>
            <Card/>
            <AddColorModal category={this.state.category} newCategory={this.state.newCategory}
                           categorySetter={input => this.setState({category: input})}
                           newCategorySetter={input => this.setState({newCategory: input})}
                           categories={this.props.categories}
                           name={this.state.name} nameSetter={input => this.setState({name: input})}
                           code={this.state.code} codeSetter={input => this.setState({code: input})}
                           handleClick={this.handleClick}/>
        </div>;
    }

    handleClick() {
        const request = {
            category: this.state.category !== "" ? this.state.category : this.state.newCategory,
            name: this.state.name,
            code: this.state.code,
        };
        axios.post(`${process.env.WEBAPI_HOST}/api/v1/colors`, request)
            .then(() => {
                this.props.fetchColors();
                this.props.fetchColorCategories();
                $("#color-add-modal").modal("toggle");
                this.setState({category: "", newCategory: "", name: "", code: ""});
            })
            .catch(({response}) => toast.warn(response.data.error.message));
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
                <div className="modal-header">
                    <span className="modal-title">Add Color</span>
                </div>
                <div className="modal-body">
                    <ColorCategoryInput category={props.category} newCategory={props.newCategory}
                                        categorySetter={props.categorySetter} newCategorySetter={props.newCategorySetter}
                                        categories={props.categories}/>
                    <ColorNameInput name={props.name} nameSetter={props.nameSetter}/>
                    <ColorCodeInput code={props.code} codeSetter={props.codeSetter}/>
                </div>
                <div className="modal-footer">
                    <button type="button" className="btn btn-secondary" data-dismiss="modal">Cancel</button>
                    <button type="button" className="btn btn-primary" onClick={props.handleClick}>
                        Add Color
                    </button>
                </div>
            </div>
        </div>
    </div>
);

const ColorCategoryInput = props => {
    const categories = props.categories
        .sort((a, b) => a > b ? 1 : -1)
        .map(c => <option key={c} value={c}>{c}</option>);
    const selectDisabled = props.newCategory !== "";
    const textDisabled = props.category !== "";

    return <div>
        <label className="mb-0">Color Category:</label>
        <div className="input-group">
            <select className="input-group-prepend custom-select" value={props.category}
                    onChange={e => props.categorySetter(e.target.value)} disabled={selectDisabled}>
                <option key="0" value="">Choose from dropdown</option>
                {categories}
            </select>
            <input type="text" className="form-control" value={props.newCategory}
                   placeholder="or input new" maxLength="20"
                   onChange={e => props.newCategorySetter(e.target.value)} disabled={textDisabled}/>
        </div>
    </div>;
};

const ColorNameInput = props => (
    <div>
        <label className="mb-0">Color Name:</label>
        <input type="text" className="form-control" placeholder="eg. Red" maxLength="30"
               value={props.name} onChange={e => props.nameSetter(e.target.value)}/>
    </div>
);

const ColorCodeInput = props => (
    <div>
        <label className="mb-0">Base Color Code:</label>
        <input type="text" className="form-control" placeholder="eg. #ff0000" maxLength="7"
               value={props.code} onChange={e => props.codeSetter(e.target.value)}/>
    </div>
);

const mapStateToProps = state => ({
    categories: state.colorCategory.categories
});

const mapDispatchToProps = dispatch => ({
    fetchColors: () => dispatch(searchBar.fetchColors()),
    fetchColorCategories: () => dispatch(colorCategory.fetchColorCategories()),
});

export default connect(mapStateToProps, mapDispatchToProps)(AddColorCard);
