import React, {Component} from "react";
import {SelectableGroup} from "react-selectable-fast";
import axios from "axios";
import ColorBoard from "./ColorBoard";
import {connect} from "react-redux";
import {actions as vote} from "../../../../modules/vote";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faArrowRight} from "@fortawesome/free-solid-svg-icons";
import {faInfoCircle} from "@fortawesome/free-solid-svg-icons";

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
            <VotingPageButtons className="border-bottom border-secondary p-3 mb-3"
                               history={this.props.history} handleSubmitClick={this.handleSubmitClick}/>
            <UserGuide/>
            <SelectableGroup enableDeselect allowClickWithoutSelected
                             onSelectionFinish={this.handleSelectionFinish}>
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

const VotingPageButtons = ({className, history, handleSubmitClick}) => (
    <div className={className + " row"}>
        <button className="ml-auto btn btn-secondary m-3" onClick={() => history.push("/statistics")}>
            <span>Go to statistics <FontAwesomeIcon icon={faArrowRight}/></span>
        </button>
        <button className="btn btn-primary m-3" onClick={handleSubmitClick}>
            <span>Submit <FontAwesomeIcon icon={faArrowRight}/></span>
        </button>
    </div>
);

const UserGuide = () => (
    <div className="col-8 m-auto card" style={{backgroundColor: "#444a50"}}>
        <div className="card-body">
            <h5>
                <span className="badge badge-info"><FontAwesomeIcon icon={faInfoCircle}/> Usage</span>
            </h5>
            <p>Drag and select colors below which you think that matches the color name chosen in the right side bar. Then press the <span className="badge badge-primary">Submit <FontAwesomeIcon icon={faArrowRight}/></span> button to submit the data and also see the
                statistics.</p>
        </div>
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
