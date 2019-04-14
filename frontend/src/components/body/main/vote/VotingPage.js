import React, {Component} from "react";
import {DeselectAll, SelectableGroup} from "react-selectable-fast";
import axios from "axios";
import ColorBoard from "./ColorBoard";
import {connect} from "react-redux";
import {actions as vote} from "../../../../modules/vote";

class VotingPage extends Component {

    constructor(props) {
        super(props);

        this.handleSelectionFinish = this.handleSelectionFinish.bind(this);
        this.handleSubmitClick = this.handleSubmitClick.bind(this);
        this.submit = this.submit.bind(this);
    }

    render() {
        if (this.props.baseColor === null) {
            return null;
        }
        console.log("rendering [voting page]",
            "baseColor.code:", this.props.baseColor.code,
            "codeList[0]:", this.props.colorCodeList.length !== 0 ? this.props.colorCodeList[0] : null);

        return <div>
            <VotingHeader baseColor={this.props.baseColor}/>
            <VotingPageButtons history={this.props.history} handleSubmitClick={this.handleSubmitClick}/>
            <SelectableGroup enableDeselect allowClickWithoutSelected
                             onSelectionFinish={this.handleSelectionFinish}>
                <DeselectAllButton/>
                <ColorBoard colorCodeList={this.props.colorCodeList}/>
            </SelectableGroup>
        </div>;
    }

    handleSelectionFinish(selectedItems) {
        this.props.setSelectedColorCodeList(selectedItems.map(item => item.props.colorCode));
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
        const {lang, name} = this.props.baseColor;
        const url = `${process.env.WEBAPI_HOST}/api/v1/votes`;
        const body = {"lang": lang, "name": name, "colors": this.props.selectedColorCodeList};
        axios.post(url, body).then(() => this.props.history.push("/statistics"));
    }
}

const VotingHeader = ({baseColor}) => (
    <div className="card bg-dark border border-secondary">
        <div className="card-body">
            <div className="row ml-0 mr-0">
                <div className="col-3 card bg-dark border border-secondary p-2 text-center">
                    <div className="row">
                        <span className="col-4 border-right border-secondary p-3">{baseColor.lang}</span>
                        <span className="col-8 p-3">{baseColor.name}</span>
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
    boardSideLength: state.board.sideLength,
    baseColor: state.board.baseColor,
    colorCodeList: state.board.colorCodeList,
    selectedColorCodeList: state.vote.selectedColorCodeList,
    userId: state.user.id,
});

const mapDispatchToProps = dispatch => ({
    setSelectedColorCodeList: colorCodeList => dispatch(vote.setSelectedColorCodeList(colorCodeList)),
});

export default connect(mapStateToProps, mapDispatchToProps)(VotingPage);
