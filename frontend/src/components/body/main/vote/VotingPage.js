import React, {Component} from "react";
import {DeselectAll, SelectableGroup} from "react-selectable-fast";
import axios from "axios";
import ColorBoard from "./ColorBoard";
import {connect} from "react-redux";

class VotingPage extends Component {

    constructor(props) {
        super(props);
        this.state = {};
        this.selected = [];

        this.handleSelectionFinish = this.handleSelectionFinish.bind(this);
        this.handleSubmitClick = this.handleSubmitClick.bind(this);
        this.submit = this.submit.bind(this);
    }

    render() {
        console.log("rendering voting page");
        if (this.props.displayedColor === null) {
            return null;
        }

        return <div>
            <VotingHeader displayedColor={this.props.displayedColor}/>
            <VotingPageButtons history={this.props.history} handleSubmitClick={this.handleSubmitClick}/>
            <SelectableGroup enableDeselect allowClickWithoutSelected
                             onSelectionFinish={this.handleSelectionFinish}>
                <DeselectAllButton/>
                <ColorBoard colorCodes={this.props.displayedColorList} boardSideLength={this.props.boardSideLength}/>
            </SelectableGroup>
        </div>;
    }

    handleSelectionFinish(selectedItems) {
        this.selected = selectedItems.map(item => item.props.colorCode);
        console.log(this.selected);
    };

    handleSubmitClick() {
        const userId = this.props.userId;
        if (userId === undefined || userId === null) {
            this.props.loginModalRef.current.openLoginModal(this.submit);
        } else {
            this.submit();
        }
    }

    submit() {
        const {lang, name} = this.props.displayedColor;
        const url = `${process.env.WEBAPI_HOST}/api/v1/votes`;
        const body = {"lang": lang, "name": name, "colors": this.selected};
        axios.post(url, body).then(() => this.props.history.push("/statistics"));
    }
}

const VotingHeader = ({displayedColor}) => (
    <div className="card bg-dark border border-secondary">
        <div className="card-body">
            <div className="row ml-0 mr-0">
                <div className="col-3 card bg-dark border border-secondary p-2 text-center">
                    <div className="row">
                        <span className="col-4 border-right border-secondary p-3">{displayedColor.lang}</span>
                        <span className="col-8 p-3">{displayedColor.name}</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
);

const VotingPageButtons = ({history, handleSubmitClick}) => (
    <div className="row">
        <button className="ml-auto btn btn-secondary m-3" onClick={() => history.push("/statistics")}>
            Skip to statistics
        </button>
        <button className="btn btn-primary m-3" onClick={handleSubmitClick}>
            Submit
        </button>
    </div>
);

const DeselectAllButton = () => (
    <div className="row">
        <DeselectAll className="ml-auto btn btn-secondary m-3">Clear</DeselectAll>
    </div>
);

const mapStateToProps = state => ({
    displayedColor: state.colors.displayedColor,
    displayedColorList: state.colors.displayedColorList,
    boardSideLength: state.colors.boardSideLength,
    userId: state.user.id,
});

export default connect(mapStateToProps)(VotingPage);
